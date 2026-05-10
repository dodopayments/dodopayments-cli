import { createSignal } from 'solid-js';
import { useTerminalDimensions } from '@opentui/solid';
import { colors, glyphs } from '../theme';

interface InputBarProps {
  onSubmit: (value: string) => void;
  onInput?: (value: string) => void;
  placeholder?: string;
  disabled?: boolean;
}

export const InputBar = (props: InputBarProps) => {
  const [value, setValue] = createSignal('');
  const dims = useTerminalDimensions();
  const dividerWidth = () => Math.max(0, dims().width - 2);

  const handleInput = (next: string) => {
    setValue(next);
    props.onInput?.(next);
  };

  const handleSubmit = (final: string) => {
    const trimmed = final.trim();
    if (!trimmed) return;
    props.onSubmit(trimmed);
    setValue('');
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
          focused={!props.disabled}
          value={value()}
          placeholder={props.placeholder ?? 'Type a command. /help to list all.'}
          onInput={handleInput}
          onSubmit={handleSubmit}
        />
      </box>
    </box>
  );
};
