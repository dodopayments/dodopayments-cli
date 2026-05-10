import React from 'react';
import path from 'node:path';
import os from 'node:os';
import { Box, Text } from 'ink';
import { colors, glyphs } from '../theme';
import { version } from '../../../package.json';

interface WelcomeBannerProps {
  authMode?: string | null;
}

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
  const sep = ` ${glyphs.separator} `;
  return (
    <Box flexDirection="column" alignItems="center" paddingY={1}>
      <Text color={colors.brand} bold>dodopayments</Text>
      <Box marginTop={1}>
        <Text color={colors.textMuted}>/help for commands</Text>
        <Text color={colors.textDim}>{sep}</Text>
        <Text color={colors.textMuted}>/login to auth</Text>
      </Box>
      <Box marginTop={1}>
        <Text color={colors.textDim}>v{version}</Text>
        <Text color={colors.textDim}>{sep}</Text>
        <Text color={modeColor(authMode)}>{formatMode(authMode)}</Text>
        <Text color={colors.textDim}>{sep}</Text>
        <Text color={colors.textDim}>{formatCwd()}</Text>
      </Box>
    </Box>
  );
};
