import { input } from '@inquirer/prompts';
import type DodoPayments from 'dodopayments';

export default async function WebhookListener({
    API_KEY,
    dodoClient
}: {
    API_KEY: string,
    dodoClient: DodoPayments
}) {
    const endpoint = await input({
        message: 'Enter the endpoint URL: ',
    });

    let targetedEndpoint: string;
    if (process.env.DODO_WH_TEST_SERVER_URL) {
        targetedEndpoint = `https://${process.env.DODO_WH_TEST_SERVER_URL}/`
    } else {
        targetedEndpoint = 'https://wsserver.dodopayments.tech/'
    }

    // This will check if the webhook is already added. If not added, it will add it.
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
            description: 'This webhook as been automatically added by Dodo Payments CLI for testing webhooks',
        });
    }


    let wss_url: string;
    if (process.env.DODO_WH_TEST_SERVER_URL) {
        wss_url = `wss://${process.env.DODO_WH_TEST_SERVER_URL}/connect`
    } else {
        wss_url = 'wss://wsserver.dodopayments.tech/connect'
    }

    const ws = new WebSocket(wss_url, {
        headers: {
            'api-key': API_KEY
        }
    });

    ws.onopen = () => {
        console.log('Successfully connected to Dodo Payments CLI Webhook Server ✅');
    };

    ws.onmessage = async (e) => {
        try {
            const data = JSON.parse(e.data.toString());
            if (data.requestId) {
                console.log('Received webhook!');
                console.log('Event Type:', data.payload.type);
                const response = await fetch(endpoint, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        // retrive headers from the request
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

                console.log('Status:', response.status);
                console.log('Response:', responseBody);
                console.log(); // Blank space

                ws.send(JSON.stringify({
                    type: 'webhook_response',
                    requestId: data.requestId,
                    status: response.status,
                    body: responseBody,
                    headers: responseHeaders
                }));
            }
        } catch (error) {
            console.error('Failed to parse or respond to message', error);
        }
    };

    ws.onclose = (event) => {
        console.log('Disconnected from Dodo Payments CLI Webhook Server');
        console.log('Reason:', event.reason);
        console.log('Code:', event.code);
    };

    ws.onerror = (e) => {
        console.log('Error in Dodo Payments CLI Webhook Server:', e);
    };
}