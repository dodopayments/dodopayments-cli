import { useTerminalDimensions } from '@opentui/solid';
import { colors, glyphs } from '../theme';
import { useTui } from '../context';

interface InputBarProps {
  onSubmit: (value: string) => void;
  onInput?: (value: string) => void;
  value?: string;
  placeholder?: string;
  disabled?: boolean;
}

export const InputBar = (props: InputBarProps) => {
  const dims = useTerminalDimensions();
  const dividerWidth = () => Math.max(0, dims().width - 2);
  const { promptActive } = useTui();
  const isFocused = () => !props.disabled && !promptActive();

  const handleInput = (next: string) => {
    props.onInput?.(next);
  };

  const handleSubmit = (final: string) => {
    const trimmed = final.trim();
    if (!trimmed) return;
    props.onSubmit(trimmed);
  };

  return (
    <box flexDirection="column" flexShrink={0}>
      <box paddingLeft={1} paddingRight={1} flexShrink={0}>
        <text fg={colors.textDim}>{'─'.repeat(dividerWidth())}</text>
      </box>
      <box
        flexDirection="row"
        paddingLeft={1}
        paddingRight={1}
        backgroundColor={colors.brandBlack}
      >
        <text fg={colors.accentLime}>{`${glyphs.prompt} `}</text>
        <input
          flexGrow={1}
          focused={isFocused()}
          value={props.value ?? ''}
          placeholder={props.placeholder ?? 'Type a command. /help to list all.'}
          onInput={handleInput}
          onSubmit={handleSubmit}
        />
      </box>
    </box>
  );
};
