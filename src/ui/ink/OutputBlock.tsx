import React, { useState } from 'react';
import { Box, Text, useInput } from 'ink';
import chalk from 'chalk';
import type { BlockType } from './types';
import Spinner from 'ink-spinner';
import SelectInput from 'ink-select-input';
import Markdown from 'ink-markdown-es';
import { boxes, colors, glyphs } from '../theme';
import { HELP_FOOTER, HELP_GROUPS } from './help-structure';

const MAX_COL_WIDTH = 45;
const MIN_COL_WIDTH = 10;

const NUMERIC_KEY_HINTS = /(amount|price|total|cost|value|count|quantity|cycles|percent|fee)/i;
const DATE_KEY_HINTS = /(created|updated|deleted|timestamp|date|_at\b|on\b)/i;
const CURRENCY_OR_NUMBER = /^[\$€£₹¥]\s*-?\d[\d,]*(\.\d+)?%?$|^-?\d[\d,]*(\.\d+)?%?$/;

const isNumericColumn = (key: string, samples: any[]): boolean => {
  if (DATE_KEY_HINTS.test(key)) return false;
  if (NUMERIC_KEY_HINTS.test(key)) return true;
  let numericHits = 0;
  let nonEmpty = 0;
  for (const v of samples) {
    const s = String(v ?? '').trim().split(/\s/)[0] ?? '';
    if (!s) continue;
    nonEmpty++;
    if (CURRENCY_OR_NUMBER.test(s)) numericHits++;
  }
  return nonEmpty > 0 && numericHits / nonEmpty >= 0.7;
};

const HELP_HEADING_COLORS: Record<string, string> = {
  PRODUCTS: '#7FC4D4',
  PAYMENTS: '#F5A623',
  CUSTOMERS: '#E85BCF',
  DISCOUNTS: '#C6FE1E',
  LICENCES: '#38BDF8',
  ADDONS: '#7FC4D4',
  REFUNDS: '#F5A623',
  CHECKOUT: '#C6FE1E',
  WEBHOOKS: '#E85BCF',
  AI: '#C6FE1E',
  AUTH: '#38BDF8',
  SESSION: '#737470',
};

const SimpleTable = ({ data }: { data: any[] }) => {
  if (!data || data.length === 0) return null;
  const keys = Object.keys(data[0]);

  const columnWidths = keys.reduce((acc, k) => {
    const maxLen = Math.max(
      k.length,
      ...data.map((row) => String(row[k] ?? '').length)
    );
    acc[k] = Math.min(Math.max(maxLen + 2, MIN_COL_WIDTH), MAX_COL_WIDTH);
    return acc;
  }, {} as Record<string, number>);

  const numericFlags: Record<string, boolean> = {};
  for (const k of keys) {
    numericFlags[k] = isNumericColumn(k, data.map((row) => row[k]));
  }

  return (
    <Box flexDirection="column" {...boxes.table}>
      <Box paddingX={1}>
        {keys.map((k) => (
          <Box key={k} width={columnWidths[k]}>
            <Text bold color={colors.accentSky}>{k}</Text>
          </Box>
        ))}
      </Box>
      <Box paddingX={1}>
        {keys.map((k) => (
          <Box key={`dash-${k}`} width={columnWidths[k] ?? MIN_COL_WIDTH}>
            <Text color={colors.textDim}>{'─'.repeat((columnWidths[k] ?? MIN_COL_WIDTH) - 1)}</Text>
          </Box>
        ))}
      </Box>
      {data.map((row, i) => {
        const isAltRow = i % 2 === 1;
        const baseColor = isAltRow ? colors.textMuted : colors.textPrimary;
        return (
          <Box key={i} paddingX={1}>
            {keys.map((k) => {
              const valStr = String(row[k] ?? '');
              const w = columnWidths[k] ?? MIN_COL_WIDTH;
              const displayStr = valStr.length > w - 1
                ? valStr.substring(0, Math.max(0, w - 2)) + '…'
                : valStr;
              const cellColor = numericFlags[k] ? colors.accentAmber : baseColor;
              return (
                <Box key={k} width={w}>
                  <Text color={cellColor}>{displayStr}</Text>
                </Box>
              );
            })}
          </Box>
        );
      })}
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
    <Box {...boxes.prompt} paddingX={1} width="100%">
      <Box marginRight={1}>
        <Text color={colors.accentMagenta}>{block.label}</Text>
      </Box>
      <Text>{renderValue()}</Text>
    </Box>
  );
};

