import { createOpenRouter } from '@openrouter/ai-sdk-provider';
import { generateText } from 'ai';
import type { ModelMessage } from 'ai';
import { createMCPClient } from '@ai-sdk/mcp';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';
import { resolveCredentials } from '../../utils/auth';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function getAIModel(ctx: CommandContext) {
  const { apiKey, mode } = await resolveCredentials(ctx, false);
  const proxyMode = mode === 'test_mode' ? 'test' : 'live';

  const openrouter = createOpenRouter({
    baseURL: 'https://ai-proxy.dodopayments.tech/proxy',
    apiKey,
    headers: { 'X-Dodo-Mode': proxyMode },
  });
  return openrouter('openai/gpt-4o');
}

/**
 * Classify an error into a user-friendly category with diagnostic detail.
 */
function classifyError(error: any, phase: string): string {
  const raw = error?.message ?? String(error);
  const status = error?.statusCode ?? error?.response?.status ?? error?.status;
  const code = error?.code ?? error?.cause?.code;

  // Auth errors
  if (status === 401) {
    return 'Authentication failed — your Dodo API key is invalid or expired. Run /login again.';
  }
  if (status === 403) {
    return 'Access denied — your API key does not have permission for AI features.';
  }

  // Rate limiting
  if (status === 429) {
    return 'AI rate limit reached. Please wait a moment and try again.';
  }

  // Network / connection errors
  if (code === 'ECONNREFUSED') {
    return `[${phase}] Connection refused — the AI proxy server at ai-proxy.dodopayments.tech is unreachable. Check your network connection.`;
  }
  if (code === 'ENOTFOUND' || code === 'EAI_AGAIN') {
    return `[${phase}] DNS lookup failed — cannot resolve the AI proxy host. Check your internet connection.`;
  }
  if (code === 'ETIMEDOUT' || code === 'ESOCKETTIMEDOUT') {
    return `[${phase}] Connection timed out — the AI proxy did not respond in time.`;
  }
  if (code === 'ECONNRESET' || raw.includes('socket hang up')) {
    return `[${phase}] Connection was reset by the server. This may be a temporary issue — try again.`;
  }
  if (/connection closed/i.test(raw)) {
    return `[${phase}] The server closed the connection unexpectedly. This usually means the AI proxy rejected the request (invalid API key, server overloaded, or the response was too large). Verify your API key with /login and try again.`;
  }

  // HTTP errors from the proxy/LLM
  if (status && status >= 500) {
    return `[${phase}] Server error (HTTP ${status}) — the AI proxy or upstream LLM returned an internal error. Try again later.`;
  }
  if (status && status >= 400) {
    return `[${phase}] Request rejected (HTTP ${status}): ${raw}`;
  }

  // Parsing / JSON errors
  if (/JSON|Unexpected token|parse/i.test(raw)) {
    return `[${phase}] Failed to parse the AI response — the server returned invalid data. Raw: ${raw.slice(0, 200)}`;
  }

  // MCP-specific errors
  if (raw.includes('spawn') || raw.includes('ENOENT')) {
    return `[${phase}] Failed to start MCP subprocess — 'npx' may not be installed or not in PATH. Ensure Node.js is properly installed.`;
  }
  if (/MCP|transport/i.test(raw)) {
    return `[${phase}] MCP transport error: ${raw.slice(0, 200)}`;
  }

  // Abort / timeout
  if (/abort|AbortError/i.test(raw)) {
    return `[${phase}] Request was aborted: ${raw.slice(0, 200)}`;
  }
  if (/timed? ?out/i.test(raw)) {
    return `[${phase}] ${raw}`;
  }

  // Fallback with as much detail as possible
  const extras = [
    status ? `status=${status}` : null,
    code ? `code=${code}` : null,
    error?.cause ? `cause=${error.cause.message ?? error.cause}` : null,
  ].filter(Boolean).join(', ');

  return `[${phase}] ${raw}${extras ? ` (${extras})` : ''}`;
}

/**
 * Race a promise against a timeout. Rejects with a descriptive error if the deadline expires.
 */
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
    ctx.addBlock({ type: 'error', message: 'Please provide a question. You can also just type your question directly without /ai.' });
    return;
  }

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Connecting to AI...' });

  let knowledgeClient: Awaited<ReturnType<typeof createMCPClient>> | null = null;
  let execClient: Awaited<ReturnType<typeof createMCPClient>> | null = null;

  try {
    // Phase 1: Resolve credentials & create model
    let model;
    try {
      model = await getAIModel(ctx);
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'Auth')), { _classified: true });
    }

    // Phase 2: Initialize MCP clients (with timeouts)
    ctx.updateBlock(spinnerId, { label: 'Starting MCP tools...' });

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

    ctx.updateBlock(spinnerId, { label: 'Starting execution tools...' });

    try {
      execClient = await withTimeout(
        createMCPClient({
          transport: new StdioClientTransport({
            command: 'npx',
            args: ['-y', 'mcp-remote@latest', 'https://mcp.dodopayments.com/sse'],
            env: process.env as Record<string, string>,
            stderr: 'pipe',
          }),
        }),
        MCP_INIT_TIMEOUT,
        'MCP Execution server',
      );
    } catch (e: any) {
      throw Object.assign(new Error(classifyError(e, 'MCP Exec Init')), { _classified: true });
    }

    // Phase 3: Load tools
    ctx.updateBlock(spinnerId, { label: 'Loading tools...' });
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

    ctx.updateBlock(spinnerId, { label: 'Thinking...' });

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
Today is ${new Date().toDateString()}.`;

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

      ctx.updateBlock(spinnerId, { label: `Thinking... (step ${step + 2})` });
    }

    ctx.removeBlock(spinnerId);

    if (finalText) {
      ctx.addBlock({ type: 'streaming', text: finalText });
    } else {
      ctx.addBlock({ type: 'error', message: 'AI did not return a response. Please try rephrasing your question.' });
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
