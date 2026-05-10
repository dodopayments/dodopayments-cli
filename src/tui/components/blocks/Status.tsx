import { colors, glyphs } from '../../theme';

export const Success = (props: { message: string }) => (
  <box flexDirection="row">
    <text fg={colors.success}>{`${glyphs.check} `}</text>
    <text fg={colors.textPrimary}>{props.message}</text>
  </box>
);

export const Error = (props: { message: string }) => (
  <box
    borderStyle="rounded"
    borderColor={colors.error}
    paddingLeft={1}
    paddingRight={1}
    flexDirection="row"
  >
    <text fg={colors.error}>{`${glyphs.cross} `}</text>
    <text fg={colors.textPrimary}>{props.message}</text>
  </box>
);

export const Info = (props: { message: string }) => (
  <box flexDirection="row">
    <text fg={colors.info}>{`${glyphs.bullet} `}</text>
    <text fg={colors.textPrimary}>{props.message}</text>
  </box>
);

export const Empty = () => (
  <text fg={colors.textDim}>{`${glyphs.dot} No results found.`}</text>
);

export const Link = (props: { text: string; url: string }) => (
  <box flexDirection="row">
    <text fg={colors.textMuted}>{`${props.text} `}</text>
    <text fg={colors.info}>{props.url}</text>
  </box>
);

export const Event = (props: { event: any }) => (
  <box flexDirection="row">
    <text fg={colors.testMode}>{`[${new Date().toLocaleTimeString()}] `}</text>
    <text fg={colors.textPrimary}>{JSON.stringify(props.event)}</text>
  </box>
);

const ANSI_REGEX = /\u001b\[[0-9;]*m/g;
const stripAnsi = (s: string): string => s.replace(ANSI_REGEX, '');

export const Streaming = (props: { text: string }) => (
  <text fg={colors.textPrimary}>{stripAnsi(props.text)}</text>
);
