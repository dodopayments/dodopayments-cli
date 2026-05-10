import { createSignal, onCleanup, onMount } from 'solid-js';
import { colors } from '../../theme';
import { useTui } from '../../context';

interface InlineInputProps {
  label: string;
  secure?: boolean;
  onSubmit: (value: string) => void;
}

export const InlineInput = (props: InlineInputProps) => {
  const [value, setValue] = createSignal('');
  const { setPromptActive } = useTui();

  onMount(() => {
    setPromptActive(true);
    onCleanup(() => setPromptActive(false));
  });

  return (
    <box
      borderStyle="rounded"
      borderColor={colors.accentMagenta}
      paddingLeft={1}
      paddingRight={1}
      flexDirection="column"
    >
      <text fg={colors.accentMagenta}>{props.label}</text>
      <input
        focused
        value={value()}
        placeholder={props.secure ? '••••••' : ''}
        onInput={((v: string) => setValue(v)) as any}
        onSubmit={((v: string) => props.onSubmit(v)) as any}
      />
    </box>
  );
};
