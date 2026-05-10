import React, { useState, useRef, useEffect } from 'react';
import { Box, useApp, useInput } from 'ink';
import { WelcomeBanner } from './WelcomeBanner';
import { StatusBar } from './StatusBar';
import { MessageList } from './MessageList';
import { InputBar } from './InputBar';
import { resolveCredentials, setSessionMode } from '../../utils/auth';
import type { Message, BlockType, BlockVariant } from './types';
import type { CommandContext } from './CommandContext';
import { handleCommand } from './router';
import {
  checkForUpdates,
  detectInstallMethod,
  dispatchSilentUpdate,
  consumePendingSilentUpdate,
  type UpdateInfo,
  type InstallMethod,
} from '../../utils/update';
import { version } from '../../../package.json';
import { UpdateNotification } from './UpdateNotification';

export const App = () => {
  const { exit } = useApp();
  const [messages, setMessages] = useState<Message[]>([]);
  const [isProcessing, setIsProcessing] = useState(true);
  const [showBanner, setShowBanner] = useState(true);
  const [isInitialized, setIsInitialized] = useState(false);
  const [authInfo, setAuthInfo] = useState<{ mode: string; key: string } | null>(null);
  const [updateInfo, setUpdateInfo] = useState<UpdateInfo | null>(null);
  const [installMethod] = useState<InstallMethod>(() => detectInstallMethod());

  useEffect(() => {
    const completed = consumePendingSilentUpdate();
    if (completed && completed.to !== completed.from) {
      addBlockToLatestSystemMessage({
        type: 'success',
        message: `Updated to v${completed.to} (was v${completed.from}). Restart to use the new version.`,
      });
    }

    const checkUpdate = async () => {
      const info = await checkForUpdates(version);
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
    };
    checkUpdate();
  }, []);

  useEffect(() => {
    const init = async () => {
      try {
        const creds = await resolveCredentials(ctx);
        setSessionMode(creds.mode);
        const maskedKey = creds.apiKey.slice(0, 10) + '...' + creds.apiKey.slice(-3);
        setAuthInfo({ mode: creds.mode, key: maskedKey });
        setIsInitialized(true);
      } catch (e: any) {
        if (e.message !== 'CONFIG_NOT_FOUND' && !e.message.includes('login')) {
          addBlockToLatestSystemMessage({ type: 'error', message: e.message });
        }
        setAuthInfo(null);
      } finally {
        setIsInitialized(true);
        setIsProcessing(false);
      }
    };
    init();
  }, []);

  useInput((_ch, key) => {
    if (key.escape) {
      exit();
    }
  });


  const addMessage = (role: 'user' | 'system', text?: string) => {
    const id = Math.random().toString(36).substring(7);
    setMessages((prev) => [...prev, { id, role, text, blocks: [] }]);
    return id;
  };

  const addBlockToLatestSystemMessage = (block: BlockVariant) => {
    const id = Math.random().toString(36).substring(7);
    const newBlock = { ...block, id } as BlockType;
    setMessages((prev) => {
      const copy = [...prev];
      let sysMsg = copy[copy.length - 1];
      if (!sysMsg || sysMsg.role !== 'system') {
        sysMsg = { id: Math.random().toString(36).substring(7), role: 'system', blocks: [] };
        copy.push(sysMsg);
      }
      sysMsg.blocks.push(newBlock);
      return copy;
    });
    return id;
  };

  const updateBlock = (id: string, update: Partial<BlockVariant>) => {
    setMessages((prev) => {
      const copy = [...prev];
      for (const msg of copy) {
        const idx = msg.blocks.findIndex((b) => b.id === id);
        if (idx !== -1) {
          msg.blocks[idx] = { ...msg.blocks[idx], ...update } as BlockType;
          break;
        }
      }
      return copy;
    });
  };

  const removeBlock = (id: string) => {
    setMessages((prev) => {
      const copy = [...prev];
      for (const msg of copy) {
        const idx = msg.blocks.findIndex((b) => b.id === id);
        if (idx !== -1) {
          msg.blocks.splice(idx, 1);
          break;
        }
      }
      return copy;
    });
  };

  const promptInput = (label: string, secure?: boolean): Promise<string> => {
    return new Promise((resolve) => {
      const blockId = addBlockToLatestSystemMessage({
        type: 'inline-input',
        label,
        secure,
        onSubmit: (val: string) => {
          removeBlock(blockId);
          addBlockToLatestSystemMessage({ type: 'streaming', text: `${label} ${secure ? '***' : val.trim()}` });
          resolve(val.trim());
        },
      } as any);
    });
  };

  const promptSelect = (label: string, options: { label: string, value: string }[]): Promise<string> => {
    return new Promise((resolve) => {
      const blockId = addBlockToLatestSystemMessage({
        type: 'inline-select',
        label,
        options,
        onSubmit: (val: string) => {
          removeBlock(blockId);
          const selectedLabel = options.find(o => o.value === val)?.label || val;
          addBlockToLatestSystemMessage({ type: 'streaming', text: `${label || '>'} ${selectedLabel}` });
          resolve(val);
        },
      } as any);
    });
  };

  const promptConfirm = (message: string): Promise<boolean> => {
    return new Promise((resolve) => {
      const blockId = addBlockToLatestSystemMessage({
        type: 'confirm',
        message,
        onConfirm: () => {
          removeBlock(blockId);
          addBlockToLatestSystemMessage({ type: 'streaming', text: `${message} Yes` });
          resolve(true);
        },
        onCancel: () => {
          removeBlock(blockId);
          addBlockToLatestSystemMessage({ type: 'streaming', text: `${message} No` });
          resolve(false);
        },
      } as any);
    });
  };

  const ctx: CommandContext = {
    addBlock: addBlockToLatestSystemMessage,
    updateBlock,
    removeBlock,
    promptInput,
    promptSelect,
    promptConfirm,
  };

  const handleSubmit = async (input: string) => {
    if (showBanner) setShowBanner(false);

    if (input === '/clear' || input === 'clear') {
      setMessages([]);
      return;
    }
    if (input === '/exit' || input === '/quit' || input === 'exit' || input === 'quit') {
      exit();
      return;
    }

    addMessage('user', input);
    setIsProcessing(true);

    try {
      await handleCommand(input, ctx, exit);
    } catch (e: any) {
      ctx.addBlock({ type: 'error', message: e.message || String(e) });
    } finally {
      setIsProcessing(false);
    }
  };

  return (
    <Box flexDirection="column" minHeight={10}>
      {showBanner && isInitialized && <WelcomeBanner authMode={authInfo?.mode} />}
      {updateInfo && <UpdateNotification info={updateInfo} method={installMethod} />}
      <MessageList messages={messages} />
      {!showBanner && isInitialized && <StatusBar authInfo={authInfo} />}
      <InputBar
        onSubmit={handleSubmit}
        onClear={() => setMessages([])}
        onExit={exit}
        isActive={!isProcessing}
      />
    </Box>
  );
};
