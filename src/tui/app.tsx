/**
 * Solid + OpenTUI render tree. Loaded only via `bootstrap.ts` AFTER the
 * Solid babel plugin is installed, so the JSX here compiles to Solid factory
 * calls (not React.createElement).
 *
 * Phase 2 layout contract (do not break in later phases without updating
 * the plan):
 *   - Top group: Welcome (visible until messages.length > 0)
 *   - Middle: <scrollbox flexGrow={1} stickyScroll stickyStart="bottom">
 *   - Bottom chrome group: <box flexShrink={0}> wrapping StatusBar +
 *     InputBar + HintBar. The flexShrink={0} wrapper is mandatory --
 *     without it the scrollbox starves the chrome of height (verified
 *     during spike v2).
 */

import { createSignal, Show } from 'solid-js';
import { render } from '@opentui/solid';
import { Welcome } from './components/Welcome';
import { StatusBar } from './components/StatusBar';
import { InputBar } from './components/InputBar';
import { HintBar } from './components/HintBar';
import { TuiContextProvider, type AuthInfo } from './context';

const App = () => {
  const [authInfo] = createSignal<AuthInfo>(null);
  const [input, setInput] = createSignal('');
  const [paletteVisible] = createSignal(false);
  const [isProcessing] = createSignal(false);
  const [promptActive, setPromptActive] = createSignal(false);
  const [hasMessages] = createSignal(false);

  return (
    <TuiContextProvider
      value={{ authInfo, input, paletteVisible, isProcessing, promptActive, setPromptActive }}
    >
      <box flexDirection="column" width="100%" height="100%">
        <Show when={!hasMessages()}>
          <Welcome />
        </Show>
        <scrollbox
          flexGrow={1}
          stickyScroll={true}
          stickyStart="bottom"
          paddingLeft={2}
          paddingRight={2}
        />
        <box flexShrink={0} flexDirection="column">
          <StatusBar />
          <InputBar onSubmit={() => {}} onInput={setInput} />
          <HintBar />
        </box>
      </box>
    </TuiContextProvider>
  );
};

export const mountTuiApp = (): void => {
  render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
};
