import { For, Show } from 'solid-js';
import type { Message } from '../types';
import { colors, glyphs } from '../theme';
import { renderBlock } from './blocks';

export const MessageRow = (props: { message: Message }) => (
  <box flexDirection="column" paddingBottom={1}>
    <Show when={props.message.role === 'user' && props.message.text}>
      <box flexDirection="row" alignItems="stretch" flexShrink={0} paddingBottom={1}>
        <box width={1} flexShrink={0} backgroundColor={colors.accentLime}>
          <text fg={colors.accentLime}> </text>
        </box>
        <box
          flexDirection="row"
          paddingLeft={2}
          paddingRight={2}
          backgroundColor={colors.brandBlack}
          flexShrink={0}
        >
          <text fg={colors.accentLime}>{`${glyphs.prompt} `}</text>
          <text fg={colors.textPrimary}>{props.message.text}</text>
        </box>
      </box>
    </Show>
    <For each={props.message.blocks}>{(b) => renderBlock(b)}</For>
  </box>
);
