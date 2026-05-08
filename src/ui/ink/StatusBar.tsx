import React, { useEffect, useState } from 'react';
import { Box, Text } from 'ink';
import { resolveCredentials } from '../../utils/auth';

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

  if (!authInfo) {
    return (
      <Box paddingX={1} paddingBottom={1}>
        <Text color="gray">Not logged in. Run </Text>
        <Text color="#07BC70">/login</Text>
      </Box>
    );
  }

  const isTest = authInfo.mode === 'test_mode';
  return (
    <Box paddingX={1} paddingBottom={1}>
      <Text color={isTest ? 'yellow' : '#07BC70'}>◉ {authInfo.mode}</Text>
      <Text color="gray">  {authInfo.key} using stored credentials</Text>
    </Box>
  );
};
