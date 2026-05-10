import { For } from 'solid-js';
import { colors, glyphs } from '../../theme';
import { isNumericValue } from './tableLogic';

export const Detail = (props: { data: Record<string, any> }) => {
  const entries = () => Object.entries(props.data);
  return (
    <box flexDirection="column" paddingTop={1} paddingBottom={1}>
      <For each={entries()}>
        {([key, value]) => {
          const valStr = String(value);
          const numeric = isNumericValue(key, value);
          const valueColor = numeric ? colors.accentAmber : colors.textPrimary;
          return (
            <box flexDirection="row">
              <box width={20} justifyContent="flex-end" paddingRight={1}>
                <text fg={colors.textMuted}>{key}</text>
              </box>
              <text fg={colors.textDim}>{`${glyphs.separator} `}</text>
              <text fg={valueColor}>{valStr}</text>
            </box>
          );
        }}
      </For>
    </box>
  );
};
