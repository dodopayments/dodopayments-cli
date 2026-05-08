import fs from 'node:fs';
import path from 'node:path';
import os from 'node:os';
import { spawn, spawnSync } from 'node:child_process';
import { fileURLToPath } from 'node:url';

const UPDATE_CACHE_PATH = path.join(os.homedir(), '.dodopayments', 'update-cache.json');
const PACKAGE_NAME = 'dodopayments-cli';
const SIX_HOURS = 6 * 60 * 60 * 1000;

export type UpdateDelta = 'patch' | 'minor' | 'major' | null;
export type InstallMethod = 'npm' | 'bun' | 'binary' | 'unknown';

export interface UpdateInfo {
  currentVersion: string;
  latestVersion: string;
  delta: Exclude<UpdateDelta, null>;
}

export interface UpdateCache {
  latestVersion: string;
  lastChecked: number;
  /** Set after a silent update is dispatched so the next launch can confirm completion. */
  pendingSilentUpdate?: { from: string; to: string; dispatchedAt: number };
}

/**
 * Checks NPM for a newer version (cached for 6 hours).
 * Returns full update info (current, latest, delta) or null when up-to-date.
 */
export async function checkForUpdates(currentVersion: string): Promise<UpdateInfo | null> {
  try {
    const cache = await refreshUpdateCache();
    if (!cache) return null;
    const delta = getVersionDelta(currentVersion, cache.latestVersion);
    if (!delta) return null;
    return { currentVersion, latestVersion: cache.latestVersion, delta };
  } catch {
    return null;
  }
}

async function refreshUpdateCache(): Promise<UpdateCache | null> {
  const configDir = path.dirname(UPDATE_CACHE_PATH);
  if (!fs.existsSync(configDir)) {
    fs.mkdirSync(configDir, { recursive: true });
  }

  let cache: UpdateCache | null = readCache();
  const now = Date.now();

  if (!cache || now - cache.lastChecked > SIX_HOURS) {
    try {
      const response = await fetch(`https://registry.npmjs.org/${PACKAGE_NAME}/latest`, {
        signal: AbortSignal.timeout(2000),
      });
      if (response.ok) {
        const data = (await response.json()) as { version: string };
        cache = {
          ...(cache ?? {}),
          latestVersion: data.version,
          lastChecked: now,
        };
        writeCache(cache);
      }
    } catch {}
  }
  return cache;
}

function readCache(): UpdateCache | null {
  if (!fs.existsSync(UPDATE_CACHE_PATH)) return null;
  try {
    return JSON.parse(fs.readFileSync(UPDATE_CACHE_PATH, 'utf-8')) as UpdateCache;
  } catch {
    return null;
  }
}

function writeCache(cache: UpdateCache): void {
  try {
    fs.writeFileSync(UPDATE_CACHE_PATH, JSON.stringify(cache));
  } catch {}
}

/**
 * Returns the kind of version bump from `current` to `latest`, or null if
 * `latest` is not strictly newer.
 */
export function getVersionDelta(current: string, latest: string): UpdateDelta {
  const c = parseSemver(current);
  const l = parseSemver(latest);
  if (!c || !l) return current !== latest ? 'major' : null;

  if (l.major > c.major) return 'major';
  if (l.major < c.major) return null;
  if (l.minor > c.minor) return 'minor';
  if (l.minor < c.minor) return null;
  if (l.patch > c.patch) return 'patch';
  if (l.patch < c.patch) return null;
  if (c.prerelease && !l.prerelease) return 'patch';
  return null;
}

function parseSemver(v: string): {
  major: number;
  minor: number;
  patch: number;
  prerelease: string | null;
} | null {
  const stripped = v.replace(/^v/, '');
  const [core, prerelease = null] = stripped.split('-', 2) as [string, string?];
  const parts = (core?.split('+')[0] ?? '').split('.').map((n) => parseInt(n, 10));
  if (parts.length < 3 || parts.some((n) => Number.isNaN(n))) return null;
  return {
    major: parts[0]!,
    minor: parts[1]!,
    patch: parts[2]!,
    prerelease: prerelease ? prerelease.split('+')[0] ?? null : null,
  };
}

