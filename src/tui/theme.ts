/**
 * Design tokens for the Dodo Payments TUI. Sole source of truth for brand
 * colors, glyphs, and the ASCII wordmark. Brand values from
 * dodopayments.com/brand. Do not introduce hex literals or named ANSI
 * colors elsewhere -- everything must reference these tokens.
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

export const helpHeadingColors: Record<string, string> = {
  PRODUCTS: colors.accentCyan,
  PAYMENTS: colors.accentAmber,
  CUSTOMERS: colors.accentMagenta,
  DISCOUNTS: colors.accentLime,
  LICENCES: colors.accentSky,
  ADDONS: colors.accentCyan,
  REFUNDS: colors.accentAmber,
  CHECKOUT: colors.accentLime,
  WEBHOOKS: colors.accentMagenta,
  AI: colors.accentLime,
  AUTH: colors.accentSky,
  SESSION: colors.textMuted,
};


