import React, { useEffect, useState } from 'react';
import { Box, Text } from 'ink';
import { resolveCredentials } from '../../utils/auth';
import { colors, glyphs } from '../theme';

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
        const maskedKey = apiKey.slice(0, 10) + '...' + apiKey.slice(-3);
        setAuthInfo({ mode, key: maskedKey });
      })
      .catch(() => setAuthInfo(null));
  }, [propAuthInfo]);

  const sep = ` ${glyphs.separator} `;

  if (!authInfo) {
    return (
      <Box paddingX={1} paddingBottom={1}>
        <Text color={colors.textDim}>{glyphs.separator} </Text>
        <Text color={colors.textMuted}>Not signed in. Run </Text>
        <Text color={colors.brand}>/login</Text>
        <Text color={colors.textMuted}> to get started</Text>
      </Box>
    );
  }

  const isTest = authInfo.mode === 'test_mode';
  const dotColor = isTest ? colors.testMode : colors.liveMode;
  const label = isTest ? 'TEST MODE' : 'LIVE MODE';

  return (
    <Box paddingX={1} paddingBottom={1}>
      <Text color={dotColor} bold>{glyphs.dot} {label}</Text>
      <Text color={colors.textDim}>{sep}</Text>
      <Text color={colors.textMuted}>{authInfo.key}</Text>
      <Text color={colors.textDim}>{sep}</Text>
      <Text color={colors.textMuted}>stored credentials</Text>
    </Box>
  );
};
