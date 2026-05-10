import { Show } from 'solid-js';
import { useTerminalDimensions } from '@opentui/solid';
import { colors, glyphs } from '../theme';
import { version } from '../../../package.json';

const PAYMENTS_MIN_WIDTH = 84;

export const Welcome = () => {
  const dims = useTerminalDimensions();
  const showPayments = () => dims().width >= PAYMENTS_MIN_WIDTH;

  return (
    <box flexDirection="column" paddingLeft={2} paddingTop={1} flexShrink={0}>
      <ascii_font
        text="DODO"
        font="block"
        color={[colors.brand, colors.brandLime]}
        selectable={false}
      />
      <Show when={showPayments()}>
        <ascii_font
          text="PAYMENTS"
          font="block"
          color={[colors.brandLime, colors.brand]}
          selectable={false}
        />
      </Show>
      <box flexDirection="row" paddingTop={1}>
        <text fg={colors.textPrimary}>Welcome to </text>
        <text fg={colors.brand} attributes={1}>Dodo Payments</text>
        <text fg={colors.textPrimary}>! Let’s get you set up.</text>
      </box>
      <text fg={colors.textMuted}>{`dodopayments cli  ${glyphs.separator}  v${version}`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} type ${glyphs.prompt} /help to see commands`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} type ${glyphs.prompt} /login to authenticate`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} press ctrl+c to exit`}</text>
    </box>
  );
};
