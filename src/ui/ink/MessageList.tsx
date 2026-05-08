import React from 'react';
import { Box, Text } from 'ink';
import type { Message } from './types';
import { OutputBlock } from './OutputBlock';
import { colors, glyphs } from '../theme';

interface MessageListProps {
  messages: Message[];
}

export const MessageList = ({ messages }: MessageListProps) => {
  return (
    <Box flexDirection="column" paddingX={1}>
      {messages.map((msg) => (
        <Box key={msg.id} flexDirection="column" marginBottom={1}>
          {msg.role === 'user' ? (
            <Text color={colors.textMuted}>{glyphs.prompt} {msg.text}</Text>
          ) : null}
          {msg.blocks.map((block, i) => (
            <OutputBlock key={block.id || i} block={block} />
          ))}
        </Box>
      ))}
    </Box>
  );
};
