import React from 'react';
import { Box, Text, useStdout } from 'ink';
import Gradient from 'ink-gradient';
import { renderWordmark } from '../wordmark';
import { colors } from '../theme';

const DODO = renderWordmark('DODO');
const PAYMENTS = renderWordmark('PAYMENTS');
const PAYMENTS_WIDTH = (PAYMENTS[0]?.length ?? 0) + 6;

export const WelcomeBanner = () => {
  const { stdout } = useStdout();
  const cols = stdout?.columns ?? 80;
  const showPayments = cols >= PAYMENTS_WIDTH;

  return (
    <Box flexDirection="column" paddingX={1}>
      <Gradient colors={[colors.brand, colors.brandLime]}>
        <Text>{DODO.join('\n')}</Text>
      </Gradient>
      {showPayments && (
        <Gradient colors={[colors.brandLime, colors.brand]}>
          <Text>{PAYMENTS.join('\n')}</Text>
        </Gradient>
      )}
      <Box marginTop={1}>
        <Text color={colors.textPrimary}>Welcome to </Text>
        <Text color={colors.brand} bold>Dodo Payments</Text>
        <Text color={colors.textPrimary}>! Let’s get you set up.</Text>
      </Box>
    </Box>
  );
};
