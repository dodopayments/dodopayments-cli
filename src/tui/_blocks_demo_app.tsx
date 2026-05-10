/**
 * Phase 4 throwaway: cycles through every block variant with realistic
 * fixture data. Use the InputBar to type variant names (success, error,
 * info, empty, link, table, detail, help, spinner-cycle, streaming, event,
 * input, select, confirm, all). Deleted in Phase 8.
 */

import { createSignal, For } from 'solid-js';
import { render } from '@opentui/solid';
import { colors, glyphs } from './theme';
import { TuiContextProvider, type AuthInfo } from './context';
import { createMessageStore, type CommandContext } from './CommandContext';
import { Welcome } from './components/Welcome';
import { StatusBar } from './components/StatusBar';
import { InputBar } from './components/InputBar';
import { HintBar } from './components/HintBar';
import { MessageRow } from './components/MessageRow';

const PRODUCTS_FIXTURE = [
  { name: 'Pro Plan', product_id: 'pdt_aaa111', created_at: '2025-01-12', price: '$29.00' },
  { name: 'Team Plan', product_id: 'pdt_bbb222', created_at: '2025-02-03', price: '$99.00' },
  { name: 'Enterprise', product_id: 'pdt_ccc333', created_at: '2025-03-21', price: '$499.00' },
];

const CUSTOMER_FIXTURE = {
  customer_id: 'cus_xxx',
  name: 'Ada Lovelace',
  email: 'ada@example.com',
  created_at: '2024-11-02',
  total_spent_amount: '$1,540.00',
  active_subscriptions: 3,
};

async function runBlock(arg: string, ctx: CommandContext) {
  switch (arg) {
    case 'success':
      ctx.addBlock({ type: 'success', message: 'It worked.' });
      return;
    case 'error':
      ctx.addBlock({ type: 'error', message: 'Something exploded.' });
      return;
    case 'info':
      ctx.addBlock({ type: 'info', message: 'For your information.' });
      return;
    case 'empty':
      ctx.addBlock({ type: 'empty' });
      return;
    case 'link':
      ctx.addBlock({
        type: 'link',
        text: 'Open dashboard:',
        url: 'https://dashboard.dodopayments.com/products/new',
      });
      return;
    case 'table':
      ctx.addBlock({ type: 'table', data: PRODUCTS_FIXTURE });
      return;
    case 'detail':
      ctx.addBlock({ type: 'detail', data: CUSTOMER_FIXTURE });
      return;
    case 'help':
      ctx.addBlock({ type: 'help' });
      return;
    case 'event':
      ctx.addBlock({
        type: 'event',
        event: { type: 'payment.success', id: 'pay_demo', amount_cents: 4900 },
      });
      return;
    case 'streaming':
      ctx.addBlock({
        type: 'streaming',
        text: '## Streaming markdown\n\nHere is **bold**, _italic_, and `inline code`.\n\n```ts\nconst x = 1;\n```',
      });
      return;
    case 'spinner-cycle': {
      const id = ctx.addBlock({ type: 'spinner', label: 'fetching products…' });
      await new Promise((r) => setTimeout(r, 1500));
      ctx.updateBlock(id, { type: 'spinner', label: 'still fetching…' });
      await new Promise((r) => setTimeout(r, 1500));
      ctx.removeBlock(id);
      ctx.addBlock({ type: 'success', message: 'Done.' });
      return;
    }
    case 'input': {
      const value = await ctx.promptInput('Enter your API key:');
      ctx.addBlock({ type: 'success', message: `got: ${value}` });
      return;
    }
    case 'select': {
      const value = await ctx.promptSelect('Pick a plan:', [
        { label: 'Pro', value: 'pro' },
        { label: 'Team', value: 'team' },
        { label: 'Enterprise', value: 'enterprise' },
      ]);
      ctx.addBlock({ type: 'success', message: `picked: ${value}` });
      return;
    }
    case 'confirm': {
      const ok = await ctx.promptConfirm('Are you sure?');
      ctx.addBlock({ type: 'success', message: `confirmed: ${ok}` });
      return;
    }
    case 'all': {
      for (const v of ['success', 'error', 'info', 'empty', 'link', 'event', 'table', 'detail', 'streaming', 'help']) {
        await runBlock(v, ctx);
      }
      return;
    }
    default:
      ctx.addBlock({
        type: 'info',
        message:
          'try: success | error | info | empty | link | table | detail | help | event | streaming | spinner-cycle | input | select | confirm | all',
      });
  }
}

const App = () => {
  const store = createMessageStore();
  const [input, setInput] = createSignal('');
  const [authInfo] = createSignal<AuthInfo>(null);
  const [paletteVisible] = createSignal(false);
  const [isProcessing] = createSignal(false);
  const [promptActive, setPromptActive] = createSignal(false);

  const handleSubmit = async (raw: string) => {
    store.pushUserEcho(`${glyphs.prompt} ${raw}`);
    await runBlock(raw.trim(), store.ctx);
  };

  return (
    <TuiContextProvider
      value={{ authInfo, input, paletteVisible, isProcessing, promptActive, setPromptActive }}
    >
      <box flexDirection="column" width="100%" height="100%">
        {store.messages().length === 0 ? <Welcome /> : null}
        <scrollbox
          flexGrow={1}
          stickyScroll={true}
          stickyStart="bottom"
          paddingLeft={2}
          paddingRight={2}
        >
          <For each={store.messages()}>{(m) => <MessageRow message={m} />}</For>
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

export const mountBlocksDemo = (): void => {
  render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
};
