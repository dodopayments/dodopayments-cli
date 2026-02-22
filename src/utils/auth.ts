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

const CONFIG_DIR = path.join(os.homedir(), '.dodopayments');
const CONFIG_PATH = path.join(CONFIG_DIR, 'api-key');

function ensureConfigDir(): void {
  fs.mkdirSync(CONFIG_DIR, { recursive: true });
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
  ensureConfigDir();

  let existingConfig: Config = {};

  if (configExists()) {
    try {
      existingConfig = JSON.parse(fs.readFileSync(CONFIG_PATH, 'utf-8'));
    } catch {
      existingConfig = {};
    }
  }

  existingConfig[mode] = apiKey;
  fs.writeFileSync(
    CONFIG_PATH,
    JSON.stringify(existingConfig, null, 2),
    'utf-8',
  );
}

export function resetConfig(): void {
  fs.rmSync(CONFIG_PATH, { force: true });
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

  const modes = Object.keys(config) as Mode[];

  if (modes.length === 0) {
    console.error('No valid credentials found. Please login again.');
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
