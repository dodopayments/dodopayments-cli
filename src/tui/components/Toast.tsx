import { For, Show } from 'solid-js';
import { useToasts } from '../toast';
import { colors, glyphs } from '../theme';

const variantColor = {
  success: colors.success,
  info: colors.info,
  error: colors.error,
} as const;

const variantGlyph = {
  success: glyphs.check,
  info: glyphs.bullet,
  error: glyphs.cross,
} as const;

export const ToastViewport = () => (
  <Show when={useToasts().length > 0}>
    <box
      position="absolute"
      top={1}
      right={2}
      flexDirection="column"
      zIndex={1000}
    >
      <For each={useToasts()}>
        {(toast) => (
          <box
            borderStyle="rounded"
            borderColor={variantColor[toast.variant]}
            backgroundColor={colors.brandBlack}
            paddingLeft={1}
            paddingRight={1}
            flexDirection="row"
            marginBottom={1}
          >
            <text fg={variantColor[toast.variant]}>{`${variantGlyph[toast.variant]} `}</text>
            <text fg={colors.textPrimary}>{toast.message}</text>
          </box>
        )}
      </For>
    </box>
  </Show>
);
