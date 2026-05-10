import React from 'react';
import { Box, Text } from 'ink';
import { colors, glyphs } from '../theme';

interface Hint {
  key: string;
  label: string;
}

interface HintBarProps {
  paletteVisible?: boolean;
  isProcessing?: boolean;
  inputEmpty?: boolean;
}

const sep = ` ${glyphs.separator} `;

const renderHints = (hints: Hint[]) => (
  <>
    {hints.map((h, i) => (
      <React.Fragment key={h.key}>
        {i > 0 && <Text color={colors.textDim}>{sep}</Text>}
        <Text color={colors.accentLime}>{h.key}</Text>
        <Text color={colors.textDim}> {h.label}</Text>
      </React.Fragment>
    ))}
  </>
);

export const HintBar = ({ paletteVisible, isProcessing, inputEmpty }: HintBarProps) => {
  const hints: Hint[] = paletteVisible
    ? [
        { key: '↑↓', label: 'navigate' },
        { key: '↵', label: 'select' },
        { key: 'tab', label: 'autocomplete' },
        { key: 'esc', label: 'cancel' },
      ]
    : isProcessing
    ? [{ key: 'esc', label: 'cancel' }]
    : inputEmpty
    ? [
        { key: '/', label: 'palette' },
        { key: '/help', label: 'commands' },
        { key: '↑↓', label: 'history' },
        { key: 'ctrl+c', label: 'exit' },
      ]
    : [
        { key: '↵', label: 'submit' },
        { key: 'tab', label: 'autocomplete' },
        { key: 'ctrl+l', label: 'clear' },
      ];

  return (
    <Box paddingX={1}>
      {renderHints(hints)}
    </Box>
  );
};
