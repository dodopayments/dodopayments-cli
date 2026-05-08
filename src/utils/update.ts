import fs from 'node:fs';
import path from 'node:path';
import os from 'node:os';

const UPDATE_CACHE_PATH = path.join(os.homedir(), '.dodopayments', 'update-cache.json');

export interface UpdateCache {
  latestVersion: string;
  lastChecked: number;
}

/**
 * Checks for updates in the background.
 * Returns the latest version if an update is available, null otherwise.
 */
export async function checkForUpdates(currentVersion: string): Promise<string | null> {
  try {
    // 1. Ensure config directory exists
    const configDir = path.dirname(UPDATE_CACHE_PATH);
    if (!fs.existsSync(configDir)) {
      fs.mkdirSync(configDir, { recursive: true });
    }

    // 2. Read cache
    let cache: UpdateCache | null = null;
    if (fs.existsSync(UPDATE_CACHE_PATH)) {
      try {
        cache = JSON.parse(fs.readFileSync(UPDATE_CACHE_PATH, 'utf-8'));
      } catch (e) {
        // Corrupted cache, ignore
      }
    }

    const now = Date.now();
    const SIX_HOURS = 6 * 60 * 60 * 1000;

    // 3. If cache is old or missing, check NPM
    if (!cache || now - cache.lastChecked > SIX_HOURS) {
      try {
        const response = await fetch('https://registry.npmjs.org/dodopayments-cli/latest', {
          signal: AbortSignal.timeout(2000), // Don't hang the CLI for too long
        });
        
        if (response.ok) {
          const data = await response.json() as { version: string };
          const latestVersion = data.version;
          
          cache = {
            latestVersion,
            lastChecked: now,
          };
          
          fs.writeFileSync(UPDATE_CACHE_PATH, JSON.stringify(cache));
        }
      } catch (e) {
        // Network error or timeout, return cached value if we have it
      }
    }

    if (cache && isNewerVersion(cache.latestVersion, currentVersion)) {
      return cache.latestVersion;
    }
  } catch (e) {
    // Fallback for any unexpected errors
  }

  return null;
}

function isNewerVersion(latest: string, current: string): boolean {
  try {
    const l = latest.split('.').map((n) => parseInt(n, 10));
    const c = current.split('.').map((n) => parseInt(n, 10));
    
    for (let i = 0; i < 3; i++) {
      const lNum = l[i] || 0;
      const cNum = c[i] || 0;
      if (lNum > cNum) return true;
      if (lNum < cNum) return false;
    }
  } catch (e) {
    // If parsing fails, just compare as strings
    return latest !== current;
  }
  return false;
}
