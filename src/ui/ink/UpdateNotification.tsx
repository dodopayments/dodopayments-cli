import React from 'react';
import { Box, Text } from 'ink';
import { boxes, colors, glyphs } from '../theme';
import type { InstallMethod, UpdateInfo } from '../../utils/update';

interface UpdateNotificationProps {
  info: UpdateInfo;
  method: InstallMethod;
}

export const UpdateNotification: React.FC<UpdateNotificationProps> = ({ info, method }) => {
  const isMajor = info.delta === 'major';
  const canSelfUpdate = method === 'npm' || method === 'bun';

  return (
    <Box {...boxes.warning} paddingX={2} marginY={1} flexDirection="column">
      <Text>
        <Text color={colors.warning} bold>{glyphs.bullet} {isMajor ? 'Major update available' : 'Update available'} </Text>
        <Text color={colors.textMuted}>v{info.currentVersion} </Text>
        <Text color={colors.textDim}>{glyphs.arrow} </Text>
        <Text color={colors.brand} bold>v{info.latestVersion}</Text>
      </Text>
      <Box marginTop={1}>
        {canSelfUpdate ? (
          <Text>
            <Text color={colors.textMuted}>Run </Text>
            <Text color={colors.accentLime} bold>/update</Text>
            <Text color={colors.textMuted}> to install now.</Text>
          </Text>
        ) : (
          <Text>
            <Text color={colors.textMuted}>Refer to </Text>
            <Text color={colors.info} bold underline>https://github.com/dodopayments/dodopayments-cli</Text>
            <Text color={colors.textMuted}> for the update guide.</Text>
          </Text>
        )}
      </Box>
    </Box>
  );
};
