import { colors, glyphs } from '../../theme';

export const Success = (props: { message: string }) => (
  <text>
    <span fg={colors.success}>{`${glyphs.check} `}</span>
    <span fg={colors.textPrimary}>{props.message}</span>
  </text>
);

export const Error = (props: { message: string }) => (
  <box borderStyle="round" borderColor={colors.error} paddingLeft={1} paddingRight={1}>
    <text>
      <span fg={colors.error}>{`${glyphs.cross} `}</span>
      <span fg={colors.textPrimary}>{props.message}</span>
    </text>
  </box>
);

export const Info = (props: { message: string }) => (
  <text>
    <span fg={colors.info}>{`${glyphs.bullet} `}</span>
    <span fg={colors.textPrimary}>{props.message}</span>
  </text>
);

export const Empty = () => (
  <text fg={colors.textDim}>{`${glyphs.dot} No results found.`}</text>
);

export const Link = (props: { text: string; url: string }) => (
  <text>
    <span fg={colors.textMuted}>{`${props.text} `}</span>
    <a href={props.url} fg={colors.info}>{props.url}</a>
  </text>
);

export const Event = (props: { event: any }) => (
  <text>
    <span fg={colors.testMode}>{`[${new Date().toLocaleTimeString()}] `}</span>
    <span fg={colors.textPrimary}>{JSON.stringify(props.event)}</span>
  </text>
);

const ANSI_REGEX = /\u001b\[[0-9;]*m/g;
const stripAnsi = (s: string): string => s.replace(ANSI_REGEX, '');

export const Streaming = (props: { text: string }) => (
  <markdown content={stripAnsi(props.text)} />
);
