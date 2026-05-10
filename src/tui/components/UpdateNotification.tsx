import { Show } from 'solid-js';
import { colors, glyphs } from '../theme';
import type { InstallMethod, UpdateInfo } from '../../utils/update';

interface UpdateNotificationProps {
  info: UpdateInfo;
  method: InstallMethod;
}

export const UpdateNotification = (props: UpdateNotificationProps) => {
  const isMajor = () => props.info.delta === 'major';
  const canSelfUpdate = () => props.method === 'npm' || props.method === 'bun';

  return (
    <box
      borderStyle="rounded"
      borderColor={colors.warning}
      paddingLeft={2}
      paddingRight={2}
      paddingTop={1}
      paddingBottom={1}
      flexDirection="column"
    >
      <box flexDirection="row">
        <text fg={colors.warning} attributes={1}>
          {`${glyphs.bullet} ${isMajor() ? 'Major update available' : 'Update available'} `}
        </text>
        <text fg={colors.textMuted}>{`v${props.info.currentVersion} `}</text>
        <text fg={colors.textDim}>{`${glyphs.arrow} `}</text>
        <text fg={colors.brand} attributes={1}>{`v${props.info.latestVersion}`}</text>
      </box>
      <Show
        when={canSelfUpdate()}
        fallback={
          <box flexDirection="row">
            <text fg={colors.textMuted}>Refer to </text>
            <text fg={colors.info} attributes={5}>
              https://github.com/dodopayments/dodopayments-cli
            </text>
            <text fg={colors.textMuted}> for the update guide.</text>
          </box>
        }
      >
        <box flexDirection="row">
          <text fg={colors.textMuted}>Run </text>
          <text fg={colors.accentLime} attributes={1}>/update</text>
          <text fg={colors.textMuted}> to install now.</text>
        </box>
      </Show>
    </box>
  );
};
