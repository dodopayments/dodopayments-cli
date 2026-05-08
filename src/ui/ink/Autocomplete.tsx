import React from 'react';
import { Box, Text } from 'ink';
import { COMMANDS } from '../../lib/commands';
import { boxes, colors, glyphs } from '../theme';

interface AutocompleteProps {
  input: string;
  selectedIndex: number;
}

export const Autocomplete = ({ input, selectedIndex }: AutocompleteProps) => {
  if (!input.startsWith('/')) return null;

  const suggestions = COMMANDS.filter((c) =>
    c.command.toLowerCase().startsWith(input.toLowerCase()),
  );

  if (suggestions.length === 0) return null;

  if (suggestions.length === 1 && suggestions[0]!.command.toLowerCase() === input.toLowerCase()) {
    return null;
  }

  return (
    <Box flexDirection="column" paddingX={1} {...boxes.panel} width="100%">
      {suggestions.map((cmd, i) => {
        const isSelected = i === selectedIndex;
        const cmdColor = isSelected ? colors.brandLime : colors.textPrimary;
        const arrow = isSelected ? `${glyphs.prompt} ` : '  ';
        return (
          <Box key={cmd.command}>
            <Text color={cmdColor} bold={isSelected}>
              {arrow}
              {cmd.command}
            </Text>
            {cmd.command.endsWith('list') && (
              <Text color={colors.textDim}> {'<page>'}</Text>
            )}
            {(cmd.command.endsWith('info') || cmd.command.endsWith('update') || cmd.command.endsWith('delete')) && (
              <Text color={colors.textDim}> {'<id>'}</Text>
            )}
            <Text color={colors.textMuted}> {glyphs.separator} {cmd.description}</Text>
          </Box>
        );
      })}
      <Box justifyContent="flex-end">
        <Text color={colors.textDim}>Tab/{glyphs.arrow} to complete</Text>
      </Box>
    </Box>
  );
};

export const getSuggestions = (input: string) => {
  if (!input.startsWith('/')) return [];
  return COMMANDS.filter((c) =>
    c.command.toLowerCase().startsWith(input.toLowerCase()),
  );
};
