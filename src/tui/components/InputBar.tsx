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
    <box
      flexDirection="row"
      paddingLeft={1}
      paddingRight={1}
      flexShrink={0}
      backgroundColor={colors.brandBlack}
    >
      <text fg={colors.accentLime}>{`${glyphs.prompt} `}</text>
      <input
        flexGrow={1}
        focused={isFocused()}
        value={props.value ?? ''}
        placeholder={props.placeholder ?? 'Type a command. /help to list all.'}
        onInput={handleInput as any}
        onSubmit={handleSubmit as any}
      />
    </box>
  );
};
