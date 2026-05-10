/**
 * Phase 3 throwaway: exercises every CommandContext primitive
 * (addBlock/updateBlock/removeBlock + each prompt) via fake /test commands.
 * Uses minimal placeholder block renderers since the real ones land in Phase 4.
 * Deleted in Phase 8.
 */

import { createSignal, For, Match, Switch } from 'solid-js';
import { render } from '@opentui/solid';
import { colors, glyphs } from './theme';
import { TuiContextProvider, type AuthInfo } from './context';
import { createMessageStore, type CommandContext } from './CommandContext';
import { Welcome } from './components/Welcome';
import { StatusBar } from './components/StatusBar';
import { InputBar } from './components/InputBar';
import { HintBar } from './components/HintBar';
import type { BlockType } from './types';

const renderBlock = (block: BlockType) => {
  return (
    <Switch fallback={<text fg={colors.textDim}>· {block.type}</text>}>
      <Match when={block.type === 'spinner' && block}>
        {(b) => <text fg={colors.warning}>⠋ {(b() as { label: string }).label}</text>}
      </Match>
      <Match when={block.type === 'success' && block}>
        {(b) => <text fg={colors.success}>{glyphs.check} {(b() as { message: string }).message}</text>}
      </Match>
      <Match when={block.type === 'error' && block}>
        {(b) => <text fg={colors.error}>{glyphs.cross} {(b() as { message: string }).message}</text>}
      </Match>
      <Match when={block.type === 'info' && block}>
        {(b) => <text fg={colors.info}>{glyphs.bullet} {(b() as { message: string }).message}</text>}
      </Match>
      <Match when={block.type === 'inline-input' && block}>
        {(b) => {
          const blk = b() as { label: string; onSubmit: (v: string) => void };
          const [val, setVal] = createSignal('');
          return (
            <box flexDirection="column" borderStyle="round" borderColor={colors.accentMagenta} padding={1}>
              <text fg={colors.textPrimary}>{blk.label}</text>
              <input
                focused
                value={val()}
                onInput={setVal}
                onSubmit={(v) => blk.onSubmit(v)}
              />
            </box>
          );
        }}
      </Match>
      <Match when={block.type === 'confirm' && block}>
        {(b) => {
          const blk = b() as { message: string; onConfirm: () => void; onCancel: () => void };
          return (
            <select
              focused
              options={[
                { name: 'Yes', value: 'yes', description: blk.message },
                { name: 'No', value: 'no', description: blk.message },
              ]}
              onSelect={(_i, opt) => {
                if (opt?.value === 'yes') blk.onConfirm();
                else blk.onCancel();
              }}
            />
          );
        }}
      </Match>
      <Match when={block.type === 'inline-select' && block}>
        {(b) => {
          const blk = b() as {
            label?: string;
            options: { label: string; value: string }[];
            onSubmit: (v: string) => void;
          };
          return (
            <box flexDirection="column">
              {blk.label ? <text fg={colors.textPrimary}>{blk.label}</text> : null}
              <select
                focused
                options={blk.options.map((o) => ({ name: o.label, value: o.value, description: o.label }))}
                onSelect={(_i, opt) => {
                  if (opt) blk.onSubmit(opt.value);
                }}
              />
            </box>
          );
        }}
      </Match>
    </Switch>
  );
};

const App = () => {
  const store = createMessageStore();
  const [input, setInput] = createSignal('');
  const [authInfo] = createSignal<AuthInfo>(null);
  const [paletteVisible] = createSignal(false);
  const [isProcessing] = createSignal(false);

  const handleSubmit = async (raw: string) => {
    store.pushUserEcho(`${glyphs.prompt} ${raw}`);
    await runFakeCommand(raw, store.ctx);
  };

  return (
    <TuiContextProvider value={{ authInfo, input, paletteVisible, isProcessing }}>
      <box flexDirection="column" width="100%" height="100%">
        {store.messages().length === 0 ? <Welcome /> : null}
        <scrollbox
          flexGrow={1}
          stickyScroll={true}
          stickyStart="bottom"
          paddingLeft={2}
          paddingRight={2}
        >
          <For each={store.messages()}>
            {(m) => (
              <box flexDirection="column">
                {m.text ? (
                  <text fg={m.role === 'user' ? colors.textPrimary : colors.textMuted}>
                    {m.text}
                  </text>
                ) : null}
                <For each={m.blocks}>{(b) => renderBlock(b)}</For>
              </box>
            )}
          </For>
        </scrollbox>
        <box flexShrink={0} flexDirection="column">
          <StatusBar />
          <InputBar onSubmit={handleSubmit} onInput={setInput} />
          <HintBar />
        </box>
      </box>
    </TuiContextProvider>
  );
};

async function runFakeCommand(raw: string, ctx: CommandContext): Promise<void> {
  const arg = raw.replace(/^\/test\s*/, '').trim();
  switch (arg) {
    case 'success':
      ctx.addBlock({ type: 'success', message: 'success block rendered' });
      return;
    case 'error':
      ctx.addBlock({ type: 'error', message: 'error block rendered' });
      return;
    case 'info':
      ctx.addBlock({ type: 'info', message: 'info block rendered' });
      return;
    case 'spinner-cycle': {
      const id = ctx.addBlock({ type: 'spinner', label: 'working…' });
      await new Promise((r) => setTimeout(r, 1500));
      ctx.updateBlock(id, { type: 'spinner', label: 'still working…' });
      await new Promise((r) => setTimeout(r, 1500));
      ctx.removeBlock(id);
      ctx.addBlock({ type: 'success', message: 'spinner cycle complete (update + remove verified)' });
      return;
    }
    case 'prompt-input': {
      const value = await ctx.promptInput('Enter a test value:');
      ctx.addBlock({ type: 'success', message: `got input: ${value}` });
      return;
    }
    case 'prompt-confirm': {
      const ok = await ctx.promptConfirm('Confirm test?');
      ctx.addBlock({ type: 'success', message: `confirmed: ${ok}` });
      return;
    }
    case 'prompt-select': {
      const value = await ctx.promptSelect('Pick one:', [
        { label: 'Option A', value: 'a' },
        { label: 'Option B', value: 'b' },
        { label: 'Option C', value: 'c' },
      ]);
      ctx.addBlock({ type: 'success', message: `picked: ${value}` });
      return;
    }
    default:
      ctx.addBlock({
        type: 'info',
        message: 'try: /test success | /test error | /test info | /test spinner-cycle | /test prompt-input | /test prompt-confirm | /test prompt-select',
      });
  }
}

export const mountCtxDemo = (): void => {
  render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
};
