import { Box, Text } from 'ink';
import { version } from '../../../package.json';

const LOGO_TEXT = [
  `  ______          _        ______                                _       `,
  `  |  _  \\        | |       | ___ \\                              | |      `,
  `  | | | |___   __| | ___   | |_/ /_ _ _   _ _ __ ___   ___ _ __ | |_ ___ `,
  `  | | | / _ \\ / _\` |/ _ \\  |  __/ _\` | | | | '_ \` _ \\ / _ \\ '_ \\| __/ __|`,
  `  | |/ / (_) | (_| | (_) | | | | (_| | |_| | | | | | |  __/ | | | |_\\__ \\`,
  `  |___/ \\___/ \\__,_|\\___/  \\_|  \\__,_|\\__, |_| |_| |_|\\___|_| |_|\\__|___/`,
  `                                       __/ |                             `,
  `                                      |___/                              `,
];

export const WelcomeBanner = () => {
  return (
    <Box flexDirection="column" paddingX={1}>
      <Box flexDirection="row" alignItems="center">
        {/* Text logo on the right */}
        <Box flexDirection="column" justifyContent="center" flexShrink={0}>
          {LOGO_TEXT.map((line, i) => (
            <Text key={`logo-${i}`} color="#07BC70" bold wrap="truncate">{line}</Text>
          ))}
          <Box marginTop={1} flexDirection="row" justifyContent="space-between" width="100%">
            <Text color="gray" wrap="truncate">  The official CLI to manage Dodo Payments!</Text>
            <Text color="gray">v{version}</Text>
          </Box>
        </Box>
      </Box>

      <Text> </Text>
    </Box>
  );
};
