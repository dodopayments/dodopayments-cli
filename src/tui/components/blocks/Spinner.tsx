import { createSignal, onCleanup, onMount } from 'solid-js';
import { colors } from '../../theme';

const FRAMES = ['⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'];

export const Spinner = (props: { label: string }) => {
  const [frame, setFrame] = createSignal(0);
  onMount(() => {
    const t = setInterval(() => setFrame((f) => (f + 1) % FRAMES.length), 80);
    onCleanup(() => clearInterval(t));
  });
  return (
    <box flexDirection="row">
      <text fg={colors.accentAmber}>{FRAMES[frame()] ?? FRAMES[0]}</text>
      <text fg={colors.textMuted}>{` ${props.label}`}</text>
    </box>
  );
};
