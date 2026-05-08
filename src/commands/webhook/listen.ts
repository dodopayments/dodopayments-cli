import type DodoPayments from 'dodopayments';
import type { CommandContext } from '../../ui/ink/CommandContext';

export default async function WebhookListener({
  API_KEY,
  dodoClient,
  ctx,
}: {
  API_KEY: string;
  dodoClient: DodoPayments;
  ctx: CommandContext;
}) {
  const endpoint = await ctx.promptInput('Enter the endpoint URL: ');

  let targetedEndpoint: string;
  if (process.env.DODO_WH_TEST_SERVER_URL) {
    targetedEndpoint = `https://${process.env.DODO_WH_TEST_SERVER_URL}/`;
  } else {
    targetedEndpoint = 'https://wsserver.dodopayments.tech/';
  }

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Connecting...' });

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
      ctx.addBlock({ type: 'success', message: 'Successfully connected to Dodo Payments CLI Webhook Server ✅' });
    };

    ws.onmessage = async (e) => {
      try {
        const data = JSON.parse(e.data.toString());
        if (data.requestId) {
          const response = await fetch(endpoint, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              ...data.headers,
            },
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
        ctx.addBlock({ type: 'error', message: `Failed to parse or respond to message: ${String(error)}` });
      }
    };

    ws.onclose = (event) => {
      ctx.addBlock({ type: 'error', message: `Disconnected from Webhook Server. Reason: ${event.reason} Code: ${event.code}` });
      resolve();
    };

    ws.onerror = (e) => {
      ctx.addBlock({ type: 'error', message: `Error in Webhook Server: ${(e as any).message || 'Unknown error'}` });
    };
  });
}
