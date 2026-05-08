import { Box, Text } from 'ink';
import { version } from '../../../package.json';
import { colors, glyphs } from '../theme';

export const WelcomeBanner = () => {
  return (
    <Box flexDirection="column" paddingX={1} marginBottom={1}>
      <Box flexDirection="row" justifyContent="space-between">
        <Box>
          <Text color={colors.brandLime} bold>{`  ${glyphs.bullet}  `}</Text>
          <Text color={colors.brand} bold wrap="truncate">Dodo Payments</Text>
        </Box>
        <Text color={colors.textDim}>v{version}</Text>
      </Box>
      <Box>
        <Text color={colors.textMuted}>     Billing & payments, from your terminal.</Text>
      </Box>
    </Box>
  );
};
