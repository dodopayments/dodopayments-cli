import type { BlockVariant } from './types';

export interface CommandContext {
  addBlock: (block: BlockVariant) => string;
  updateBlock: (id: string, block: Partial<BlockVariant>) => void;
  removeBlock: (id: string) => void;
  promptInput: (label: string, secure?: boolean) => Promise<string>;
  promptSelect: (label: string, options: {label: string, value: string}[]) => Promise<string>;
  promptConfirm: (message: string) => Promise<boolean>;
}
