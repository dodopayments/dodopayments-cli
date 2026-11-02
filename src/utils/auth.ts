import path from 'node:path';
import fs from 'node:fs';
import { CONFIG_DIR, decrypt, ensureConfigDir, writeEncrypted } from './secure-store';
import { oauthCredentials, oauthModes } from './oauth';

export type Mode = 'test_mode' | 'live_mode';
export type Config = Partial<Record<Mode, string>>;
export type ResolvedCredentials = {
  apiKey: string;
  mode: Mode;
  /** `oauth` when the token comes from `dodo login` in the browser or with a device code. */
  source?: 'api_key' | 'oauth';
};
export type LogoutTarget = Mode | 'all';

const CONFIG_PATH = path.join(CONFIG_DIR, 'config.json');
const STATE_PATH = path.join(CONFIG_DIR, 'state.json');
const LEGACY_CONFIG_PATH = path.join(CONFIG_DIR, 'api-key');
const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];

let sessionMode: Mode | null = null;

export type State = {
  mode?: Mode;
};

function readState(): State {
  try {
    if (fs.existsSync(STATE_PATH)) {
      return JSON.parse(fs.readFileSync(STATE_PATH, 'utf-8'));
    }
  } catch {
    // Ignore errors
  }
  return {};
}

function writeState(state: State) {
  ensureConfigDir();
  fs.writeFileSync(STATE_PATH, JSON.stringify(state, null, 2));
}

export function setSessionMode(mode: Mode) {
  sessionMode = mode;
  writeState({ mode });
}

export function getSessionMode(): Mode | null {
  if (!sessionMode) {
    const state = readState();
    if (state.mode) {
      sessionMode = state.mode;
    }
  }
  return sessionMode;
}

export function clearSessionMode() {
  sessionMode = null;
  writeState({});
}

async function migrate(): Promise<void> {
  // Try to migrate from legacy api-key file if config.json doesn't exist
  if (!fs.existsSync(CONFIG_PATH) && fs.existsSync(LEGACY_CONFIG_PATH)) {
    try {
      const content = fs.readFileSync(LEGACY_CONFIG_PATH, 'utf-8');
      const legacyConfig = JSON.parse(content) as Config;
      await writeConfig(legacyConfig);
    } catch {
      // Ignore migration errors
    } finally {
      try {
        fs.rmSync(LEGACY_CONFIG_PATH, { force: true });
      } catch {
        // Ignore rm errors
      }
    }
  }
}

function getConfiguredModesFromConfig(config: Config): Mode[] {
  return ALL_MODES.filter((mode) => {
    const apiKey = config[mode];
    return typeof apiKey === 'string' && apiKey.length > 0;
  });
}

async function writeConfig(config: Config): Promise<void> {
  ensureConfigDir();
  
  const configuredModes = getConfiguredModesFromConfig(config);

  if (configuredModes.length === 0) {
    await resetConfig();
    return;
  }

  // Filter config to only include configured modes
  const filteredConfig: Config = {};
  for (const mode of configuredModes) {
    filteredConfig[mode] = config[mode];
  }

  writeEncrypted(CONFIG_PATH, JSON.stringify(filteredConfig, null, 2));
}

export async function readConfig(): Promise<Config> {
  await migrate();
  if (!fs.existsSync(CONFIG_PATH)) {
    throw new Error('CONFIG_NOT_FOUND');
  }

  const raw = fs.readFileSync(CONFIG_PATH, 'utf-8');
  let config: Config;

  try {
    // Try to decrypt
    const decrypted = decrypt(raw);
    config = JSON.parse(decrypted) as Config;
  } catch (err) {
    // If decryption fails, it might be in plaintext (from a previous version)
    try {
      config = JSON.parse(raw) as Config;
      // Re-save as encrypted
      await writeConfig(config);
    } catch {
      throw new Error('CONFIG_NOT_FOUND');
    }
  }

  if (getConfiguredModesFromConfig(config).length === 0) {
    throw new Error('CONFIG_NOT_FOUND');
  }

  return config;
}

export async function configExists(): Promise<boolean> {
  await migrate();
  if (process.env.DODO_PAYMENTS_API_KEY) return true;
  if (oauthModes().length > 0) return true;
  if (!fs.existsSync(CONFIG_PATH)) return false;
  
  try {
    const config = await readConfig();
    return getConfiguredModesFromConfig(config).length > 0;
  } catch {
    return false;
  }
}

