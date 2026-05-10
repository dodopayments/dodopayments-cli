import { For, Show, createMemo, createSignal, onMount, untrack } from 'solid-js';
import { render, useKeyboard } from '@opentui/solid';
import { Welcome } from './components/Welcome';
import { InputBar } from './components/InputBar';
import { HintRow } from './components/HintRow';
import { Footer } from './components/Footer';
import { MessageRow } from './components/MessageRow';
import { Palette, rankCommands } from './components/Palette';
import { UpdateNotification } from './components/UpdateNotification';
import { TuiContextProvider, type AuthInfo } from './context';
import { createMessageStore } from './CommandContext';
import { handleCommand } from './router';
import { resolveCredentials, setSessionMode } from '../utils/auth';
import {
  checkForUpdates,
  consumePendingSilentUpdate,
  detectInstallMethod,
  dispatchSilentUpdate,
  type InstallMethod,
  type UpdateInfo,
} from '../utils/update';
import { version } from '../../package.json';

const App = () => {
  const [authInfo, setAuthInfo] = createSignal<AuthInfo>(null);
  const [input, setInput] = createSignal('');
  const [paletteIndex, setPaletteIndex] = createSignal(0);
  const [paletteDismissed, setPaletteDismissed] = createSignal(false);
  const [isProcessing, setIsProcessing] = createSignal(false);
  const [promptActive, setPromptActive] = createSignal(false);
  const [updateInfo, setUpdateInfo] = createSignal<UpdateInfo | null>(null);
  const installMethod: InstallMethod = detectInstallMethod();

  const store = createMessageStore();
  const hasMessages = createMemo(() => store.messages().length > 0);
  const paletteVisible = createMemo(() => {
    const v = input();
    if (!v.startsWith('/')) return false;
    if (paletteDismissed() || promptActive()) return false;
    const firstWord = v.split(' ')[0] ?? '';
    if (v.includes(' ') && firstWord.length > 1) {
      const exactMatch = rankCommands(firstWord).some((c) => c.command === firstWord);
      if (exactMatch) return false;
    }
    return true;
  });
  const palettePool = createMemo(() => rankCommands(input()));

  onMount(() => {
    resolveCredentials(undefined, false)
      .then(({ mode, apiKey }) => {
        setSessionMode(mode);
        const masked = apiKey.slice(0, 10) + '...' + apiKey.slice(-3);
        setAuthInfo({ mode, key: masked });
      })
      .catch(() => setAuthInfo(null));

    const completed = consumePendingSilentUpdate();
    if (completed && completed.to !== completed.from) {
      store.ctx.addBlock({
        type: 'success',
        message: `Updated to v${completed.to} (was v${completed.from}). Restart to use the new version.`,
      });
    }

    void checkForUpdates(version).then((info) => {
      if (!info) return;
      if (info.delta === 'major') {
        setUpdateInfo(info);
        return;
      }
      if (installMethod === 'npm' || installMethod === 'bun') {
        dispatchSilentUpdate(info, installMethod);
      } else {
        setUpdateInfo(info);
      }
    });
  });

  const onInputChange = (next: string) => {
    setInput(next);
    setPaletteIndex(0);
    if (paletteDismissed() && next.length > 0) setPaletteDismissed(false);
  };

  const completeWith = (cmd: string) => {
    const next = `${cmd} `;
    setInput(next);
    setPaletteDismissed(true);
  };

  const onSubmit = async (raw: string) => {
    const trimmed = raw.trim();
    if (!trimmed) return;
    if (paletteVisible()) {
      const pool = untrack(palettePool);
      const pick = pool[paletteIndex()];
      const looksIncomplete =
        !!pick && pick.command !== trimmed && pick.command.startsWith(trimmed);
      if (looksIncomplete) {
        completeWith(pick.command);
        return;
      }
    }
    store.pushUserEcho(trimmed);
    setInput('');
    setPaletteDismissed(false);
    setIsProcessing(true);
    try {
      await handleCommand(trimmed, store.ctx, () => process.exit(0));
    } catch (e: any) {
      store.ctx.addBlock({
        type: 'error',
        message: e?.message ?? String(e),
      });
    } finally {
      setIsProcessing(false);
    }
  };

  useKeyboard((key) => {
    if (!paletteVisible()) return;
    const pool = untrack(palettePool);
    if (key.name === 'up') {
      key.preventDefault();
      setPaletteIndex((i) => Math.max(0, i - 1));
    } else if (key.name === 'down') {
      key.preventDefault();
      setPaletteIndex((i) => Math.min(pool.length - 1, i + 1));
    } else if (key.name === 'tab') {
      key.preventDefault();
      const pick = pool[paletteIndex()];
      if (pick) completeWith(pick.command);
    } else if (key.name === 'escape') {
      key.preventDefault();
      setPaletteDismissed(true);
    }
  });

  return (
    <TuiContextProvider
      value={{
        authInfo,
        input,
        setInput,
        paletteVisible,
        paletteIndex,
        isProcessing,
        promptActive,
        setPromptActive,
      }}
    >
      <box flexDirection="column" width="100%" height="100%">
        <Show
          when={hasMessages()}
          fallback={<Welcome />}
        >
          <scrollbox
            flexGrow={1}
            stickyScroll={true}
            stickyStart="bottom"
            paddingLeft={2}
            paddingRight={2}
          >
            <Show when={updateInfo()}>
              {(info: () => UpdateInfo) => (
                <UpdateNotification info={info()} method={installMethod} />
              )}
            </Show>
            <For each={store.messages()}>{(m) => <MessageRow message={m} />}</For>
          </scrollbox>
        </Show>
        <box flexShrink={0} flexDirection="column" paddingLeft={2} paddingRight={2}>
          <InputBar onSubmit={onSubmit} onInput={onInputChange} value={input()} />
          <HintRow />
        </box>
        <Footer />
        <Palette
          query={input()}
          selectedIndex={paletteIndex()}
          visible={paletteVisible()}
        />
      </box>
    </TuiContextProvider>
  );
};

export const mountTuiApp = (): void => {
  render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
};