/**
 * Detects how this CLI binary was installed. Used to choose the right
 * upgrade command (and to skip auto-update for standalone binaries where
 * we cannot safely overwrite the executable).
 */
export function detectInstallMethod(): InstallMethod {
  // Bun's --compile produces a single-file executable; in that case
  // import.meta.url points inside the embedded bundle and there is no
  // package.json on disk we can safely upgrade.
  const isBunCompiled = typeof (globalThis as any).Bun !== 'undefined'
    && process.argv[0]?.endsWith('dodo')
    && !canResolveScriptOnDisk();
  if (isBunCompiled) return 'binary';

  const scriptPath = resolveScriptPath();
  if (!scriptPath) return 'unknown';

  const nodeModulesIdx = scriptPath.lastIndexOf(`${path.sep}node_modules${path.sep}`);
  if (nodeModulesIdx === -1) return 'unknown';

  const installRoot = scriptPath.slice(0, nodeModulesIdx);
  if (installRoot.includes(`${path.sep}.bun${path.sep}`)) return 'bun';
  if (fs.existsSync(path.join(installRoot, 'bun.lockb'))) return 'bun';
  return 'npm';
}

function resolveScriptPath(): string | null {
  try {
    return fileURLToPath(import.meta.url);
  } catch {
    return null;
  }
}

function canResolveScriptOnDisk(): boolean {
  const p = resolveScriptPath();
  return !!p && fs.existsSync(p);
}

/**
 * Spawns a detached background process to upgrade the global install.
 * Resolves immediately; the upgrade itself runs after this process exits.
 * Writes a `pendingSilentUpdate` marker so the next launch can confirm.
 */
export function dispatchSilentUpdate(info: UpdateInfo, method: InstallMethod): boolean {
  if (method !== 'npm' && method !== 'bun') return false;

  const cmd = method === 'bun' ? 'bun' : 'npm';
  const args = method === 'bun'
    ? ['add', '-g', `${PACKAGE_NAME}@latest`]
    : ['install', '-g', `${PACKAGE_NAME}@latest`];

  try {
    const child = spawn(cmd, args, {
      detached: true,
      stdio: 'ignore',
      windowsHide: true,
    });
    child.unref();

    const cache = readCache() ?? { latestVersion: info.latestVersion, lastChecked: Date.now() };
    cache.pendingSilentUpdate = {
      from: info.currentVersion,
      to: info.latestVersion,
      dispatchedAt: Date.now(),
    };
    writeCache(cache);
    return true;
  } catch {
    return false;
  }
}

/**
 * Reads and clears any pending silent-update marker. Used on launch to surface
 * a one-line confirmation block when a previous background update completed.
 */
export function consumePendingSilentUpdate(): { from: string; to: string } | null {
  const cache = readCache();
  if (!cache?.pendingSilentUpdate) return null;
  const pending = cache.pendingSilentUpdate;
  delete cache.pendingSilentUpdate;
  writeCache(cache);
  return { from: pending.from, to: pending.to };
}

/**
 * Synchronous, foreground update — used by the `/update` command. Captures
 * stdout/stderr so the UI can surface them.
 */
export function runForegroundUpdate(method: InstallMethod): {
  ok: boolean;
  stdout: string;
  stderr: string;
  command: string;
} {
  const cmd = method === 'bun' ? 'bun' : 'npm';
  const args = method === 'bun'
    ? ['add', '-g', `${PACKAGE_NAME}@latest`]
    : ['install', '-g', `${PACKAGE_NAME}@latest`];
  const command = `${cmd} ${args.join(' ')}`;

  try {
    const result = spawnSync(cmd, args, { encoding: 'utf-8', windowsHide: true });
    return {
      ok: result.status === 0,
      stdout: result.stdout ?? '',
      stderr: result.stderr ?? '',
      command,
    };
  } catch (e: any) {
    return { ok: false, stdout: '', stderr: e?.message ?? String(e), command };
  }
}
