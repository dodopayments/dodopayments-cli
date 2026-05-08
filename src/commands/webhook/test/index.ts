import Bun from 'bun';

Bun.serve({
    fetch: async (req) => {
        const parseURL = new URL(req.url);
        console.log('Path:', parseURL.pathname);
        console.log('Headers:', req.headers);
        console.log('Body:', await req.json());
        return new Response('YAHOO!', { status: 200 });
    },
    port: 8080
});

console.log('Test server listening on port 8080. Logs all incoming webhooks.');