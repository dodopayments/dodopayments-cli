const targets: any = [
    "bun-darwin-x64",
    "bun-darwin-arm64",
    "bun-linux-x64",
    "bun-linux-arm64",
    "bun-windows-x64"
];

// Loop through all targets and build them
for (const target of targets) {
    await Bun.build({
        entrypoints: ["./index.ts"],
        outdir: `./dist/dodo-cli-${target.replace('bun-', '')}`,
        target: "bun",
        compile: {
            // `as any` cause Bun doesn't provide type of exact platforms
            target: target as any,
        }
    });
}

console.log(`Build complete! Binaries are in ./dist`);