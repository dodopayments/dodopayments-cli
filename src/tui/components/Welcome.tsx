import { Show } from 'solid-js';
import { useTerminalDimensions } from '@opentui/solid';
import { colors } from '../theme';
import { TipLine } from './TipLine';

const PAYMENTS_MIN_WIDTH = 84;

export const Welcome = () => {
  const dims = useTerminalDimensions();
  const showPayments = () => dims().width >= PAYMENTS_MIN_WIDTH;

  return (
    <box
      flexGrow={1}
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      paddingTop={2}
      paddingBottom={2}
    >
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
      <box flexShrink={0} paddingTop={2}>
        <TipLine />
      </box>
    </box>
  );
};
