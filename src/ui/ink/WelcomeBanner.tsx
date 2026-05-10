import React from 'react';
import path from 'node:path';
import os from 'node:os';
import { Box, Text } from 'ink';
import { colors, glyphs } from '../theme';
import { version } from '../../../package.json';

interface WelcomeBannerProps {
  authMode?: string | null;
}

const LOGO = [
  '█▀▀▄ █▀▀█ █▀▀▄ █▀▀█',
  '█  █ █  █ █  █ █  █',
  '█▄▄▀ ▀▀▀▀ █▄▄▀ ▀▀▀▀',
];

const formatCwd = (): string => {
  const cwd = process.cwd();
  const home = os.homedir();
  if (cwd === home) return '~';
  if (cwd.startsWith(home + path.sep)) return '~' + cwd.slice(home.length);
  return cwd;
};

const formatMode = (mode?: string | null): string => {
  if (!mode) return 'NOT SIGNED IN';
  return mode === 'test_mode' ? 'TEST' : 'LIVE';
};

const modeColor = (mode?: string | null): string => {
  if (!mode) return colors.textMuted;
  return mode === 'test_mode' ? colors.testMode : colors.liveMode;
};

export const WelcomeBanner = ({ authMode }: WelcomeBannerProps) => {
  return (
    <Box flexDirection="column" paddingX={2} paddingY={1}>
      <Box flexDirection="column">
        {LOGO.map((row, i) => (
          <Text key={i} color={colors.brand}>{row}</Text>
        ))}
      </Box>
      <Box marginTop={1} flexDirection="column">
        <Box>
          <Text color={colors.textMuted}>dodopayments cli</Text>
          <Text color={colors.textDim}> {glyphs.separator} v{version}</Text>
        </Box>
        <Box>
          <Text color={modeColor(authMode)}>{formatMode(authMode)}</Text>
          <Text color={colors.textDim}> {glyphs.separator} </Text>
          <Text color={colors.textMuted}>{formatCwd()}</Text>
        </Box>
      </Box>
      <Box marginTop={1} flexDirection="column">
        <Box>
          <Text color={colors.textDim}>type </Text>
          <Text color={colors.accentLime}>/help</Text>
          <Text color={colors.textDim}> to see commands</Text>
        </Box>
        <Box>
          <Text color={colors.textDim}>type </Text>
          <Text color={colors.accentLime}>/</Text>
          <Text color={colors.textDim}> to open the command palette</Text>
        </Box>
        {!authMode && (
          <Box>
            <Text color={colors.textDim}>type </Text>
            <Text color={colors.accentLime}>/login</Text>
            <Text color={colors.textDim}> to sign in</Text>
          </Box>
        )}
      </Box>
    </Box>
  );
};
