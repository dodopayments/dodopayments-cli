import { createSignal } from 'solid-js';

export type ToastVariant = 'success' | 'info' | 'error';

export type Toast = {
  id: number;
  message: string;
  variant: ToastVariant;
};

const [toasts, setToasts] = createSignal<Toast[]>([]);
let nextId = 1;
const timers = new Map<number, ReturnType<typeof setTimeout>>();

export const useToasts = toasts;

export const dismissToast = (id: number): void => {
  const t = timers.get(id);
  if (t) {
    clearTimeout(t);
    timers.delete(id);
  }
  setToasts((prev) => prev.filter((toast) => toast.id !== id));
};

export const showToast = (
  message: string,
  options: { variant?: ToastVariant; durationMs?: number } = {},
): number => {
  const id = nextId++;
  const variant = options.variant ?? 'success';
  const durationMs = options.durationMs ?? 2000;
  setToasts((prev) => [...prev, { id, message, variant }]);
  if (durationMs > 0) {
    const handle = setTimeout(() => dismissToast(id), durationMs);
    timers.set(id, handle);
  }
  return id;
};
