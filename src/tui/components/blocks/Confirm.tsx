import { onCleanup, onMount } from 'solid-js';
import { colors } from '../../theme';
import { useTui } from '../../context';

interface ConfirmProps {
  message: string;
  onConfirm: () => void;
  onCancel: () => void;
}

export const Confirm = (props: ConfirmProps) => {
  const { setPromptActive } = useTui();
  onMount(() => {
    setPromptActive(true);
    onCleanup(() => setPromptActive(false));
  });

  return (
    <box flexDirection="column">
      <text fg={colors.textPrimary}>{props.message}</text>
      <select
        focused
        options={[
          { name: 'Yes', value: 'yes', description: 'confirm' },
          { name: 'No', value: 'no', description: 'cancel' },
        ]}
        onSelect={(_i, opt) => {
          if (opt?.value === 'yes') props.onConfirm();
          else props.onCancel();
        }}
      />
    </box>
  );
};
