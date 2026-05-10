import { execSync } from 'node:child_process';
import os from 'node:os';
import path from 'node:path';
import { version } from '../../../package.json';
import { colors } from '../theme';

const homedir = os.homedir();

const tildify = (p: string): string => {
  if (p === homedir) return '~';
  if (p.startsWith(homedir + path.sep)) return '~' + p.slice(homedir.length);
  return p;
};

const detectBranch = (): string | null => {
  try {
    const out = execSync('git rev-parse --abbrev-ref HEAD', {
      stdio: ['ignore', 'pipe', 'ignore'],
      timeout: 200,
    })
      .toString()
      .trim();
    return out && out !== 'HEAD' ? out : null;
  } catch {
    return null;
  }
};

const cwd = tildify(process.cwd());
const branch = detectBranch();
const leftLabel = branch ? `${cwd}:${branch}` : cwd;

export const Footer = () => (
  <box flexDirection="row" flexShrink={0} paddingLeft={1} paddingRight={1}>
    <box flexGrow={1} flexDirection="row">
      <text fg={colors.textDim}>{leftLabel}</text>
    </box>
    <text fg={colors.textDim}>{`v${version}`}</text>
  </box>
);