export async function saveConfig(mode: Mode, apiKey: string): Promise<void> {
  let existingConfig: Config = {};

  if (await configExists()) {
    try {
      existingConfig = await readConfig();
    } catch {
      existingConfig = {};
    }
  }

  existingConfig[mode] = apiKey;
  await writeConfig(existingConfig);
}

export async function resetConfig(): Promise<void> {
  if (fs.existsSync(CONFIG_PATH)) {
    try {
      fs.unlinkSync(CONFIG_PATH);
    } catch {
      // Ignore errors
    }
  }
  clearSessionMode();
}

export async function clearConfig(target: LogoutTarget): Promise<{
  hadInvalidConfig: boolean;
  removedModes: Mode[];
}> {
  // Only the API key file. The OAuth sign-in and the environment variable are not in it.
  await migrate();
  if (!fs.existsSync(CONFIG_PATH)) {
    return { hadInvalidConfig: false, removedModes: [] };
  }

  let config: Config;
  try {
    config = await readConfig();
  } catch {
    await resetConfig();
    return { hadInvalidConfig: true, removedModes: [] };
  }

  const configuredModes = getConfiguredModesFromConfig(config);

  if (target === 'all') {
    await resetConfig();
    return { hadInvalidConfig: false, removedModes: configuredModes };
  }

  const hasStoredMode = Object.prototype.hasOwnProperty.call(config, target);

  if (!hasStoredMode) {
    return { hadInvalidConfig: false, removedModes: [] };
  }

  delete config[target];
  await writeConfig(config);

  if (getSessionMode() === target) {
    clearSessionMode();
  }

  return { hadInvalidConfig: false, removedModes: [target] };
}

import type { CommandContext } from '../tui/CommandContext';

async function readConfigOrEmpty(): Promise<Config> {
  await migrate();
  if (!fs.existsSync(CONFIG_PATH)) return {};
  try {
    return await readConfig();
  } catch {
    return {};
  }
}

/** The modes that have a pasted API key. */
export async function apiKeyModes(): Promise<Mode[]> {
  return getConfiguredModesFromConfig(await readConfigOrEmpty());
}

export function configPath(): string {
  return CONFIG_PATH;
}

/** The modes that have a pasted API key or an OAuth sign-in. */
export async function signedInModes(): Promise<Mode[]> {
  return availableModes(await readConfigOrEmpty());
}

async function availableModes(config: Config): Promise<Mode[]> {
  const oauth = oauthModes();
  return ALL_MODES.filter((mode) => Boolean(config[mode]) || oauth.includes(mode));
}

/** A pasted API key comes first, then the OAuth sign-in. */
async function credentialFor(config: Config, mode: Mode): Promise<ResolvedCredentials> {
  const apiKey = config[mode];
  if (apiKey) return { mode, apiKey, source: 'api_key' };
  return { mode, apiKey: await oauthCredentials(mode), source: 'oauth' };
}

export async function resolveCredentials(
  ctx?: CommandContext,
  prompt: boolean = true,
  forcePrompt: boolean = false,
): Promise<ResolvedCredentials> {
  const envKey = process.env.DODO_PAYMENTS_API_KEY;
  if (envKey) {
    const envMode = process.env.DODO_PAYMENTS_ENVIRONMENT === 'live_mode' ? 'live_mode' : 'test_mode';
    return { mode: envMode, apiKey: envKey };
  }

  if (!sessionMode) {
    const state = readState();
    if (state.mode) {
      sessionMode = state.mode;
    }
  }

  const config = await readConfigOrEmpty();
  const modes = await availableModes(config);

  if (modes.length === 0) {
    throw new Error('Sign in first. Run /login to get started.');
  }

  if (sessionMode && !forcePrompt) {
    // A persisted mode with no credentials is ignored.
    if (modes.includes(sessionMode)) {
      return credentialFor(config, sessionMode);
    }
    sessionMode = null;
  }

  if (modes.length === 1 || (!prompt && !forcePrompt)) {
    const mode = modes[0] as Mode;
    if (prompt) setSessionMode(mode);
    return credentialFor(config, mode);
  }

  if (!ctx) {
    throw new Error('Multiple environments configured. Run the CLI without arguments to choose one.');
  }

  const selectedMode = (await ctx.promptSelect(
    'Environment',
    modes.map((mode) => ({
      label: mode === 'test_mode' ? 'Test Mode' : 'Live Mode',
      value: mode,
    })),
  )) as Mode;

  setSessionMode(selectedMode);
  return credentialFor(config, selectedMode);
}
