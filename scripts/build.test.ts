import { describe, expect, test } from 'bun:test';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { version } from '../package.json';

const nodePath = Bun.which('node');
if (!nodePath) throw new Error('Node.js is required to test the npm artifact');

function runWithNode(args: string[]) {
  const isolatedHome = mkdtempSync(join(tmpdir(), 'dodo-cli-test-'));
  const nodeBinDirectory = dirname(nodePath);

  try {
    return spawnSync(nodePath, [resolve('dist/index.js'), ...args], {
      encoding: 'utf8',
      env: {
        ...process.env,
        HOME: isolatedHome,
        USERPROFILE: isolatedHome,
        PATH: nodeBinDirectory,
        Path: nodeBinDirectory,
      },
    });
  } finally {
    rmSync(isolatedHome, { recursive: true, force: true });
  }
}

describe('npm artifact under Node.js without Bun', () => {
  test('prints its version', () => {
    const result = runWithNode(['--version']);

    expect(result.status).toBe(0);
    expect(result.stdout).toBe(`v${version}\n`);
    expect(result.stderr).not.toContain('Error');
  });

  test('reaches the login command', () => {
    const result = runWithNode(['login']);

    expect(result.status).toBe(0);
    expect(result.stdout).toContain('Usage: dodo login <api-key> <test|live>');
    expect(result.stderr).toBe('API key and mode are required.\n');
  });
});
