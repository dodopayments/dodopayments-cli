/**
 * Design tokens for the Dodo Payments OpenTUI/SolidJS layer.
 * Mirrors `src/ui/theme.ts` (the Ink theme) so visual parity is exact during
 * the dual-stack period (Phases 1-7). When the legacy Ink path is removed in
 * Phase 8, this file becomes the sole source of truth.
 *
 * Brand values from dodopayments.com/brand. Do not introduce hex literals or
 * named ANSI colors elsewhere.
 */

export const colors = {
  brand: '#07BC70',
  brandLime: '#C6FE1E',
  brandForest: '#003F28',
  brandBlack: '#1F2023',

  surfaceDeep: '#0D0D0D',
  border: '#212423',

  textPrimary: '#FFFFFF',
  textMuted: '#737470',
  textDim: '#535452',

  success: '#07BC70',
  error: '#EF4444',
  warning: '#F5A623',
  info: '#38BDF8',

  accentSky: '#38BDF8',
  accentAmber: '#F5A623',
  accentMagenta: '#E85BCF',
  accentCyan: '#7FC4D4',
  accentLime: '#C6FE1E',

  testMode: '#F5A623',
  liveMode: '#07BC70',
} as const;

export const glyphs = {
  prompt: '❯',
  bullet: '◆',
  dot: '●',
  check: '✓',
  cross: '✗',
  arrow: '→',
  separator: '·',
} as const;

export const spacing = { xs: 0, sm: 1, md: 2, lg: 3 } as const;

/**
 * Three-row ASCII wordmark for the welcome screen. Renders correctly at any
 * terminal width >= 24 cols (the wordmark itself is 19 cols + padding).
 */
export const LOGO_DODO = [
  '█▀▀▄ █▀▀█ █▀▀▄ █▀▀█',
  '█  █ █  █ █  █ █  █',
  '█▄▄▀ ▀▀▀▀ █▄▄▀ ▀▀▀▀',
] as const;
