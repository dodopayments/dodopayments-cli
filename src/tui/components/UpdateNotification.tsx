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
      borderStyle="round"
      borderColor={colors.warning}
      paddingLeft={2}
      paddingRight={2}
      paddingTop={1}
      paddingBottom={1}
      flexDirection="column"
    >
      <text>
        <span fg={colors.warning} attributes={1}>
          {`${glyphs.bullet} ${isMajor() ? 'Major update available' : 'Update available'} `}
        </span>
        <span fg={colors.textMuted}>{`v${props.info.currentVersion} `}</span>
        <span fg={colors.textDim}>{`${glyphs.arrow} `}</span>
        <span fg={colors.brand} attributes={1}>{`v${props.info.latestVersion}`}</span>
      </text>
      <Show
        when={canSelfUpdate()}
        fallback={
          <text>
            <span fg={colors.textMuted}>Refer to </span>
            <a href="https://github.com/dodopayments/dodopayments-cli" fg={colors.info} attributes={5}>
              https://github.com/dodopayments/dodopayments-cli
            </a>
            <span fg={colors.textMuted}> for the update guide.</span>
          </text>
        }
      >
        <text>
          <span fg={colors.textMuted}>Run </span>
          <span fg={colors.accentLime} attributes={1}>/update</span>
          <span fg={colors.textMuted}> to install now.</span>
        </text>
      </Show>
    </box>
  );
};
