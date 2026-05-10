import os from 'node:os';
import path from 'node:path';
import fs from 'node:fs';
import keytar from 'keytar';

export type Mode = 'test_mode' | 'live_mode';
export type Config = Partial<Record<Mode, string>>;
export type ResolvedCredentials = {
  apiKey: string;
  mode: Mode;
};
export type LogoutTarget = Mode | 'all';

const SERVICE_NAME = 'dodopayments-cli';
const CONFIG_DIR = path.join(os.homedir(), '.dodopayments');
const CONFIG_PATH = path.join(CONFIG_DIR, 'api-key');
const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];

let sessionMode: Mode | null = null;

export function setSessionMode(mode: Mode) {
  sessionMode = mode;
}

export function getSessionMode(): Mode | null {
  return sessionMode;
}

async function migrate(): Promise<void> {
  if (fs.existsSync(CONFIG_PATH)) {
    try {
      const content = fs.readFileSync(CONFIG_PATH, 'utf-8');
      const config = JSON.parse(content) as Config;
      for (const mode of ALL_MODES) {
        const apiKey = config[mode];
        if (apiKey) {
          await keytar.setPassword(SERVICE_NAME, mode, apiKey);
        }
      }
    } catch {
      // Ignore migration errors
    } finally {
      fs.rmSync(CONFIG_PATH, { force: true });
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
  const configuredModes = getConfiguredModesFromConfig(config);

  if (configuredModes.length === 0) {
    await resetConfig();
    return;
  }

  for (const mode of ALL_MODES) {
    const apiKey = config[mode];
    if (apiKey) {
      await keytar.setPassword(SERVICE_NAME, mode, apiKey);
    } else {
      await keytar.deletePassword(SERVICE_NAME, mode);
    }
  }
}

export async function configExists(): Promise<boolean> {
  await migrate();
  for (const mode of ALL_MODES) {
    const password = await keytar.getPassword(SERVICE_NAME, mode);
    if (password) return true;
  }
  return false;
}

export async function readConfig(): Promise<Config> {
  await migrate();
  const config: Config = {};
  let hasAny = false;

  for (const mode of ALL_MODES) {
    const password = await keytar.getPassword(SERVICE_NAME, mode);
    if (password) {
      config[mode] = password;
      hasAny = true;
    }
  }

  if (!hasAny) {
    throw new Error('CONFIG_NOT_FOUND');
  }

  return config;
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
  for (const mode of ALL_MODES) {
    await keytar.deletePassword(SERVICE_NAME, mode);
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
