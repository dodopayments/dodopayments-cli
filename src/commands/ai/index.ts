import { createOpenRouter } from '@openrouter/ai-sdk-provider';
import { generateText } from 'ai';
import type { ModelMessage } from 'ai';
import { createMCPClient } from '@ai-sdk/mcp';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';
import { resolveCredentials } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

export async function getAIModel(ctx: CommandContext) {
  const { apiKey, mode } = await resolveCredentials(ctx, false);
  return buildModel(apiKey, mode);
}

function buildModel(apiKey: string, mode: 'test_mode' | 'live_mode') {
  const proxyMode = mode === 'test_mode' ? 'test' : 'live';
  const openrouter = createOpenRouter({
    baseURL: 'https://ai-proxy.dodopayments.tech/proxy',
    apiKey,
    headers: { 'X-Dodo-Mode': proxyMode },
  });
  return openrouter('openai/gpt-5.4-mini');
}

function classifyError(error: any, phase: string): string {
  const raw = error?.message ?? String(error);
  const status = error?.statusCode ?? error?.response?.status ?? error?.status;
  const code = error?.code ?? error?.cause?.code;

  if (status === 401) {
    return 'Authentication failed. Your API key is invalid or expired. Run /login to refresh.';
  }
  if (status === 403) {
    return 'Access denied. Your API key does not have permission for the assistant.';
  }
  if (status === 429) {
    return 'Rate limit reached. Wait a moment and try again.';
  }

  if (code === 'ECONNREFUSED') {
    return `Couldn't reach the assistant. Check your network and try again. (${phase}: connection refused)`;
  }
  if (code === 'ENOTFOUND' || code === 'EAI_AGAIN') {
    return `Couldn't reach the assistant. Check your network and try again. (${phase}: DNS lookup failed)`;
  }
  if (code === 'ETIMEDOUT' || code === 'ESOCKETTIMEDOUT') {
    return `Couldn't reach the assistant. Try again in a moment. (${phase}: timeout)`;
  }
  if (code === 'ECONNRESET' || raw.includes('socket hang up')) {
    return `Connection reset. Try again in a moment. (${phase})`;
  }
  if (/connection closed/i.test(raw)) {
    return `The server closed the connection. Verify your API key with /login and try again. (${phase})`;
  }

  if (status && status >= 500) {
    return `Server error (HTTP ${status}). Try again in a moment. (${phase})`;
  }
  if (status && status >= 400) {
    return `Request rejected (HTTP ${status}). ${raw} (${phase})`;
  }

  if (/JSON|Unexpected token|parse/i.test(raw)) {
    return `The assistant returned invalid data. Try again. (${phase})`;
  }

  if (raw.includes('spawn') || raw.includes('ENOENT')) {
    return `Couldn't start the assistant. Make sure 'npx' is installed and in your PATH. (${phase})`;
  }
  if (/MCP|transport/i.test(raw)) {
    return `Communication error: ${raw.slice(0, 200)} (${phase})`;
  }

  if (/abort|AbortError/i.test(raw)) {
    return `Request aborted: ${raw.slice(0, 200)} (${phase})`;
  }
  if (/timed? ?out/i.test(raw)) {
    return `Timeout. ${raw} (${phase})`;
  }

  const extras = [
    status ? `status=${status}` : null,
    code ? `code=${code}` : null,
    error?.cause ? `cause=${error.cause.message ?? error.cause}` : null,
  ].filter(Boolean).join(', ');

  return `${raw}${extras ? ` (${extras})` : ''} (${phase})`;
}

function withTimeout<T>(promise: Promise<T>, ms: number, label: string): Promise<T> {
  return new Promise<T>((resolve, reject) => {
    const timer = setTimeout(
      () => reject(new Error(`Timed out after ${Math.round(ms / 1000)}s waiting for ${label}`)),
      ms,
    );
    promise.then(
      (v) => { clearTimeout(timer); resolve(v); },
      (e) => { clearTimeout(timer); reject(e); },
    );
  });
}

const MCP_INIT_TIMEOUT = 60_000; // 60 seconds for MCP subprocess to start
const LLM_REQUEST_TIMEOUT = 60_000; // 60 seconds per LLM step

