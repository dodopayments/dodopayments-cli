/**
 * Phase 4 throwaway: bootstrap entry for the all-blocks demo.
 * Deleted in Phase 8.
 *
 * Run: bun src/tui/_blocks_demo.ts
 */

await import('@opentui/solid/preload');
const { mountBlocksDemo } = await import('./_blocks_demo_app');
mountBlocksDemo();
