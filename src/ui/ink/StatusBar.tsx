import React, { useEffect, useState } from 'react';
import { Box, Text } from 'ink';
import { resolveCredentials } from '../../utils/auth';
import { colors, glyphs } from '../theme';
import { version } from '../../../package.json';

interface StatusBarProps {
  authInfo?: { mode: string; key: string } | null;
}

export const StatusBar = ({ authInfo: propAuthInfo }: StatusBarProps) => {
  const [authInfo, setAuthInfo] = useState<{ mode: string; key: string } | null>(null);

  useEffect(() => {
    if (propAuthInfo !== undefined) {
      setAuthInfo(propAuthInfo);
      return;
    }

    resolveCredentials(undefined, false)
      .then(({ mode, apiKey }) => {
        const maskedKey = apiKey.slice(0, 8) + '…' + apiKey.slice(-3);
        setAuthInfo({ mode, key: maskedKey });
      })
      .catch(() => setAuthInfo(null));
  }, [propAuthInfo]);

  const sep = ` ${glyphs.separator} `;

  if (!authInfo) {
    return (
      <Box paddingX={1}>
        <Box flexGrow={1}>
          <Text color={colors.textMuted}>not signed in</Text>
          <Text color={colors.textDim}>{sep}</Text>
          <Text color={colors.accentLime}>/login</Text>
        </Box>
        <Text color={colors.textDim}>v{version}</Text>
      </Box>
    );
  }

  const isTest = authInfo.mode === 'test_mode';
  const modeLabel = isTest ? 'TEST' : 'LIVE';
  const modeColor = isTest ? colors.testMode : colors.liveMode;

  return (
    <Box paddingX={1}>
      <Box flexGrow={1}>
        <Text color={modeColor}>{modeLabel}</Text>
        <Text color={colors.textDim}>{sep}</Text>
        <Text color={colors.textMuted}>{authInfo.key}</Text>
      </Box>
      <Text color={colors.textDim}>v{version}</Text>
    </Box>
  );
};
