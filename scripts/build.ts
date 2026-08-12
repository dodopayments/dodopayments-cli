/**
 * Production bundler. Bun's `bun build` does NOT honor bunfig.toml's
 * preload (per OpenTUI #122), so we drive Bun.build directly with the
 * @opentui/solid bun-plugin to apply the Solid babel transform on every
 * `.tsx/.jsx` file.
 *
 * The npm artifact targets Node so package-manager installs can run CLI
 * commands without Bun. OpenTUI still needs Bun for its FFI implementation;
 * src/index.ts re-launches only the interactive TUI under Bun.
 */

import solidPlugin from '@opentui/solid/bun-plugin';

const result = await Bun.build({
  entrypoints: ['./src/index.ts'],
  outdir: './dist',
  target: 'node',
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

const sizeKb = (main.size / 1024).toFixed(0);
console.log(`\u2713 Built ${main.path} (${sizeKb} KB)`);
