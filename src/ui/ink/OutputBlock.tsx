import React, { useState } from 'react';
import { Box, Text, useInput } from 'ink';
import chalk from 'chalk';
import type { BlockType } from './types';
import Spinner from 'ink-spinner';
import SelectInput from 'ink-select-input';
import Markdown from 'ink-markdown-es';

const SimpleTable = ({ data }: { data: any[] }) => {
  if (!data || data.length === 0) return null;
  const keys = Object.keys(data[0]);

  // Calculate dynamic column widths (min 10, max 45)
  const columnWidths = keys.reduce((acc, k) => {
    const maxLen = Math.max(
      k.length,
      ...data.map((row) => String(row[k] ?? '').length)
    );
    acc[k] = Math.min(Math.max(maxLen + 2, 10), 45); // +2 for spacing
    return acc;
  }, {} as Record<string, number>);

  return (
    <Box flexDirection="column" borderStyle="single" borderColor="#07BC70">
      <Box paddingX={1}>
        {keys.map((k) => (
          <Box key={k} width={columnWidths[k]}>
            <Text bold color="#07BC70">{k}</Text>
          </Box>
        ))}
      </Box>
      <Box paddingX={1}>
        {keys.map((k) => (
          <Box key={`dash-${k}`} width={columnWidths[k] ?? 10}>
            <Text color="gray">{'—'.repeat((columnWidths[k] ?? 10) - 1)}</Text>
          </Box>
        ))}
      </Box>
      {data.map((row, i) => (
        <Box key={i} paddingX={1}>
          {keys.map((k) => {
            const valStr = String(row[k] ?? '');
            const w = columnWidths[k] ?? 10;
            // Truncate if exceeds max width - 1
            const displayStr = valStr.length > w - 1 
              ? valStr.substring(0, Math.max(0, w - 4)) + '...' 
              : valStr;
            return (
              <Box key={k} width={w}>
                <Text>{displayStr}</Text>
              </Box>
            );
          })}
        </Box>
      ))}
    </Box>
  );
};

const InlineInput = ({ block }: { block: any }) => {
  const [val, setVal] = useState('');
  const [cursor, setCursor] = useState(0);

  useInput((ch, key) => {
    if (key.return) {
      block.onSubmit(val);
    } else if (key.backspace) {
      if (cursor > 0) {
        setVal((prev) => prev.slice(0, cursor - 1) + prev.slice(cursor));
        setCursor((prev) => prev - 1);
      }
    } else if (key.delete) {
      if (cursor < val.length) {
        setVal((prev) => prev.slice(0, cursor) + prev.slice(cursor + 1));
      }
    } else if (key.leftArrow) {
      setCursor((prev) => Math.max(0, prev - 1));
    } else if (key.rightArrow) {
      setCursor((prev) => Math.min(val.length, prev + 1));
    } else if (ch && !key.ctrl && !key.meta) {
      setVal((prev) => prev.slice(0, cursor) + ch + prev.slice(cursor));
      setCursor((prev) => prev + ch.length);
    }
  });

  const renderValue = () => {
    const displayVal = block.secure ? '*'.repeat(val.length) : val;
    if (displayVal.length === 0) return chalk.inverse(' ');
    let rendered = '';
    for (let i = 0; i < displayVal.length; i++) {
      rendered += i === cursor ? chalk.inverse(displayVal[i]!) : displayVal[i];
    }
    if (cursor === displayVal.length) rendered += chalk.inverse(' ');
    return rendered;
  };

  return (
    <Box borderStyle="round" borderColor="#07BC70" paddingX={1} width="100%">
      <Box marginRight={1}>
        <Text color="#07BC70">{block.label}</Text>
      </Box>
      <Text>{renderValue()}</Text>
    </Box>
  );
};

export const OutputBlock = ({ block }: { block: BlockType }) => {
  switch (block.type) {
    case 'spinner':
      return (
        <Box>
          <Text color="#07BC70"><Spinner type="dots" /></Text>
          <Text> {block.label}</Text>
        </Box>
      );
    case 'table':
      return (
        <Box>
          <SimpleTable data={block.data} />
        </Box>
      );
    case 'detail':
      return (
        <Box flexDirection="column" paddingY={1}>
          {Object.entries(block.data).map(([key, value]) => (
            <Box key={key}>
              <Box width={20} justifyContent="flex-end" paddingRight={1}>
                <Text color="#07BC70">{key}:</Text>
              </Box>
              <Text>{String(value)}</Text>
            </Box>
          ))}
        </Box>
      );
    case 'error':
      return (
        <Box borderStyle="single" borderColor="red" paddingX={1}>
          <Text color="red">{block.message}</Text>
        </Box>
      );
    case 'success':
      return (
        <Box>
          <Text color="#07BC70">✓ </Text>
          <Text>{block.message}</Text>
        </Box>
      );
    case 'link':
      return (
        <Box>
          <Text>To view, go to: </Text>
          <Text color="blueBright" underline>{block.url}</Text>
        </Box>
      );
    case 'empty':
      return (
        <Box>
          <Text color="gray">No results found.</Text>
        </Box>
      );
    case 'streaming':
      return (
        <Box flexDirection="column">
          <Markdown>{block.text}</Markdown>
        </Box>
      );
    case 'event':
      return (
        <Box>
          <Text color="yellow">[{new Date().toLocaleTimeString()}] </Text>
          <Text>{JSON.stringify(block.event)}</Text>
        </Box>
      );
    case 'help':
      return (
        <Box flexDirection="column">
          <Text bold>Commands:</Text>
          <Text color="gray">See autocomplete suggestions for full list.</Text>
        </Box>
      );
    case 'inline-input':
      return <InlineInput block={block} />;
    case 'inline-select':
      return (
        <Box flexDirection="column">
          {block.label && <Text>{block.label}</Text>}
          <SelectInput
            items={block.options}
            onSelect={(item) => block.onSubmit(item.value)}
          />
        </Box>
      );
    case 'confirm':
      return (
        <Box flexDirection="column">
          <Text>{block.message}</Text>
          <Box gap={1}>
            <SelectInput
              items={[
                { label: 'Yes', value: 'yes' },
                { label: 'No', value: 'no' }
              ]}
              onSelect={(item) => item.value === 'yes' ? block.onConfirm() : block.onCancel()}
            />
          </Box>
        </Box>
      );
    default:
      return null;
  }
};
