import { For, Show } from 'solid-js';
import type { Message } from '../types';
import { colors, glyphs } from '../theme';
import { renderBlock } from './blocks';

export const MessageRow = (props: { message: Message }) => (
  <box flexDirection="column" paddingBottom={1}>
    <Show when={props.message.role === 'user' && props.message.text}>
      <text fg={colors.textMuted}>{`${glyphs.prompt} ${props.message.text}`}</text>
    </Show>
    <For each={props.message.blocks}>{(b) => renderBlock(b)}</For>
  </box>
);
