/**
 * Phase 2 throwaway: streams 50 fake messages over 5 seconds to validate
 * sticky-bottom scrollbox + bottom-anchored chrome layout. Loaded via
 * `_chrome_demo.ts` bootstrap. Deleted in Phase 8.
 */

import { createSignal, For, onMount } from 'solid-js';
import { render } from '@opentui/solid';
import { colors } from './theme';
import { TuiContextProvider, type AuthInfo } from './context';
import { Welcome } from './components/Welcome';
import { StatusBar } from './components/StatusBar';
import { InputBar } from './components/InputBar';
import { HintBar } from './components/HintBar';

type Msg = { id: number; who: 'user' | 'system'; text: string };

const App = () => {
  const [messages, setMessages] = createSignal<Msg[]>([]);
  const [input, setInput] = createSignal('');
  const [authInfo] = createSignal<AuthInfo>(null);
  const [paletteVisible] = createSignal(false);
  const [isProcessing] = createSignal(false);

  onMount(() => {
    let i = 0;
    const t = setInterval(() => {
      if (i >= 50) {
        clearInterval(t);
        return;
      }
      setMessages((prev) => [
        ...prev,
        {
          id: i,
          who: i % 3 === 0 ? 'user' : 'system',
          text:
            i % 3 === 0
              ? `❯ /products list ${Math.floor(i / 3) + 1}`
              : `row ${i}: pdt_${Math.random().toString(36).slice(2, 10)}`,
        },
      ]);
      i++;
    }, 100);
  });

  return (
    <TuiContextProvider value={{ authInfo, input, paletteVisible, isProcessing }}>
      <box flexDirection="column" width="100%" height="100%">
        {messages().length === 0 ? <Welcome /> : null}
        <scrollbox
          flexGrow={1}
          stickyScroll={true}
          stickyStart="bottom"
          paddingLeft={2}
          paddingRight={2}
        >
          <For each={messages()}>
            {(m) => (
              <text fg={m.who === 'user' ? colors.textPrimary : colors.textMuted}>
                {m.text}
              </text>
            )}
          </For>
        </scrollbox>
        <box flexShrink={0} flexDirection="column">
          <StatusBar />
          <InputBar
            onSubmit={(v) => {
              setMessages((prev) => [
                ...prev,
                { id: prev.length, who: 'user', text: `❯ ${v}` },
              ]);
            }}
            onInput={setInput}
          />
          <HintBar />
        </box>
      </box>
    </TuiContextProvider>
  );
};

export const mountChromeDemo = (): void => {
  render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
};
