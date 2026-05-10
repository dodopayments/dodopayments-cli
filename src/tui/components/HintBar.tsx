import { colors } from '../theme';
import { useTui } from '../context';

const HINT_PALETTE_OPEN = '↑↓ navigate  ·  ↵ select  ·  esc cancel';
const HINT_PROCESSING = 'ctrl+c cancel  ·  please wait…';
const HINT_INPUT_EMPTY = '/ palette  ·  ↑↓ history  ·  ctrl+c exit';
const HINT_INPUT_TYPING = '↵ submit  ·  backspace edit  ·  ctrl+c exit';

export const HintBar = () => {
  const { input, paletteVisible, isProcessing } = useTui();
  const text = () => {
    if (paletteVisible()) return HINT_PALETTE_OPEN;
    if (isProcessing()) return HINT_PROCESSING;
    if (input().length === 0) return HINT_INPUT_EMPTY;
    return HINT_INPUT_TYPING;
  };

  return (
    <box paddingLeft={1} flexShrink={0}>
      <text fg={colors.textDim}>{text()}</text>
    </box>
  );
};
