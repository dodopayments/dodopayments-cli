import os from 'node:os';
import path from 'node:path';
import fs from 'node:fs';
import crypto from 'node:crypto';
import { machineIdSync } from 'node-machine-id';

export type Mode = 'test_mode' | 'live_mode';
export type Config = Partial<Record<Mode, string>>;
export type ResolvedCredentials = {
  apiKey: string;
  mode: Mode;
};
export type LogoutTarget = Mode | 'all';

const CONFIG_DIR = path.join(os.homedir(), '.dodopayments');
const CONFIG_PATH = path.join(CONFIG_DIR, 'config.json');
const LEGACY_CONFIG_PATH = path.join(CONFIG_DIR, 'api-key');
const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];
const ALGORITHM = 'aes-256-gcm';

let sessionMode: Mode | null = null;
let cachedKey: Buffer | null = null;

export function setSessionMode(mode: Mode) {
  sessionMode = mode;
}

export function getSessionMode(): Mode | null {
  return sessionMode;
}

function getEncryptionKey(): Buffer {
  if (cachedKey) return cachedKey;
  let id: string;
  try {
    id = machineIdSync();
  } catch {
    id = `${os.hostname()}-${os.userInfo().username}`;
  }
  cachedKey = crypto.pbkdf2Sync(id, 'dodopayments-cli-salt', 100000, 32, 'sha256');
  return cachedKey;
}

function encrypt(text: string): string {
  const iv = crypto.randomBytes(16);
  const cipher = crypto.createCipheriv(ALGORITHM, getEncryptionKey(), iv);
  let encrypted = cipher.update(text, 'utf8', 'hex');
  encrypted += cipher.final('hex');
  const authTag = cipher.getAuthTag().toString('hex');
  return `${iv.toString('hex')}:${authTag}:${encrypted}`;
}

function decrypt(text: string): string {
  const parts = text.split(':');
  if (parts.length !== 3) throw new Error('Invalid encrypted format');
  const [ivHex, tagHex, encrypted] = parts;
  if (!ivHex || !tagHex || !encrypted) throw new Error('Invalid encrypted format');
  const iv = Buffer.from(ivHex, 'hex');
  const authTag = Buffer.from(tagHex, 'hex');

  const decipher = crypto.createDecipheriv(ALGORITHM, getEncryptionKey(), iv);
  decipher.setAuthTag(authTag);
  let decrypted = decipher.update(encrypted, 'hex', 'utf8');
  decrypted += decipher.final('utf8');
  return decrypted;
}

function ensureConfigDir() {
  if (!fs.existsSync(CONFIG_DIR)) {
    fs.mkdirSync(CONFIG_DIR, { recursive: true, mode: 0o700 });
  }
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

  const jsonStr = JSON.stringify(filteredConfig, null, 2);
  const encrypted = encrypt(jsonStr);

  fs.writeFileSync(CONFIG_PATH, encrypted, { mode: 0o600 });
  
  try {
    // Ensure permissions are strictly 0o600 even if the file already existed
    fs.chmodSync(CONFIG_PATH, 0o600);
  } catch {
    // Ignore if chmod fails (e.g. on Windows)
  }
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
}

export async function clearConfig(target: LogoutTarget): Promise<{
  hadInvalidConfig: boolean;
  removedModes: Mode[];
}> {
  if (!(await configExists())) {
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

  return { hadInvalidConfig: false, removedModes: [target] };
}

import type { CommandContext } from '../tui/CommandContext';

export async function resolveCredentials(ctx?: CommandContext, prompt: boolean = true): Promise<ResolvedCredentials> {
  if (sessionMode) {
    const config = await readConfig();
    return { mode: sessionMode, apiKey: config[sessionMode]! };
  }

  if (!(await configExists())) {
    throw new Error('Sign in first. Run /login to get started.');
  }

  let config: Config;
  try {
    config = await readConfig();
  } catch {
    await resetConfig();
    throw new Error("Couldn't load credentials. Run /login to retry.");
  }

  const modes = getConfiguredModesFromConfig(config);

  if (modes.length === 0) {
    await resetConfig();
    throw new Error('No valid credentials. Run /login to retry.');
  }

  if (modes.length === 1 || !prompt) {
    const mode = modes[0] as Mode;
    return { mode, apiKey: config[mode]! };
  }

  if (!ctx) {
    throw new Error('Multiple environments configured. Run the CLI without arguments to choose one.');
  }

  const selectedMode = await ctx.promptSelect('Environment', modes.map((mode) => ({
    label: mode === 'test_mode' ? 'Test Mode' : 'Live Mode',
    value: mode,
  }))) as Mode;

  return { mode: selectedMode, apiKey: config[selectedMode]! };
}
