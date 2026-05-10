import { For, Show } from 'solid-js';
import { colors } from '../theme';
import { useTui } from '../context';

interface HintPair {
  key: string;
  label: string;
}

const HINTS_DEFAULT: HintPair[] = [
  { key: 'tab', label: 'complete' },
  { key: '/', label: 'commands' },
  { key: 'ctrl+c', label: 'exit' },
];

const HINTS_PALETTE: HintPair[] = [
  { key: '↑↓', label: 'navigate' },
  { key: '↵', label: 'select' },
  { key: 'esc', label: 'cancel' },
];

const HINTS_PROCESSING: HintPair[] = [{ key: 'ctrl+c', label: 'cancel' }];

export const HintRow = () => {
  const { paletteVisible, isProcessing } = useTui();
  const hints = () => {
    if (paletteVisible()) return HINTS_PALETTE;
    if (isProcessing()) return HINTS_PROCESSING;
    return HINTS_DEFAULT;
  };

  return (
    <box flexDirection="row" justifyContent="flex-end" paddingRight={2} flexShrink={0}>
      <For each={hints()}>
        {(h, i) => (
          <box flexDirection="row">
            <Show when={i() > 0}>
              <text fg={colors.textDim}>{'   '}</text>
            </Show>
            <text fg={colors.textPrimary}>{h.key}</text>
            <text fg={colors.textDim}>{` ${h.label}`}</text>
          </box>
        )}
      </For>
    </box>
  );
};
