import { For, createMemo } from 'solid-js';
import { colors } from '../../theme';
import {
  computeColumnWidths,
  isNumericColumn,
  truncateCell,
  MIN_COL_WIDTH,
} from './tableLogic';

export const Table = (props: { data: any[]; statusColumn?: string }) => {
  const keys = createMemo(() => {
    const first = props.data[0];
    return first ? Object.keys(first) : [];
  });
  const widths = createMemo(() => computeColumnWidths(keys(), props.data));
  const numericFlags = createMemo(() => {
    const flags: Record<string, boolean> = {};
    for (const k of keys()) {
      flags[k] = isNumericColumn(k, props.data.map((row) => row[k]));
    }
    return flags;
  });

  if (!props.data || props.data.length === 0) return null;

  return (
    <box flexDirection="column" borderStyle="single" borderColor={colors.accentSky}>
      <box paddingLeft={1} paddingRight={1} flexDirection="row">
        <For each={keys()}>
          {(k) => (
            <box width={widths()[k] ?? MIN_COL_WIDTH}>
              <text fg={colors.accentSky} attributes={1}>{k}</text>
            </box>
          )}
        </For>
      </box>
      <box paddingLeft={1} paddingRight={1} flexDirection="row">
        <For each={keys()}>
          {(k) => {
            const w = widths()[k] ?? MIN_COL_WIDTH;
            return (
              <box width={w}>
                <text fg={colors.textDim}>{'─'.repeat(Math.max(0, w - 1))}</text>
              </box>
            );
          }}
        </For>
      </box>
      <For each={props.data}>
        {(row, i) => {
          const baseColor = colors.textPrimary;
          return (
            <box paddingLeft={1} paddingRight={1} flexDirection="row">
              <For each={keys()}>
                {(k) => {
                  const w = widths()[k] ?? MIN_COL_WIDTH;
                  const valStr = String(row[k] ?? '');
                  const display = truncateCell(valStr, w);
                  const fg = numericFlags()[k] ? colors.accentAmber : baseColor;
                  return (
                    <box width={w}>
                      <text fg={fg}>{display}</text>
                    </box>
                  );
                }}
              </For>
            </box>
          );
        }}
      </For>
    </box>
  );
};
