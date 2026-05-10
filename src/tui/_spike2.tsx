/**
 * OpenTUI Phase 0 spike v2 — proves the Solid + JSX path works without the
 * glitches that the factory-API spike (`_spike.ts`) produced.
 *
 * The v1 spike failed visually because manual `renderer.root.removeAll() + add()`
 * re-renders bypassed OpenTUI's reconciler. This file uses Solid signals + <For>
 * + JSX intrinsic elements, exactly the pattern Phase 1+ of the rewrite will use.
 *
 * Validates:
 *  1. JSX runtime works (preload + tsconfig.tui.json)
 *  2. Signals re-render without glitches
 *  3. <scrollbox> sticky-bottom + sliding window
 *  4. <input> with key handling
 *  5. <Portal> overlay (popup positioning)
 *  6. Ctrl+C exits cleanly
 *
 * Run: bun --tsconfig-override tsconfig.tui.json src/tui/_spike2.tsx
 */

import { render, useKeyboard } from "@opentui/solid";
import { createSignal, For, Show, onMount } from "solid-js";

const COLORS = {
  brand: "#07BC70",
  textPrimary: "#FFFFFF",
  textMuted: "#737470",
  textDim: "#535452",
  accentLime: "#C6FE1E",
  accentSky: "#38BDF8",
} as const;

const LOGO = [
  "█▀▀▄ █▀▀█ █▀▀▄ █▀▀█",
  "█  █ █  █ █  █ █  █",
  "█▄▄▀ ▀▀▀▀ █▄▄▀ ▀▀▀▀",
];

interface FakeMsg {
  id: number;
  who: "user" | "system";
  text: string;
}

const App = () => {
  const [input, setInput] = createSignal("");
  const [messages, setMessages] = createSignal<FakeMsg[]>(
    Array.from({ length: 50 }, (_, i) => ({
      id: i,
      who: i % 3 === 0 ? "user" : "system",
      text:
        i % 3 === 0
          ? `❯ /products list ${Math.floor(i / 3) + 1}`
          : `row ${i}: pdt_${Math.random().toString(36).slice(2, 10)} · $${(Math.random() * 100).toFixed(2)}`,
    })),
  );

  const showOverlay = () => input().startsWith("/");
  const inputEmpty = () => input().length === 0;

  useKeyboard((key) => {
    if (key.name === "return") {
      const v = input().trim();
      if (v) {
        setMessages((prev) => [
          ...prev,
          { id: prev.length, who: "user", text: `❯ ${v}` },
        ]);
      }
      setInput("");
    } else if (key.name === "backspace") {
      setInput((v) => v.slice(0, -1));
    } else if (key.name === "escape") {
      if (showOverlay()) setInput("");
    } else if (
      key.sequence &&
      key.sequence.length === 1 &&
      !key.ctrl &&
      !key.meta
    ) {
      setInput((v) => v + key.sequence);
    }
  });

  return (
    <box flexDirection="column" width="100%" height="100%">
      <box flexDirection="column" paddingTop={1} paddingLeft={2} flexShrink={0}>
        <For each={LOGO}>{(row) => <text fg={COLORS.brand}>{row}</text>}</For>
        <text> </text>
        <text fg={COLORS.textMuted}>spike v2 · OpenTUI + Solid + JSX</text>
        <text fg={COLORS.textDim}>type / to open overlay · ctrl+c to exit</text>
      </box>

      <scrollbox
        flexGrow={1}
        stickyScroll={true}
        stickyStart="bottom"
        paddingLeft={2}
        paddingRight={2}
      >
        <For each={messages()}>
          {(m) => (
            <text fg={m.who === "user" ? COLORS.textPrimary : COLORS.textMuted}>
              {m.text}
            </text>
          )}
        </For>
      </scrollbox>

      <box flexShrink={0} flexDirection="column">
        <text fg={COLORS.textDim}>{"─".repeat(120)}</text>
        <box paddingLeft={1} flexDirection="row">
          <text fg={COLORS.accentLime}>{"❯ "}</text>
          <text fg={input() ? COLORS.textPrimary : COLORS.textDim}>
            {input() || "type a command…"}
          </text>
        </box>
        <text fg={COLORS.textDim} paddingLeft={1}>
          {showOverlay()
            ? "↑↓ navigate · ↵ select · esc cancel"
            : inputEmpty()
              ? "/ palette · ↑↓ history · ctrl+c exit"
              : "↵ submit · backspace edit · ctrl+c exit"}
        </text>
      </box>

      <Show when={showOverlay()}>
        <box
          position="absolute"
          top={8}
          left={30}
          width={60}
          borderStyle="single"
          borderColor={COLORS.accentSky}
          backgroundColor="#0D0D0D"
          padding={1}
          flexDirection="column"
        >
          <text fg={COLORS.accentSky}>◆ command palette · spike v2</text>
          <text fg={COLORS.textDim}>{"─".repeat(56)}</text>
          <text fg={COLORS.accentLime}>{`❯ ${input()}`}</text>
          <text fg={COLORS.textDim}>{"─".repeat(56)}</text>
          <text fg={COLORS.accentLime}>
            ❯ /payments list · List your payments
          </text>
          <text fg={COLORS.textPrimary}>
            {"  "}/products list · List your products
          </text>
          <text fg={COLORS.textPrimary}>
            {"  "}/customers list · List your customers
          </text>
        </box>
      </Show>
    </box>
  );
};

onMount(() => {});

render(() => <App />, { exitOnCtrlC: true, targetFps: 30 });