const HelpPanel = () => {
  const headingWidth = 12;
  const commandWidth = 22;

  return (
    <Box flexDirection="column">
      <Box marginBottom={1}>
        <Text color={colors.brandLime} bold>{glyphs.bullet} </Text>
        <Text color={colors.textPrimary} bold>Dodo Payments CLI</Text>
      </Box>
      {HELP_GROUPS.map((group, gi) => {
        const headingColor = HELP_HEADING_COLORS[group.heading] ?? colors.textMuted;
        return (
          <Box key={group.heading} flexDirection="column" marginBottom={gi === HELP_GROUPS.length - 1 ? 1 : 0}>
            {group.items.map((item, ii) => (
              <Box key={item.command}>
                <Box width={headingWidth}>
                  <Text color={headingColor} bold>{ii === 0 ? group.heading : ''}</Text>
                </Box>
                <Box width={commandWidth}>
                  <Text color={colors.textPrimary}>{item.command}</Text>
                </Box>
                <Box width={10}>
                  <Text color={colors.textDim}>{item.args ?? ''}</Text>
                </Box>
                <Text color={colors.textMuted}>{item.description}</Text>
              </Box>
            ))}
          </Box>
        );
      })}
      <Box>
        <Text color={colors.textDim}>{HELP_FOOTER}</Text>
      </Box>
    </Box>
  );
};

export const OutputBlock = ({ block }: { block: BlockType }) => {
  switch (block.type) {
    case 'spinner':
      return (
        <Box>
          <Text color={colors.accentAmber}><Spinner type={glyphs.spinnerType} /></Text>
          <Text color={colors.textMuted}> {block.label}</Text>
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
          {Object.entries(block.data).map(([key, value]) => {
            const valStr = String(value);
            const firstToken = valStr.trim().split(/\s/)[0] ?? '';
            const isDate = DATE_KEY_HINTS.test(key);
            const isNumeric = !isDate && (NUMERIC_KEY_HINTS.test(key) || CURRENCY_OR_NUMBER.test(firstToken));
            const valueColor = isNumeric ? colors.accentAmber : colors.textPrimary;
            return (
              <Box key={key}>
                <Box width={20} justifyContent="flex-end" paddingRight={1}>
                  <Text color={colors.textMuted}>{key}</Text>
                </Box>
                <Text color={colors.textDim}>{glyphs.separator} </Text>
                <Text color={valueColor}>{valStr}</Text>
              </Box>
            );
          })}
        </Box>
      );
    case 'error':
      return (
        <Box {...boxes.error} paddingX={1}>
          <Text color={colors.error}>{glyphs.cross} </Text>
          <Text color={colors.textPrimary}>{block.message}</Text>
        </Box>
      );
    case 'success':
      return (
        <Box>
          <Text color={colors.success}>{glyphs.check} </Text>
          <Text color={colors.textPrimary}>{block.message}</Text>
        </Box>
      );
    case 'link':
      return (
        <Box>
          <Text color={colors.textMuted}>{block.text} </Text>
          <Text color={colors.info} underline>{block.url}</Text>
        </Box>
      );
    case 'empty':
      return (
        <Box>
          <Text color={colors.textDim}>{glyphs.dot} No results found.</Text>
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
          <Text color={colors.testMode}>[{new Date().toLocaleTimeString()}] </Text>
          <Text color={colors.textPrimary}>{JSON.stringify(block.event)}</Text>
        </Box>
      );
    case 'help':
      return <HelpPanel />;
    case 'inline-input':
      return <InlineInput block={block} />;
    case 'inline-select':
      return (
        <Box flexDirection="column">
          {block.label && <Text color={colors.textMuted}>{block.label}</Text>}
          <SelectInput
            items={block.options}
            onSelect={(item) => block.onSubmit(item.value)}
          />
        </Box>
      );
    case 'confirm':
      return (
        <Box flexDirection="column">
          <Text color={colors.textPrimary}>{block.message}</Text>
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
