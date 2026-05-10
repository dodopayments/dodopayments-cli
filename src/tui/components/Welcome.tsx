/**
 * Welcome screen shown when the TUI mounts with no messages yet.
 * Layout mirrors opencode's left-aligned welcome: ASCII wordmark in brand
 * green, then dim metadata rows, then inline action hints.
 *
 * This is a pure render component with zero state. It disappears as soon as
 * `messages().length > 0` (gated by the parent App in Phase 2).
 */

import { For } from 'solid-js';
import { colors, glyphs, LOGO_DODO } from '../theme';
import { version } from '../../../package.json';

export const Welcome = () => {
  return (
    <box flexDirection="column" paddingTop={1} paddingLeft={2} flexShrink={0}>
      <For each={LOGO_DODO}>{(row) => <text fg={colors.brand}>{row}</text>}</For>
      <text> </text>
      <text fg={colors.textMuted}>{`dodopayments cli  ${glyphs.separator}  v${version}`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} type ${glyphs.prompt} /help to see commands`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} type ${glyphs.prompt} /login to authenticate`}</text>
      <text fg={colors.textDim}>{`${glyphs.separator} press ctrl+c to exit`}</text>
    </box>
  );
};
