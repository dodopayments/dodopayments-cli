import { describe, expect, test } from 'bun:test';
import os from 'node:os';
import path from 'node:path';
import { assertOutsideHomes, realHomes } from '../../scripts/test-preload';
import { configDir } from './secure-store';

describe('test isolation', () => {
  test('the tests use a temporary config directory, not the home directory', () => {
    expect(configDir()).toBe(process.env.DODO_CONFIG_DIR!);
    expect(configDir().startsWith(os.tmpdir())).toBe(true);
    expect(() => assertOutsideHomes(configDir())).not.toThrow();
  });

  test('the guard stops a config directory in the home directory', () => {
    for (const home of realHomes()) {
      expect(() => assertOutsideHomes(path.join(home, '.dodopayments'))).toThrow('must not use');
      expect(() => assertOutsideHomes(home)).toThrow('must not use');
    }
    expect(() => assertOutsideHomes('/home/me/.dodopayments', ['/home/me'])).toThrow('must not use');
    expect(() => assertOutsideHomes('/home/me2/.dodopayments', ['/home/me'])).not.toThrow();
  });

  test('with no DODO_CONFIG_DIR, the config directory is in the home directory, and the guard stops it', () => {
    const saved = process.env.DODO_CONFIG_DIR;
    delete process.env.DODO_CONFIG_DIR;
    try {
      expect(configDir()).toBe(path.join(os.homedir(), '.dodopayments'));
      expect(() => assertOutsideHomes(configDir())).toThrow('must not use');
    } finally {
      if (saved !== undefined) process.env.DODO_CONFIG_DIR = saved;
    }
  });
});
