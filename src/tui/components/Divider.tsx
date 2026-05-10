import { useTerminalDimensions } from '@opentui/solid';
import { colors } from '../theme';

export const Divider = () => {
  const dims = useTerminalDimensions();
  const width = () => Math.max(0, dims().width - 2);
  return (
    <box paddingLeft={1} paddingRight={1} flexShrink={0}>
      <text fg={colors.textDim}>{'─'.repeat(width())}</text>
    </box>
  );
};
