/**
 * Headless (non-TTY) implementation of CommandContext. Shells through to
 * console.log / console.error so commands invoked as `dodo products list 1`
 * produce script-friendly output without mounting any TUI. Prompts throw --
 * pipelines should not call promptInput/promptSelect/promptConfirm.
 */

import type { CommandContext } from './CommandContext';
import type { BlockVariant } from './types';

export const defaultContext: CommandContext = {
  addBlock: (b: BlockVariant) => {
    if (b.type === 'table') console.table(b.data);
    else if (b.type === 'detail') console.table(b.data);
    else if (b.type === 'error') console.error(b.message);
    else if (b.type === 'success') console.log('\u2713', b.message);
    else if (b.type === 'info') console.log(b.message);
    else if (b.type === 'link') console.log('To view, go to:', b.url);
    else if (b.type === 'streaming') process.stdout.write(b.text);
    else if (b.type === 'event') console.log(b.event);
    else if (b.type === 'empty') console.log('No results found.');
    else if (b.type === 'help') console.log('Commands...');
    return '';
  },
  updateBlock: () => {},
  removeBlock: () => {},
  promptInput: async () => {
    throw new Error('Cannot prompt in non-TTY mode');
  },
  promptSelect: async () => {
    throw new Error('Cannot prompt in non-TTY mode');
  },
  promptConfirm: async () => {
    throw new Error('Cannot prompt in non-TTY mode');
  },
  clear: () => console.clear(),
};
