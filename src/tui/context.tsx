/**
 * Solid context exposing the shared TUI signals to deeply-nested components
 * (StatusBar, InputBar, HintBar, blocks). Replaces the prop-drilling pattern
 * the Ink TUI used. Created in Phase 2; gets `messages` + `addBlock` plumbing
 * in Phase 3 when CommandContext lands.
 */

import { createContext, useContext, type Accessor } from 'solid-js';

export type AuthInfo = { mode: 'test_mode' | 'live_mode'; key: string } | null;

export interface TuiContextValue {
  authInfo: Accessor<AuthInfo>;
  input: Accessor<string>;
  paletteVisible: Accessor<boolean>;
  isProcessing: Accessor<boolean>;
}

const TuiContext = createContext<TuiContextValue>();

export const TuiContextProvider = TuiContext.Provider;

export const useTui = (): TuiContextValue => {
  const ctx = useContext(TuiContext);
  if (!ctx) throw new Error('useTui must be called inside <TuiContextProvider>');
  return ctx;
};
