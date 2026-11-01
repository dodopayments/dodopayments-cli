import os from 'node:os';
import path from 'node:path';
import fs from 'node:fs';
import crypto from 'node:crypto';
import { machineIdSync } from 'node-machine-id';

export const CONFIG_DIR = path.join(os.homedir(), '.dodopayments');
const ALGORITHM = 'aes-256-gcm';

let cachedKey: Buffer | null = null;

function getEncryptionKey(): Buffer {
  if (cachedKey) return cachedKey;
  let id: string;
  try {
    id = machineIdSync();
  } catch {
    id = `${os.hostname()}-${os.userInfo().username}`;
  }
  cachedKey = crypto.pbkdf2Sync(id, 'dodopayments-cli-salt', 100000, 32, 'sha256');
  return cachedKey;
}

export function encrypt(text: string): string {
  const iv = crypto.randomBytes(16);
  const cipher = crypto.createCipheriv(ALGORITHM, getEncryptionKey(), iv);
  let encrypted = cipher.update(text, 'utf8', 'hex');
  encrypted += cipher.final('hex');
  const authTag = cipher.getAuthTag().toString('hex');
  return `${iv.toString('hex')}:${authTag}:${encrypted}`;
}

export function decrypt(text: string): string {
  const parts = text.split(':');
  if (parts.length !== 3) throw new Error('Invalid encrypted format');
  const [ivHex, tagHex, encrypted] = parts;
  if (!ivHex || !tagHex || !encrypted) throw new Error('Invalid encrypted format');
  const iv = Buffer.from(ivHex, 'hex');
  const authTag = Buffer.from(tagHex, 'hex');

  const decipher = crypto.createDecipheriv(ALGORITHM, getEncryptionKey(), iv);
  decipher.setAuthTag(authTag);
  let decrypted = decipher.update(encrypted, 'hex', 'utf8');
  decrypted += decipher.final('utf8');
  return decrypted;
}

export function ensureConfigDir() {
  if (!fs.existsSync(CONFIG_DIR)) {
    fs.mkdirSync(CONFIG_DIR, { recursive: true, mode: 0o700 });
  }
}

/** Writes an encrypted file that only the user can read. */
export function writeEncrypted(filePath: string, text: string) {
  ensureConfigDir();
  fs.writeFileSync(filePath, encrypt(text), { mode: 0o600 });
  try {
    // The mode of an existing file does not change on write.
    fs.chmodSync(filePath, 0o600);
  } catch {
    // Windows has no file modes.
  }
}
