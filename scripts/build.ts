/**
 * Production bundler. Bun's `bun build` does NOT honor bunfig.toml's
 * preload (per OpenTUI #122), so we drive Bun.build directly with the
 * @opentui/solid bun-plugin to apply the Solid babel transform on every
 * `.tsx/.jsx` file.
 *
 * The runtime path uses bunfig.toml's preload instead -- both code paths
 * end up registering the same plugin, just at different lifecycles.
 */

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
if (main) {
  const sizeKb = (main.size / 1024).toFixed(0);
  console.log(`\u2713 Built ${main.path} (${sizeKb} KB)`);
}
