import type DodoPayments from 'dodopayments';
import type { CommandContext } from '../../tui/CommandContext';
import { HTTP_SCHEME_ERROR, isHttpUrl } from '../../utils/url';

export default async function WebhookListener({
  API_KEY,
  dodoClient,
  ctx,
  endpoint: endpointArg,
}: {
  API_KEY: string;
  dodoClient: DodoPayments;
  ctx: CommandContext;
  endpoint?: string;
}) {
  let endpoint: string;
  if (endpointArg) {
    if (!isHttpUrl(endpointArg)) {
      ctx.addBlock({ type: 'error', message: HTTP_SCHEME_ERROR });
      return;
    }
    endpoint = endpointArg;
  } else {
    endpoint = '';
    while (!endpoint) {
      let value: string;
      try {
        value = await ctx.promptInput('Endpoint URL');
      } catch {
        ctx.addBlock({
          type: 'error',
          message: 'Endpoint URL is required. Usage: dodo wh listen <url>',
        });
        return;
      }
      if (isHttpUrl(value)) {
        endpoint = value;
      } else {
        ctx.addBlock({ type: 'error', message: HTTP_SCHEME_ERROR });
      }
    }
  }

  let targetedEndpoint: string;
  if (process.env.DODO_WH_TEST_SERVER_URL) {
    targetedEndpoint = `https://${process.env.DODO_WH_TEST_SERVER_URL}/`;
  } else {
    targetedEndpoint = 'https://wsserver.dodopayments.tech/';
  }

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Connecting to webhook server…' });

  let doesSetupExist = false;
  const checkSetup = await dodoClient.webhooks.list({ limit: 100 });
  checkSetup.data.forEach((e) => {
    if (e.url === targetedEndpoint) {
      doesSetupExist = true;
    }
  });

  if (!doesSetupExist) {
    await dodoClient.webhooks.create({
      url: targetedEndpoint,
      description:
        'This webhook as been automatically added by Dodo Payments CLI for testing webhooks',
    });
  }

  let wss_url: string;
  if (process.env.DODO_WH_TEST_SERVER_URL) {
    wss_url = `wss://${process.env.DODO_WH_TEST_SERVER_URL}/connect`;
  } else {
    wss_url = 'wss://wsserver.dodopayments.tech/connect';
  }

  const ws = new WebSocket(wss_url, {
    headers: {
      'api-key': API_KEY,
    },
  });

  return new Promise<void>((resolve) => {
    ws.onopen = () => {
      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Connected to webhook server.' });
    };

    ws.onmessage = async (e) => {
      try {
        const data = JSON.parse(e.data.toString());
        if (data.requestId) {
          // Build headers via Headers so casing is normalized and we don't end
          // up with duplicate `Content-Type` entries when `data.headers`
          // already includes one (e.g. `content-type`). Node/undici otherwise
          // concatenates duplicates as `application/json, application/json`,
          // which strict servers like Fastify reject with 415.
          const forwardHeaders = new Headers(data.headers ?? {});
          forwardHeaders.set('Content-Type', 'application/json');

          const response = await fetch(endpoint, {
            method: 'POST',
            headers: forwardHeaders,
            body: JSON.stringify(data.payload),
          });

          const responseText = await response.text();
          let responseBody;
          try {
            responseBody = JSON.parse(responseText);
          } catch (e) {
            responseBody = responseText;
          }

          const responseHeaders: Record<string, string> = {};
          response.headers.forEach((value, key) => {
            responseHeaders[key] = value;
          });

          ctx.addBlock({
            type: 'event',
            event: {
              type: data.payload.type,
              status: response.status,
              response: responseBody
            }
          });

          ws.send(
            JSON.stringify({
              type: 'webhook_response',
              requestId: data.requestId,
              status: response.status,
              body: responseBody,
              headers: responseHeaders,
            }),
          );
        }
      } catch (error) {
        ctx.addBlock({ type: 'error', message: `Couldn't process webhook message. ${String(error)}` });
      }
    };

    ws.onclose = (event) => {
      ctx.addBlock({ type: 'error', message: `Disconnected from webhook server. ${event.reason || ''} (code ${event.code})` });
      resolve();
    };

    ws.onerror = (e) => {
      ctx.addBlock({ type: 'error', message: `Webhook server error. ${(e as any).message || 'Unknown error'}` });
    };
  });
}
