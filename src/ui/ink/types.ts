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
  | { type: 'inline-select'; label?: string; options: {label: string, value: string}[]; onSubmit: (val: string) => void }
  | { type: 'step-form'; fields: any[]; onSubmit: (data: any) => void }
  | { type: 'confirm'; message: string; onConfirm: () => void; onCancel: () => void };

export type BlockType = BlockVariant & { id: string };

export type Message = {
  id: string;
  role: 'user' | 'system';
  text?: string;
  blocks: BlockType[];
};
