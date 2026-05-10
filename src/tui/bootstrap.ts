/**
 * Async entry into the OpenTUI render tree. Kept separate from `app.tsx`
 * so `src/index.ts` can lazy-load the TUI only on TTY launches and keep
 * the headless dispatch path free of the Solid render graph.
 *
 * The Solid babel plugin is installed via bunfig.toml's preload (Ink is
 * gone, so the global preload no longer collides with anything).
 */

export const mountTui = async (): Promise<void> => {
  const { mountTuiApp } = await import('./app');
  mountTuiApp();
};