export async function handleAI(query: string, ctx: CommandContext) {
  if (!query) {
    ctx.addBlock({ type: 'error', message: 'Question required. You can type it directly without /ai.' });
    return;
  }

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Initializing assistant…' });

  let knowledgeClient: Awaited<ReturnType<typeof createMCPClient>> | null = null;
  let execClient: Awaited<ReturnType<typeof createMCPClient>> | null = null;

  try {
    // Phase 1: Resolve credentials & create model
    let model;
    let apiKey: string;
    let mode: 'test_mode' | 'live_mode';
    try {
      const creds = await resolveCredentials(ctx, false);
      apiKey = creds.apiKey;
      mode = creds.mode;
      model = buildModel(apiKey, mode);
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'Auth')), { _classified: true });
    }

    // Phase 2: Initialize MCP clients (with timeouts)
    ctx.updateBlock(spinnerId, { label: 'Loading knowledge base…' });

    try {
      knowledgeClient = await withTimeout(
        createMCPClient({
          transport: new StdioClientTransport({
            command: 'npx',
            args: ['-y', 'mcp-remote@latest', 'https://knowledge.dodopayments.com/mcp'],
            env: process.env as Record<string, string>,
            stderr: 'pipe',
          }),
        }),
        MCP_INIT_TIMEOUT,
        'MCP Knowledge server',
      );
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'MCP Knowledge Init')), { _classified: true });
    }

    ctx.updateBlock(spinnerId, { label: 'Loading workspace…' });

    try {
      execClient = await withTimeout(
        createMCPClient({
          transport: new StdioClientTransport({
            command: 'npx',
            args: ['-y', 'dodopayments-mcp'],
            env: {
              ...(process.env as Record<string, string>),
              DODO_PAYMENTS_API_KEY: apiKey,
              DODO_PAYMENTS_ENVIRONMENT: mode,
            },
            stderr: 'pipe',
          }),
        }),
        MCP_INIT_TIMEOUT,
        'MCP Execution server',
      );
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'MCP Exec Init')), { _classified: true });
    }

    ctx.updateBlock(spinnerId, { label: 'Preparing context…' });
    let kTools, eTools;
    try {
      [kTools, eTools] = await withTimeout(
        Promise.all([knowledgeClient.tools(), execClient.tools()]),
        MCP_INIT_TIMEOUT,
        'MCP tool discovery',
      );
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'MCP Tool Loading')), { _classified: true });
    }

    // Merge tools, prefixing knowledge tools that collide with execution tools
    const tools: Record<string, any> = { ...eTools };
    for (const [key, value] of Object.entries(kTools)) {
      if (tools[key]) {
        tools[`knowledge_${key}`] = value;
      } else {
        tools[key] = value;
      }
    }

    ctx.updateBlock(spinnerId, { label: 'Analyzing your question…' });

    const systemPrompt = `You are a helpful assistant for Dodo Payments.
Answer questions about the user's payments, revenue, customers, and subscriptions.
You have access to tools:
- search_docs / knowledge_search_docs: Search the Dodo Payments SDK documentation
- execute: Execute TypeScript code against the Dodo Payments SDK to query real data

When asked about data (revenue, payments, customers, etc.), use the "execute" tool to write and run code.
For example, to list recent payments, write code like:
  const payments = await client.payments.list({ page_size: 50 });
  return payments;

The "execute" tool has access to an already-configured "client" (DodoPayments SDK instance).
Always provide a final text answer to the user. Do not stop after tool calls without responding.
Today is ${new Date().toDateString()}. Do not accept questions that are not related to Dodo Payments.`;

    // Phase 4: Multi-step LLM loop
    const MAX_STEPS = 20;
    let messages: ModelMessage[] = [
      { role: 'system', content: systemPrompt },
      { role: 'user', content: query },
    ];

    let finalText = '';

    for (let step = 0; step < MAX_STEPS; step++) {
      let result;
      try {
        result = await withTimeout(
          generateText({
            model,
            tools,
            messages,
          }),
          LLM_REQUEST_TIMEOUT,
          `LLM response (step ${step + 1})`,
        );
      } catch (e: any) {
        throw Object.assign(
          new Error(classifyError(e, step === 0 ? 'LLM Request' : `LLM Step ${step + 1}`)),
          { _classified: true },
        );
      }

      // Accumulate any text the model produced
      if (result.text) {
        finalText += result.text;
      }

      // If no tool calls, we're done
      if (result.finishReason !== 'tool-calls' || !result.toolCalls || result.toolCalls.length === 0) {
        break;
      }

      // Append the assistant + tool result messages from the SDK (properly formatted)
      messages = [...messages, ...(result.response.messages as ModelMessage[])];

      ctx.updateBlock(spinnerId, { label: `Analyzing… (step ${step + 2})` });
    }

    ctx.removeBlock(spinnerId);

    if (finalText) {
      ctx.addBlock({ type: 'markdown', text: finalText });
    } else {
      ctx.addBlock({ type: 'error', message: 'No response from the assistant. Try rephrasing your question.' });
    }
  } catch (e: any) {
    ctx.removeBlock(spinnerId);
    const message = (e as any)?._classified ? e.message : classifyError(e, 'Unknown');
    ctx.addBlock({ type: 'error', message });
  } finally {
    try { await knowledgeClient?.close(); } catch { }
    try { await execClient?.close(); } catch { }
  }
}
