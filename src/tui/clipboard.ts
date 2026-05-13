import { spawn } from 'child_process';

const OSC52_PREFIX = '\x1b]52;c;';
const OSC52_SUFFIX = '\x07';

const writeOsc52 = (text: string): boolean => {
  try {
    const b64 = Buffer.from(text, 'utf8').toString('base64');
    process.stdout.write(`${OSC52_PREFIX}${b64}${OSC52_SUFFIX}`);
    return true;
  } catch {
    return false;
  }
};

const trySpawn = (command: string, args: string[], text: string): Promise<boolean> =>
  new Promise((resolve) => {
    try {
      const child = spawn(command, args, { stdio: ['pipe', 'ignore', 'ignore'] });
      let done = false;
      const finish = (ok: boolean) => {
        if (done) return;
        done = true;
        resolve(ok);
      };
      child.on('error', () => finish(false));
      child.on('close', (code) => finish(code === 0));
      child.stdin.on('error', () => finish(false));
      child.stdin.end(text);
      setTimeout(() => finish(false), 1000);
    } catch {
      resolve(false);
    }
  });

const writeNative = async (text: string): Promise<boolean> => {
  if (process.platform === 'darwin') {
    return trySpawn('pbcopy', [], text);
  }
  if (process.platform === 'win32') {
    return trySpawn('clip', [], text);
  }
  if (process.env.WAYLAND_DISPLAY) {
    if (await trySpawn('wl-copy', [], text)) return true;
  }
  if (await trySpawn('xclip', ['-selection', 'clipboard'], text)) return true;
  return trySpawn('xsel', ['--clipboard', '--input'], text);
};

export const copyToClipboard = async (text: string): Promise<boolean> => {
  const osc = writeOsc52(text);
  const native = await writeNative(text);
  return osc || native;
};
