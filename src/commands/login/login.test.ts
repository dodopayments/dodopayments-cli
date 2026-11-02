import { expect, test } from 'bun:test';
import { mkdtempSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

test('dodo login with no terminal exits with 1', async () => {
  const home = mkdtempSync(join(tmpdir(), 'dodo-login-test-'));
  try {
    const child = Bun.spawn([process.execPath, 'src/index.ts', 'login'], {
      cwd: join(import.meta.dir, '..', '..', '..'),
      env: { ...process.env, HOME: home, USERPROFILE: home },
      stdin: 'ignore',
      stdout: 'pipe',
      stderr: 'pipe',
    });
    const [code, , stderr] = await Promise.all([
      child.exited,
      new Response(child.stdout).text(),
      new Response(child.stderr).text(),
    ]);
    expect(stderr).toContain('No terminal for the sign-in');
    expect(code).toBe(1);
  } finally {
    rmSync(home, { recursive: true, force: true });
  }
}, 30_000);
