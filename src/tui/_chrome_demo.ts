/**
 * Phase 2 throwaway: bootstrap entry for the chrome demo. Mirrors the
 * production bootstrap.ts pattern -- JSX-free + dynamic import after preload.
 * Deleted in Phase 8.
 *
 * Run: bun src/tui/_chrome_demo.ts
 */

await import('@opentui/solid/preload');
const { mountChromeDemo } = await import('./_chrome_demo_app');
mountChromeDemo();
