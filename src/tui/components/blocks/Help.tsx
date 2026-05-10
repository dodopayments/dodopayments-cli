import { For } from 'solid-js';
import { colors, glyphs, helpHeadingColors } from '../../theme';
import { HELP_FOOTER, HELP_GROUPS } from '../../help-structure';

const HEADING_WIDTH = 12;
const COMMAND_WIDTH = 22;
const ARGS_WIDTH = 10;

export const Help = () => (
  <box flexDirection="column">
    <box paddingBottom={1} flexDirection="row">
      <text fg={colors.brandLime} attributes={1}>{`${glyphs.bullet} `}</text>
      <text fg={colors.textPrimary} attributes={1}>Dodo Payments CLI</text>
    </box>
    <For each={HELP_GROUPS}>
      {(group, gi) => {
        const headingColor = helpHeadingColors[group.heading] ?? colors.textMuted;
        const isLast = gi() === HELP_GROUPS.length - 1;
        return (
          <box flexDirection="column" paddingBottom={isLast ? 1 : 0}>
            <For each={group.items}>
              {(item, ii) => (
                <box flexDirection="row">
                  <box width={HEADING_WIDTH}>
                    <text fg={headingColor} attributes={1}>
                      {ii() === 0 ? group.heading : ''}
                    </text>
                  </box>
                  <box width={COMMAND_WIDTH}>
                    <text fg={colors.textPrimary}>{item.command}</text>
                  </box>
                  <box width={ARGS_WIDTH}>
                    <text fg={colors.textDim}>{item.args ?? ''}</text>
                  </box>
                  <text fg={colors.textMuted}>{item.description}</text>
                </box>
              )}
            </For>
          </box>
        );
      }}
    </For>
    <text fg={colors.textDim}>{HELP_FOOTER}</text>
  </box>
);
