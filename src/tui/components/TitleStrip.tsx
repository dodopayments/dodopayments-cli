import { colors } from '../theme';
import { useTui } from '../context';

export const TitleStrip = () => {
  const { lastCommand } = useTui();
  const label = () => lastCommand() ?? 'dodopayments cli';
  const fg = () => (lastCommand() ? colors.textPrimary : colors.textDim);

  return (
    <box
      flexDirection="row"
      flexShrink={0}
      paddingLeft={2}
      paddingRight={2}
      backgroundColor={colors.surfaceDeep}
    >
      <text fg={fg()}>{label()}</text>
    </box>
  );
};
