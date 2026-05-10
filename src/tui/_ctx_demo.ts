/**
 * Phase 3 throwaway: bootstrap entry for CommandContext demo.
 * Deleted in Phase 8.
 *
 * Run: bun src/tui/_ctx_demo.ts
 */

await import('@opentui/solid/preload');
const { mountCtxDemo } = await import('./_ctx_demo_app');
mountCtxDemo();
