/**
 * OpenTUI Phase 0 spike — proves the runtime works in this Bun environment
 * before committing to the full rewrite.
 *
 * Validates:
 *  1. DODO ASCII wordmark renders with brand-green color (text + color)
 *  2. ScrollBox with 50 fake message rows scrolls via mouse wheel + PgUp/PgDn
 *  3. Input with `❯` lime cursor accepts keystrokes (text input + key handling)
 *  4. Centered overlay box renders when input starts with '/' (popup positioning)
 *  5. Ctrl+C exits cleanly and restores the parent terminal (alt-screen lifecycle)
 *
 * Run: bun src/tui/_spike.ts
 */

import {
  createCliRenderer,
  Box,
  Text,
  ScrollBox,
  Input,
  type CliRenderer,
} from "@opentui/core";

const COLORS = {
  brand: "#07BC70",
  textPrimary: "#FFFFFF",
  textMuted: "#737470",
  textDim: "#535452",
  accentLime: "#C6FE1E",
  accentSky: "#38BDF8",
  border: "#212423",
} as const;

const LOGO = [
  "█▀▀▄ █▀▀█ █▀▀▄ █▀▀█",
  "█  █ █  █ █  █ █  █",
  "█▄▄▀ ▀▀▀▀ █▄▄▀ ▀▀▀▀",
];

const FAKE_MESSAGES = Array.from({ length: 50 }, (_, i) => ({
  id: i,
  who: i % 3 === 0 ? "user" : "system",
  text:
    i % 3 === 0
      ? `❯ /products list ${Math.floor(i / 3) + 1}`
      : `row ${i}: pdt_${Math.random().toString(36).slice(2, 10)} · $${(Math.random() * 100).toFixed(2)}`,
}));

let inputValue = "";
let renderer: CliRenderer;

const renderApp = (input: string) => {
  const showOverlay = input.startsWith("/");

  const welcome = Box(
    { flexDirection: "column", paddingTop: 1, paddingLeft: 2 },
    ...LOGO.map((row) => Text({ content: row, fg: COLORS.brand })),
    Text({ content: "" }),
    Text({ content: "spike · OpenTUI capability proof", fg: COLORS.textMuted }),
    Text({ content: "scroll with ↑↓ PgUp/PgDn or mouse wheel", fg: COLORS.textDim }),
    Text({ content: "type / to open overlay · ctrl+c to exit", fg: COLORS.textDim }),
  );

  const messageList = ScrollBox(
    {
      stickyScroll: true,
      stickyStart: "bottom",
      flexGrow: 1,
      paddingLeft: 2,
      paddingRight: 2,
    },
    ...FAKE_MESSAGES.map((m) =>
      Text({
        content: m.text,
        fg: m.who === "user" ? COLORS.textPrimary : COLORS.textMuted,
      })
    ),
  );

  const divider = Text({
    content: "─".repeat(120),
    fg: COLORS.textDim,
  });

  const inputRow = Box(
    { flexDirection: "row", paddingLeft: 1 },
    Text({ content: "❯ ", fg: COLORS.accentLime }),
    Text({ content: input || "type a command…", fg: input ? COLORS.textPrimary : COLORS.textDim }),
  );

  const hintRow = Text({
    content: showOverlay
      ? "↑↓ navigate · ↵ select · esc cancel"
      : "/ palette · ↑↓ history · ctrl+c exit",
    fg: COLORS.textDim,
  });

  const children: any[] = [welcome, messageList, divider, inputRow, hintRow];

  if (showOverlay) {
    const overlay = Box(
      {
        position: "absolute",
        top: 8,
        left: 30,
        width: 60,
        borderStyle: "single",
        borderColor: COLORS.accentSky,
        backgroundColor: "#0D0D0D",
        padding: 1,
        flexDirection: "column",
      },
      Text({ content: "◆ command palette · spike", fg: COLORS.accentSky }),
      Text({ content: "─".repeat(56), fg: COLORS.textDim }),
      Text({ content: `❯ ${input}`, fg: COLORS.accentLime }),
      Text({ content: "─".repeat(56), fg: COLORS.textDim }),
      Text({ content: "❯ /payments list  · List your payments", fg: COLORS.accentLime }),
      Text({ content: "  /products list  · List your products", fg: COLORS.textPrimary }),
      Text({ content: "  /customers list · List your customers", fg: COLORS.textPrimary }),
    );
    children.push(overlay);
  }

  return Box(
    { flexDirection: "column", width: "100%", height: "100%" },
    ...children,
  );
};

(async () => {
  renderer = await createCliRenderer({
    exitOnCtrlC: true,
    targetFps: 30,
  });

  const root = renderApp(inputValue);
  renderer.root.add(root);

  const input = Input({
    placeholder: "spike",
    focused: true,
  });

  const keyboardHandler = (key: any) => {
    let changed = false;
    if (key.name === "return") {
      inputValue = "";
      changed = true;
    } else if (key.name === "backspace") {
      if (inputValue.length > 0) {
        inputValue = inputValue.slice(0, -1);
        changed = true;
      }
    } else if (key.name === "escape") {
      if (inputValue.startsWith("/")) {
        inputValue = "";
        changed = true;
      }
    } else if (key.sequence && key.sequence.length === 1 && !key.ctrl && !key.meta) {
      inputValue = inputValue + key.sequence;
      changed = true;
    }
    if (changed) {
      renderer.root.removeAll();
      renderer.root.add(renderApp(inputValue));
    }
  };

  renderer.keyInput.on("keypress", keyboardHandler);
})();
