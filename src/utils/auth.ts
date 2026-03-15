import os from 'node:os';
import path from 'node:path';
import fs from 'node:fs';
import { select } from '@inquirer/prompts';

export type Mode = 'test_mode' | 'live_mode';
export type Config = Partial<Record<Mode, string>>;
export type ResolvedCredentials = {
  apiKey: string;
  mode: Mode;
};
export type LogoutTarget = Mode | 'all';

const CONFIG_DIR = path.join(os.homedir(), '.dodopayments');
const CONFIG_PATH = path.join(CONFIG_DIR, 'api-key');
const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];

function ensureConfigDir(): void {
  fs.mkdirSync(CONFIG_DIR, { recursive: true });
}

function getConfiguredModesFromConfig(config: Config): Mode[] {
  return ALL_MODES.filter((mode) => {
    const apiKey = config[mode];
    return typeof apiKey === 'string' && apiKey.length > 0;
  });
}

function writeConfig(config: Config): void {
  const configuredModes = getConfiguredModesFromConfig(config);

  if (configuredModes.length === 0) {
    resetConfig();
    return;
  }

  ensureConfigDir();

  const sanitizedConfig = configuredModes.reduce<Config>((result, mode) => {
    result[mode] = config[mode];
    return result;
  }, {});

  fs.writeFileSync(CONFIG_PATH, JSON.stringify(sanitizedConfig, null, 2), 'utf-8');
}

export function configExists(): boolean {
  return fs.existsSync(CONFIG_PATH);
}

export function readConfig(): Config {
  if (!configExists()) {
    throw new Error('CONFIG_NOT_FOUND');
  }

  try {
    const parsed = JSON.parse(fs.readFileSync(CONFIG_PATH, 'utf-8'));

    if (typeof parsed !== 'object' || parsed === null) {
      throw new Error('INVALID_CONFIG');
    }

    return parsed as Config;
  } catch {
    throw new Error('INVALID_CONFIG');
  }
}

export function saveConfig(mode: Mode, apiKey: string): void {
  let existingConfig: Config = {};

  if (configExists()) {
    try {
      existingConfig = readConfig();
    } catch {
      existingConfig = {};
    }
  }

  existingConfig[mode] = apiKey;
  writeConfig(existingConfig);
}

export function resetConfig(): void {
  fs.rmSync(CONFIG_PATH, { force: true });
}

export function clearConfig(target: LogoutTarget): {
  hadInvalidConfig: boolean;
  removedModes: Mode[];
} {
  if (!configExists()) {
    return { hadInvalidConfig: false, removedModes: [] };
  }

  let config: Config;
  try {
    config = readConfig();
  } catch {
    resetConfig();
    return { hadInvalidConfig: true, removedModes: [] };
  }

  const configuredModes = getConfiguredModesFromConfig(config);

  if (target === 'all') {
    resetConfig();
    return { hadInvalidConfig: false, removedModes: configuredModes };
  }

  const hasStoredMode = Object.prototype.hasOwnProperty.call(config, target);

  if (!hasStoredMode) {
    return { hadInvalidConfig: false, removedModes: [] };
  }

  delete config[target];
  writeConfig(config);

  return { hadInvalidConfig: false, removedModes: [target] };
}

export async function resolveCredentials(): Promise<ResolvedCredentials> {
  if (!configExists()) {
    console.error('Please run `dodo login` first.');
    process.exit(1);
  }

  let config: Config;
  try {
    config = readConfig();
  } catch {
    console.error('Failed to load credentials. Please try login again.');
    resetConfig();
    process.exit(1);
  }

  const modes = getConfiguredModesFromConfig(config);

  if (modes.length === 0) {
    console.error('No valid credentials found. Please login again.');
    resetConfig();
    process.exit(1);
  }

  if (modes.length === 1) {
    const mode = modes[0] as Mode;
    return { mode, apiKey: config[mode]! };
  }

  const selectedMode = await select<Mode>({
    message: 'Choose the environment:',
    choices: modes.map((mode) => ({
      name: mode === 'test_mode' ? 'Test Mode' : 'Live Mode',
      value: mode,
    })),
  });

  return { mode: selectedMode, apiKey: config[selectedMode]! };
}
