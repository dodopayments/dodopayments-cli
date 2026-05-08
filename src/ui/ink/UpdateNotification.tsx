import React from 'react';
import { Box, Text } from 'ink';

interface UpdateNotificationProps {
  latestVersion: string;
}

export const UpdateNotification: React.FC<UpdateNotificationProps> = ({ latestVersion }) => {
  return (
    <Box
      borderStyle="round"
      borderColor="yellow"
      paddingX={2}
      marginY={1}
      flexDirection="column"
    >
      <Text>
        <Text color="yellow" bold>Update Available! </Text>
        <Text>A new version of Dodo Payments CLI is available: </Text>
        <Text color="green" bold>{latestVersion}</Text>
      </Text>
      <Box marginTop={1}>
        <Text>Refer to </Text>
        <Text color="cyan" bold>https://github.com/dodopayments/dodopayments-cli</Text>
        <Text> for the update guide.</Text>
      </Box>
    </Box>
  );
};
