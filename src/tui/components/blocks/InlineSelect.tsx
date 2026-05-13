import { onCleanup, onMount, Show } from 'solid-js';
import { colors } from '../../theme';
import { useTui } from '../../context';

interface InlineSelectOption {
  label: string;
  value: string;
}

interface InlineSelectProps {
  label?: string;
  options: InlineSelectOption[];
  onSubmit: (value: string) => void;
}

export const InlineSelect = (props: InlineSelectProps) => {
  const { setPromptActive } = useTui();
  onMount(() => {
    setPromptActive(true);
    onCleanup(() => setPromptActive(false));
  });

  return (
    <box flexDirection="column">
      <Show when={props.label}>
        <text fg={colors.textMuted}>{props.label}</text>
      </Show>
      <select
        focused
        height={Math.min(props.options.length, 10)}
        showDescription={false}
        options={props.options.map((o) => ({ name: o.label, value: o.value, description: '' }))}
        onSelect={(_i, opt) => {
          if (opt) props.onSubmit(opt.value as string);
        }}
      />
    </box>
  );
};
