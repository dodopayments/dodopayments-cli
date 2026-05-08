import React from 'react';
import { Box, Text } from 'ink';
import { boxes, colors, glyphs } from '../theme';

interface UpdateNotificationProps {
  latestVersion: string;
}

export const UpdateNotification: React.FC<UpdateNotificationProps> = ({ latestVersion }) => {
  return (
    <Box {...boxes.warning} paddingX={2} marginY={1} flexDirection="column">
      <Text>
        <Text color={colors.warning} bold>{glyphs.bullet} Update available </Text>
        <Text color={colors.textMuted}>— a new version is available: </Text>
        <Text color={colors.brand} bold>{latestVersion}</Text>
      </Text>
      <Box marginTop={1}>
        <Text color={colors.textMuted}>Refer to </Text>
        <Text color={colors.info} bold underline>https://github.com/dodopayments/dodopayments-cli</Text>
        <Text color={colors.textMuted}> for the update guide.</Text>
      </Box>
    </Box>
  );
};
