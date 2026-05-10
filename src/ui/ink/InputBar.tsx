import React, { useState, useEffect } from 'react';
import { Box, Text, useInput, useStdout } from 'ink';
import chalk from 'chalk';
import { Autocomplete, getSuggestions } from './Autocomplete';
import { colors, glyphs } from '../theme';

interface InputBarProps {
  onSubmit: (val: string) => void;
  onClear: () => void;
  onExit: () => void;
  isActive: boolean;
}

export const InputBar = ({ onSubmit, onClear, onExit, isActive }: InputBarProps) => {
  const [input, setInput] = useState('');
  const [cursorOffset, setCursorOffset] = useState(0);
  const [selectedIndex, setSelectedIndex] = useState(0);
  const [history, setHistory] = useState<string[]>([]);
  const [historyIndex, setHistoryIndex] = useState(-1);
  const [paletteDismissed, setPaletteDismissed] = useState(false);

  const rawSuggestions = getSuggestions(input);
  const exactMatch = rawSuggestions.length === 1
    && rawSuggestions[0]!.command.toLowerCase() === input.toLowerCase();
  const paletteVisible = !paletteDismissed && rawSuggestions.length > 0 && !exactMatch;
  const suggestions = paletteVisible ? rawSuggestions : [];

  useEffect(() => {
    if (selectedIndex >= suggestions.length && suggestions.length > 0) {
      setSelectedIndex(0);
    }
  }, [suggestions.length, selectedIndex]);

  useInput((ch, key) => {
    if (!isActive) return;

    if (key.upArrow) {
      if (suggestions.length > 0 && historyIndex === -1) {
        setSelectedIndex((prev) => Math.max(0, prev - 1));
      } else if (history.length > 0) {
        setHistoryIndex((prev) => {
          const next = prev === -1 ? history.length - 1 : Math.max(0, prev - 1);
          const val = history[next] || '';
          setInput(val);
          setCursorOffset(val.length);
          return next;
        });
      }
    } else if (key.downArrow) {
      if (suggestions.length > 0 && historyIndex === -1) {
        setSelectedIndex((prev) => Math.min(suggestions.length - 1, prev + 1));
      } else if (historyIndex !== -1) {
        setHistoryIndex((prev) => {
          const next = prev + 1;
          if (next >= history.length) {
            setInput('');
            setCursorOffset(0);
            return -1;
          }
          const val = history[next] || '';
          setInput(val);
          setCursorOffset(val.length);
          return next;
        });
      }
    } else if (key.tab || (key.rightArrow && suggestions.length > 0 && suggestions[selectedIndex] && cursorOffset === input.length)) {
      if (suggestions.length > 0 && suggestions[selectedIndex]) {
        const val = suggestions[selectedIndex].command + ' ';
        setInput(val);
        setCursorOffset(val.length);
      }
    } else if (key.leftArrow) {
      setCursorOffset((prev) => Math.max(0, prev - 1));
    } else if (key.rightArrow) {
      setCursorOffset((prev) => Math.min(input.length, prev + 1));
    } else if (key.backspace) {
      if (cursorOffset > 0) {
        setInput((prev) => prev.slice(0, cursorOffset - 1) + prev.slice(cursorOffset));
        setCursorOffset((prev) => prev - 1);
        setSelectedIndex(0);
        setHistoryIndex(-1);
      }
    } else if (key.delete) {
      if (cursorOffset < input.length) {
        setInput((prev) => prev.slice(0, cursorOffset) + prev.slice(cursorOffset + 1));
        setSelectedIndex(0);
        setHistoryIndex(-1);
      }
    } else if (key.return) {
      if (suggestions.length > 0 && suggestions[selectedIndex]) {
        const val = suggestions[selectedIndex].command + ' ';
        setInput(val);
        setCursorOffset(val.length);
        setPaletteDismissed(true);
        return;
      }
      const val = input.trim();
      if (val) {
        setHistory((prev) => {
          const newHist = [...prev];
          if (newHist[newHist.length - 1] !== val) {
            newHist.push(val);
          }
          return newHist;
        });
        setHistoryIndex(-1);
        setPaletteDismissed(false);
        onSubmit(val);
        setInput('');
        setCursorOffset(0);
      }
    } else if (key.escape) {
      if (suggestions.length > 0) {
        setPaletteDismissed(true);
      }
    } else if (key.ctrl && ch === 'l') {
      onClear();
    } else if (key.ctrl && ch === 'c') {
      onExit();
    } else if (ch && !key.ctrl && !key.meta) {
      setInput((prev) => prev.slice(0, cursorOffset) + ch + prev.slice(cursorOffset));
      setCursorOffset((prev) => prev + ch.length);
      setSelectedIndex(0);
      setHistoryIndex(-1);
      setPaletteDismissed(false);
    }
  });

  const renderValue = () => {
    if (input.length === 0) {
      return (
        <Text>
          {chalk.inverse(' ')}
          <Text color={colors.textDim}> Type a command. /help to list all.</Text>
        </Text>
      );
    }
    let rendered = '';
    for (let i = 0; i < input.length; i++) {
      rendered += i === cursorOffset ? chalk.inverse(input[i]!) : input[i];
    }
    if (cursorOffset === input.length) {
      rendered += chalk.inverse(' ');
    }
    return <Text>{rendered}</Text>;
  };

  const { stdout } = useStdout();
  const cols = stdout?.columns ?? 80;

  return (
    <Box flexDirection="column">
      {historyIndex === -1 && <Autocomplete input={input} selectedIndex={selectedIndex} />}
      <Box>
        <Text color={colors.textDim}>{'─'.repeat(Math.max(0, cols))}</Text>
      </Box>
      <Box paddingX={1} width="100%">
        <Box marginRight={1}>
          <Text color={colors.accentLime}>{glyphs.prompt}</Text>
        </Box>
        {isActive ? renderValue() : <Text>{input}</Text>}
      </Box>
    </Box>
  );
};
