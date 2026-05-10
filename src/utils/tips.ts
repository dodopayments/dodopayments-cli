import chalk from 'chalk';
import { colors } from '../tui/theme';

export const paginationTip = (command: string) =>
  '\n' +
  chalk.hex(colors.textDim)('— Use ') +
  chalk.hex(colors.info)(`${command} <page>`) +
  chalk.hex(colors.textDim)(' for more.');
