import { colors, glyphs } from '../theme';

interface Tip {
  prefix: string;
  command: string;
  suffix: string;
}

const TIPS: Tip[] = [
  { prefix: 'Type ', command: '/help', suffix: ' to see every command' },
  { prefix: 'Use ', command: '/login', suffix: ' to start using the API' },
  { prefix: 'Run ', command: '/products list', suffix: ' to view your catalog' },
  { prefix: 'Try ', command: '/wh trigger', suffix: ' offline to test webhooks' },
  { prefix: 'Press ', command: 'Tab', suffix: ' to autocomplete commands' },
  { prefix: 'Hit ', command: '/', suffix: ' to open the command palette' },
  { prefix: 'Use ', command: '/update', suffix: ' to grab the latest release' },
];

const tip = TIPS[Math.floor(Math.random() * TIPS.length)]!;

export const TipLine = () => (
  <box flexDirection="row" justifyContent="center" alignItems="center" flexShrink={0}>
    <text fg={colors.accentAmber}>{glyphs.dot}</text>
    <text fg={colors.accentAmber}> Tip</text>
    <text fg={colors.textPrimary}>{` ${tip.prefix}`}</text>
    <box
      paddingLeft={1}
      paddingRight={1}
      backgroundColor={colors.brandBlack}
      flexShrink={0}
    >
      <text fg={colors.accentLime}>{tip.command}</text>
    </box>
    <text fg={colors.textPrimary}>{tip.suffix}</text>
  </box>
);
