/**
 * OpenTUI bootstrap. Installs the @opentui/solid babel plugin BEFORE any
 * Solid TSX module is imported, then dynamically loads the TSX app tree.
 *
 * This file MUST stay JSX-free and MUST NOT statically import any TSX
 * module. Doing so would load (and React-transform) the Solid tree before
 * the plugin is registered, which is exactly the bug the bootstrap exists
 * to prevent. See plan v2 Phase 1 + Oracle consult bg_73e90d73.
 */

export const mountTui = async (): Promise<void> => {
  await import('@opentui/solid/preload');
  const { mountTuiApp } = await import('./app');
  mountTuiApp();
};
