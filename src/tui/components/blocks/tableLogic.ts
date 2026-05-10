/**
 * Numeric / date column detection for Table and Detail blocks. Direct port
 * of `src/ui/ink/OutputBlock.tsx` regex heuristics + isNumericColumn() so
 * column highlighting (amber for numbers) stays consistent across the two
 * TUIs through the dual-stack period.
 */

export const MAX_COL_WIDTH = 45;
export const MIN_COL_WIDTH = 10;

export const NUMERIC_KEY_HINTS =
  /(amount|price|total|cost|value|count|quantity|cycles|percent|fee)/i;
export const DATE_KEY_HINTS = /(created|updated|deleted|timestamp|date|_at\b|on\b)/i;
export const CURRENCY_OR_NUMBER =
  /^[\$€£₹¥]\s*-?\d[\d,]*(\.\d+)?%?$|^-?\d[\d,]*(\.\d+)?%?$/;

export const isNumericColumn = (key: string, samples: any[]): boolean => {
  if (DATE_KEY_HINTS.test(key)) return false;
  if (NUMERIC_KEY_HINTS.test(key)) return true;
  let numericHits = 0;
  let nonEmpty = 0;
  for (const v of samples) {
    const s = String(v ?? '').trim().split(/\s/)[0] ?? '';
    if (!s) continue;
    nonEmpty++;
    if (CURRENCY_OR_NUMBER.test(s)) numericHits++;
  }
  return nonEmpty > 0 && numericHits / nonEmpty >= 0.7;
};

export const isNumericValue = (key: string, value: unknown): boolean => {
  const valStr = String(value);
  const firstToken = valStr.trim().split(/\s/)[0] ?? '';
  if (DATE_KEY_HINTS.test(key)) return false;
  return NUMERIC_KEY_HINTS.test(key) || CURRENCY_OR_NUMBER.test(firstToken);
};

export const computeColumnWidths = (
  keys: string[],
  data: any[],
): Record<string, number> => {
  const widths: Record<string, number> = {};
  for (const k of keys) {
    const maxLen = Math.max(
      k.length,
      ...data.map((row) => String(row[k] ?? '').length),
    );
    widths[k] = Math.min(Math.max(maxLen + 2, MIN_COL_WIDTH), MAX_COL_WIDTH);
  }
  return widths;
};

export const truncateCell = (value: string, width: number): string =>
  value.length > width - 1
    ? value.substring(0, Math.max(0, width - 2)) + '…'
    : value;
