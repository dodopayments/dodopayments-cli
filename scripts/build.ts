/**
 * Production bundler. Bun's `bun build` does NOT honor bunfig.toml's
 * preload (per OpenTUI #122), so we drive Bun.build directly with the
 * @opentui/solid bun-plugin to apply the Solid babel transform on every
 * `.tsx/.jsx` file.
 *
 * The runtime path uses bunfig.toml's preload instead -- both code paths
 * end up registering the same plugin, just at different lifecycles.
 *
 * After bundling, we inject a tiny Node-compatible prelude that re-execs
 * the bundle under Bun if it's launched by Node. This is required because
 * `target: 'bun'` emits Bun-only runtime APIs (e.g. `import.meta.require`)
 * and our native deps (@opentui/core) load through Bun's loader.
 *
 * Without this prelude, `bun install -g dodopayments-cli` works on disk
 * but the symlinked `dodo` binary is invoked via `#!/usr/bin/env node`,
 * which crashes immediately with "Qg is not a function" because Node has
 * no `import.meta.require`. The prelude makes Node hand control off to
 * Bun transparently, and gives a clear error if Bun isn't installed.
 */

import { readFileSync, writeFileSync } from 'node:fs';
import solidPlugin from '@opentui/solid/bun-plugin';

const result = await Bun.build({
  entrypoints: ['./src/index.ts'],
  outdir: './dist',
  target: 'bun',
  minify: true,
  plugins: [solidPlugin],
});

if (!result.success) {
  console.error('Build failed.');
  for (const message of result.logs) console.error(message);
  process.exit(1);
}

const main = result.outputs.find((o) => o.kind === 'entry-point');
if (!main) {
  console.error('Build succeeded but no entry-point output was produced.');
  process.exit(1);
}

const outPath = main.path;

/**
 * Node-safe prelude. Must contain ZERO Bun-only APIs.
 *
 * Why this works:
 *   - Top-level await in ESM BLOCKS subsequent top-level statements
 *     until it resolves. So under Node, the `await import(...)` inside
 *     the `if` block runs to completion (and calls process.exit) before
 *     any of the bundle's Bun-only top-level code is reached.
 *   - Under Bun, the `if` is skipped and the bundle runs normally.
 *
 * Why not wrap the rest of the bundle in `else { ... }`:
 *   - The bundle contains top-level `import ...` declarations, which
 *     are syntactically forbidden inside a block. So we cannot wrap.
 *     Instead we rely on `process.exit()` inside the `if` to keep Node
 *     from ever continuing past it.
 *
 * Bun detection:
 *   - `typeof Bun !== 'undefined'` is the documented way to detect Bun.
 *     Node has no such global. Deno also lacks it; Deno would also fail
 *     on other Bun-only bits below, so the explicit error message is
 *     the correct outcome for any non-Bun runtime.
 */
const prelude = `if (typeof Bun === 'undefined') {
  const { spawnSync } = await import('node:child_process');
  const { existsSync } = await import('node:fs');
  const { fileURLToPath } = await import('node:url');
  const path = await import('node:path');
  const os = await import('node:os');

  const scriptPath = fileURLToPath(import.meta.url);
  const isWindows = process.platform === 'win32';

  const which = spawnSync(isWindows ? 'where' : 'which', ['bun'], { encoding: 'utf8' });
  let bun = null;
  if (which.status === 0 && which.stdout) {
    const first = which.stdout.split(/\\r?\\n/).find(Boolean);
    if (first && existsSync(first)) bun = first;
  }
  if (!bun) {
    const candidates = [
      path.join(os.homedir(), '.bun', 'bin', isWindows ? 'bun.exe' : 'bun'),
      '/usr/local/bin/bun',
      '/opt/homebrew/bin/bun',
    ];
    for (const c of candidates) {
      if (existsSync(c)) { bun = c; break; }
    }
  }

  if (!bun) {
    process.stderr.write(
      'dodopayments-cli requires the Bun runtime when installed via npm/bun.\\n' +
      '\\n' +
      'Install Bun:    https://bun.com/docs/installation\\n' +
      'Or download a standalone binary (no runtime required):\\n' +
      '                https://github.com/dodopayments/dodopayments-cli/releases\\n'
    );
    process.exit(1);
  }

  const child = spawnSync(bun, [scriptPath, ...process.argv.slice(2)], {
    stdio: 'inherit',
    env: process.env,
  });
  if (child.error) {
    process.stderr.write('Failed to launch bun: ' + child.error.message + '\\n');
    process.exit(1);
  }
  if (typeof child.signal === 'string' && child.signal) {
    process.kill(process.pid, child.signal);
  }
  process.exit(child.status ?? 0);
}
`;

const original = readFileSync(outPath, 'utf8');

// Splice the prelude immediately after the shebang (Bun emits it on line 1
// when bundling an executable). Everything below the shebang is untouched.
const newlineIdx = original.indexOf('\n');
const firstLine = newlineIdx >= 0 ? original.slice(0, newlineIdx) : '';
const hasShebang = firstLine.startsWith('#!');
const shebang = hasShebang ? firstLine : '';
const body = hasShebang ? original.slice(newlineIdx + 1) : original;

const patched = (shebang ? `${shebang}\n` : '') + prelude + body;
writeFileSync(outPath, patched);

const sizeKb = (Buffer.byteLength(patched, 'utf8') / 1024).toFixed(0);
console.log(`\u2713 Built ${outPath} (${sizeKb} KB)`);
