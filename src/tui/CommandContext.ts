/**
 * Solid-signal-backed implementation of the CommandContext interface that
 * every command handler in `src/commands/` consumes.
 *
 * Prompt resolution: addBlock for an inline-input/select/confirm variant
 * stores its onSubmit (or onConfirm/onCancel) callback. promptInput etc.
 * wrap the addBlock call in a Promise that resolves when the user submits.
 * The block is auto-removed on resolution so the message stream stays clean.
 */

import { createSignal, type Accessor, type Setter } from 'solid-js';
import type { BlockVariant, BlockType, Message } from './types';

export interface CommandContext {
  addBlock: (block: BlockVariant) => string;
  updateBlock: (id: string, block: Partial<BlockVariant>) => void;
  removeBlock: (id: string) => void;
  promptInput: (label: string, secure?: boolean) => Promise<string>;
  promptSelect: (label: string, options: { label: string; value: string }[]) => Promise<string>;
  promptConfirm: (message: string) => Promise<boolean>;
  clear: () => void;
}

export interface MessageStore {
  messages: Accessor<Message[]>;
  setMessages: Setter<Message[]>;
  pushUserEcho: (text: string) => void;
  ctx: CommandContext;
}

let blockCounter = 0;
const nextId = (): string => `blk_${Date.now()}_${blockCounter++}`;

export const createMessageStore = (): MessageStore => {
  const [messages, setMessages] = createSignal<Message[]>([]);

  const ensureSystemMessage = (): string => {
    const list = messages();
    const last = list[list.length - 1];
    if (last && last.role === 'system') return last.id;
    const id = nextId();
    setMessages([...list, { id, role: 'system', blocks: [] }]);
    return id;
  };

  const updateBlocks = (msgId: string, mutate: (blocks: BlockType[]) => BlockType[]) => {
    setMessages((prev) =>
      prev.map((m) => (m.id === msgId ? { ...m, blocks: mutate(m.blocks) } : m)),
    );
  };

  const addBlock: CommandContext['addBlock'] = (block) => {
    const msgId = ensureSystemMessage();
    const id = nextId();
    updateBlocks(msgId, (blocks) => [...blocks, { ...block, id } as BlockType]);
    return id;
  };

  const updateBlock: CommandContext['updateBlock'] = (id, patch) => {
    setMessages((prev) =>
      prev.map((m) => ({
        ...m,
        blocks: m.blocks.map((b) =>
          b.id === id ? ({ ...b, ...patch, id } as BlockType) : b,
        ),
      })),
    );
  };

  const removeBlock: CommandContext['removeBlock'] = (id) => {
    setMessages((prev) =>
      prev.map((m) => ({ ...m, blocks: m.blocks.filter((b) => b.id !== id) })),
    );
  };

  const promptInput: CommandContext['promptInput'] = (label, secure) => {
    return new Promise<string>((resolve) => {
      let blockId = '';
      blockId = addBlock({
        type: 'inline-input',
        label,
        secure,
        onSubmit: (value) => {
          removeBlock(blockId);
          resolve(value);
        },
      });
    });
  };

  const promptSelect: CommandContext['promptSelect'] = (label, options) => {
    return new Promise<string>((resolve) => {
      let blockId = '';
      blockId = addBlock({
        type: 'inline-select',
        label,
        options,
        onSubmit: (value) => {
          removeBlock(blockId);
          resolve(value);
        },
      });
    });
  };

  const promptConfirm: CommandContext['promptConfirm'] = (message) => {
    return new Promise<boolean>((resolve) => {
      let blockId = '';
      blockId = addBlock({
        type: 'confirm',
        message,
        onConfirm: () => {
          removeBlock(blockId);
          resolve(true);
        },
        onCancel: () => {
          removeBlock(blockId);
          resolve(false);
        },
      });
    });
  };

  const pushUserEcho = (text: string) => {
    const id = nextId();
    setMessages((prev) => [...prev, { id, role: 'user', text, blocks: [] }]);
  };

  const clear = () => {
    setMessages([]);
  };

  return {
    messages,
    setMessages,
    pushUserEcho,
    ctx: { addBlock, updateBlock, removeBlock, promptInput, promptSelect, promptConfirm, clear },
  };
};
