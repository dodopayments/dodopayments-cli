import { Show } from 'solid-js';
import { colors, glyphs } from '../theme';
import { version } from '../../../package.json';
import { useTui, type AuthInfo } from '../context';

export const StatusBar = () => {
  const { authInfo } = useTui();
  const sep = ` ${glyphs.separator} `;

  return (
    <box paddingLeft={1} paddingRight={1} flexDirection="row" flexShrink={0}>
      <box flexGrow={1} flexDirection="row">
        <Show
          when={authInfo()}
          fallback={
            <>
              <text fg={colors.textDim}>{`${glyphs.separator} `}</text>
              <text fg={colors.textMuted}>Not signed in. Run </text>
              <text fg={colors.accentLime}>/login</text>
              <text fg={colors.textMuted}> to get started</text>
            </>
          }
        >
          {(info: () => NonNullable<AuthInfo>) => {
            const isTest = info().mode === 'test_mode';
            const dotColor = isTest ? colors.testMode : colors.liveMode;
            const label = isTest ? 'TEST MODE' : 'LIVE MODE';
            return (
              <>
                <text fg={dotColor} attributes={1}>{`${glyphs.dot} ${label}`}</text>
                <text fg={colors.textDim}>{sep}</text>
                <text fg={colors.textMuted}>{info().key}</text>
                <text fg={colors.textDim}>{sep}</text>
                <text fg={colors.textMuted}>stored credentials</text>
              </>
            );
          }}
        </Show>
      </box>
      <text fg={colors.textDim}>{`v${version}`}</text>
    </box>
  );
};
