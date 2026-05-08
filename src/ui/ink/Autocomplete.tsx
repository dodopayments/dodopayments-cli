import React from 'react';
import { Box, Text } from 'ink';
import { COMMANDS } from '../../lib/commands';

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

  // Exact match hides autocomplete
  if (suggestions.length === 1 && suggestions[0]!.command.toLowerCase() === input.toLowerCase()) {
    return null;
  }

  return (
    <Box flexDirection="column" paddingX={1} borderStyle="round" borderColor="#07BC70" width="100%">
      {suggestions.map((cmd, i) => (
        <Box key={cmd.command}>
          <Text color={i === selectedIndex ? '#07BC70' : 'white'}>
            {i === selectedIndex ? '❯ ' : '  '}
            {cmd.command}
          </Text>
          {cmd.command.endsWith('list') && (
            <Text color="gray"> {'<page>'}</Text>
          )}
          {(cmd.command.endsWith('info') || cmd.command.endsWith('update') || cmd.command.endsWith('delete')) && (
            <Text color="gray"> {'<id>'}</Text>
          )}
          <Text color="gray"> — {cmd.description}</Text>
        </Box>
      ))}
    </Box>
  );
};

export const getSuggestions = (input: string) => {
  if (!input.startsWith('/')) return [];
  return COMMANDS.filter((c) =>
    c.command.toLowerCase().startsWith(input.toLowerCase()),
  );
};
