import { $ } from "bun";

const packages = [
    { name: "@opentui/core-darwin-x64@0.2.3", os: "darwin", cpu: "x64" },
    { name: "@opentui/core-linux-x64@0.2.3", os: "linux", cpu: "x64" },
    { name: "@opentui/core-linux-arm64@0.2.3", os: "linux", cpu: "arm64" },
    { name: "@opentui/core-win32-x64@0.2.3", os: "win32", cpu: "x64" },
];

for (const pkg of packages) {
    console.log(`Fetching ${pkg.name}...`);
    await $`npm install ${pkg.name} --no-save --force --os=${pkg.os} --cpu=${pkg.cpu}`;
}

console.log("Ready for cross-compilation!");