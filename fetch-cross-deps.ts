import { $ } from "bun";
import { mkdir, rename, rm } from "node:fs/promises";
import { join } from "node:path";

const packages = [
    { name: "@opentui/core-darwin-x64@0.2.3", os: "darwin", cpu: "x64" },
    { name: "@opentui/core-linux-x64@0.2.3", os: "linux", cpu: "x64" },
    { name: "@opentui/core-linux-arm64@0.2.3", os: "linux", cpu: "arm64" },
    { name: "@opentui/core-win32-x64@0.2.3", os: "win32", cpu: "x64" },
    { name: "@opentui/core-darwin-arm64@0.2.3", os: "darwin", cpu: "arm64" }
];

const tempDir = join(process.cwd(), "temp-install");
await rm(tempDir, { recursive: true, force: true });
await mkdir(tempDir, { recursive: true });

try {
    for (const pkg of packages) {
        console.log(`Fetching ${pkg.name}...`);
        
        // Install in the temporary directory to prevent npm from pruning other platforms
        await $`npm install ${pkg.name} --force --os=${pkg.os} --cpu=${pkg.cpu} --prefix ${tempDir}`;
        
        const match = pkg.name.match(/@opentui\/([^@]+)/);
        const dirName = match ? match[1] : "";
        if (!dirName) {
            throw new Error(`Could not parse package directory name from: ${pkg.name}`);
        }

        const src = join(tempDir, "node_modules", "@opentui", dirName);
        const dest = join(process.cwd(), "node_modules", "@opentui", dirName);

        await rm(dest, { recursive: true, force: true });
        await mkdir(join(process.cwd(), "node_modules", "@opentui"), { recursive: true });
        await rename(src, dest);
    }
} finally {
    await rm(tempDir, { recursive: true, force: true });
}

console.log("Ready for cross-compilation!");