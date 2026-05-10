import React from 'react';
import { Box, Text, useStdout } from 'ink';
import { COMMANDS } from '../../lib/commands';
import { boxes, colors, glyphs } from '../theme';

const PALETTE_WIDTH = 60;
const PALETTE_MAX_ITEMS = 8;

type Suggestion = (typeof COMMANDS)[number] & { score: number };

const score = (cmd: string, query: string): number => {
  if (!query || query === '/') return 1;
  const c = cmd.toLowerCase();
  const q = query.toLowerCase();
  if (c.startsWith(q)) return 1000 - (c.length - q.length);
  const idx = c.indexOf(q);
  if (idx !== -1) return 500 - idx;
  return -1;
};

const argHint = (cmd: string): string | null => {
  if (cmd.endsWith(' list')) return '<page>';
  if (cmd.endsWith(' info') || cmd.endsWith(' update') || cmd.endsWith(' delete')) return '<id>';
  return null;
};

export const getSuggestions = (input: string): Suggestion[] => {
  if (!input.startsWith('/')) return [];
  const ranked = COMMANDS
    .map((c) => ({ ...c, score: score(c.command, input) }))
    .filter((c) => c.score >= 0)
    .sort((a, b) => b.score - a.score);
  return ranked;
};

interface AutocompleteProps {
  input: string;
  selectedIndex: number;
}

export const Autocomplete = ({ input, selectedIndex }: AutocompleteProps) => {
  const { stdout } = useStdout();
  const cols = stdout?.columns ?? 80;

  if (!input.startsWith('/')) return null;
  const all = getSuggestions(input);
  if (all.length === 0) return null;
  if (all.length === 1 && all[0]!.command.toLowerCase() === input.toLowerCase()) return null;

  const visible = all.slice(0, PALETTE_MAX_ITEMS);
  const overflow = all.length - visible.length;
  const leftPad = Math.max(0, Math.floor((cols - PALETTE_WIDTH) / 2));

  return (
    <Box marginLeft={leftPad}>
      <Box flexDirection="column" {...boxes.panel} width={PALETTE_WIDTH} paddingX={1}>
        <Box>
          <Text color={colors.accentLime}>{glyphs.prompt} </Text>
          <Text color={colors.textPrimary}>{input}</Text>
        </Box>
        <Box>
          <Text color={colors.textDim}>{'─'.repeat(PALETTE_WIDTH - 4)}</Text>
        </Box>
        {visible.map((cmd, i) => {
          const isSelected = i === selectedIndex;
          const cmdColor = isSelected ? colors.accentLime : colors.textPrimary;
          const arrow = isSelected ? `${glyphs.prompt} ` : '  ';
          const hint = argHint(cmd.command);
          return (
            <Box key={cmd.command}>
              <Text color={cmdColor} bold={isSelected}>
                {arrow}{cmd.command}
              </Text>
              {hint && <Text color={colors.textDim}> {hint}</Text>}
              <Text color={colors.textMuted}> {glyphs.separator} {cmd.description}</Text>
            </Box>
          );
        })}
        {overflow > 0 && (
          <Box>
            <Text color={colors.textDim}>  +{overflow} more</Text>
          </Box>
        )}
        <Box marginTop={1} justifyContent="space-between">
          <Text color={colors.textDim}>↑↓ navigate</Text>
          <Text color={colors.textDim}>↵ select {glyphs.separator} esc cancel</Text>
        </Box>
      </Box>
    </Box>
  );
};
