/**
 * Wordmark renderer using vendored cfonts "block" font.
 * Font data: src/ui/block-font.json — sourced from cfonts (GPL-3.0,
 * Dominik Wilkowski). See src/ui/block-font.json for the original notice.
 */

import blockFont from './block-font.json';

interface BlockFont {
  lines: number;
  letterspace: string[];
  chars: Record<string, string[]>;
}

const FONT = blockFont as BlockFont;

const stripTags = (line: string): string => line.replace(/<\/?c[12]>/g, '');

export function renderWordmark(text: string): string[] {
  const upper = text.toUpperCase();
  const rows: string[] = Array.from({ length: FONT.lines }, () => '');
  for (const char of upper) {
    if (char === ' ') {
      for (let r = 0; r < FONT.lines; r++) rows[r] += '   ';
      continue;
    }
    const glyph = FONT.chars[char];
    if (!glyph) continue;
    const space = FONT.letterspace ?? Array(FONT.lines).fill(' ');
    for (let r = 0; r < FONT.lines; r++) {
      rows[r] += stripTags(glyph[r] ?? '') + (space[r] ?? '');
    }
  }
  return rows;
}

export const WORDMARK_ROWS = FONT.lines;
