/**
 * Runs before the test files (bunfig.toml [test] preload). The tests use a temporary DODO_CONFIG_DIR.
 * A test cannot use HOME for this: Bun's os.homedir() ignores a HOME that changes at run time.
 */
import { beforeEach } from 'bun:test';
import { mkdtempSync, realpathSync, rmSync } from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import { configDir } from '../src/utils/secure-store';

function canonical(dir: string): string {
  try {
    return realpathSync(dir);
  } catch {
    return path.resolve(dir);
  }
}

/** The home directory of the user: the HOME of the process, and the home in the user database. */
export function realHomes(): string[] {
  const homes = [os.homedir()];
  try {
    homes.push(os.userInfo().homedir);
  } catch {
    // The user database has no entry for this user.
  }
  return homes;
}

/** Throws when `dir` is one of the `homes` directories, or is in one of them. */
export function assertOutsideHomes(dir: string, homes: string[] = realHomes()): void {
  const target = canonical(dir);
  for (const home of homes) {
    const relative = path.relative(canonical(home), target);
    const inside = relative === '' || (relative !== '..' && !relative.startsWith(`..${path.sep}`) && !path.isAbsolute(relative));
    if (inside) {
      throw new Error(`The tests must not use ${dir}. It is in the home directory ${home}. Set DODO_CONFIG_DIR to a temporary directory.`);
    }
  }
}

const root = mkdtempSync(path.join(os.tmpdir(), 'dodo-cli-test-'));
process.env.DODO_CONFIG_DIR = path.join(root, '.dodopayments');
assertOutsideHomes(configDir());
// A test can change DODO_CONFIG_DIR. The guard examines the directory again before each test.
beforeEach(() => assertOutsideHomes(configDir()));
process.on('exit', () => rmSync(root, { recursive: true, force: true }));
