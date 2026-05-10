/**
 * Block variant union for the message stream. Each variant maps onto a
 * Solid component in `src/tui/components/blocks/` via the dispatcher in
 * `components/blocks/index.tsx`. step-form is reserved -- no command
 * currently emits it and there is no renderer for it.
 */

export type BlockVariant =
  | { type: 'spinner'; label: string }
  | { type: 'table'; data: any[]; statusColumn?: string }
  | { type: 'detail'; data: Record<string, any> }
  | { type: 'error'; message: string }
  | { type: 'success'; message: string }
  | { type: 'info'; message: string }
  | { type: 'link'; text: string; url: string }
  | { type: 'empty' }
  | { type: 'streaming'; text: string }
  | { type: 'event'; event: any }
  | { type: 'help' }
  | { type: 'inline-input'; label: string; secure?: boolean; onSubmit: (val: string) => void }
  | { type: 'inline-select'; label?: string; options: { label: string; value: string }[]; onSubmit: (val: string) => void }
  | { type: 'step-form'; fields: any[]; onSubmit: (data: any) => void }
  | { type: 'confirm'; message: string; onConfirm: () => void; onCancel: () => void };

export type BlockType = BlockVariant & { id: string };

export type Message = {
  id: string;
  role: 'user' | 'system';
  text?: string;
  blocks: BlockType[];
};
