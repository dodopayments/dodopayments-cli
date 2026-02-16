import { rename, rm } from "node:fs/promises";

const targets: any = [
    "bun-darwin-x64",
    "bun-darwin-arm64",
    "bun-linux-x64",
    "bun-linux-arm64",
    "bun-windows-x64"
];

// Loop through all targets and build them
for (const target of targets) {
    const targetName = target.replace('bun-', '');
    const tempDir = `./dist/temp-${targetName}`;

    const build = await Bun.build({
        entrypoints: ["./src/index.ts"],
        outdir: tempDir,
        target: "bun",
        compile: {
            // `as any` cause Bun doesn't provide type of exact platforms
            target: target as any,
        },
    });

    const artifact = build.outputs[0];
    if (build.success && artifact) {
        const binaryPath = artifact.path;

        if (binaryPath) {
            let finalName = `dodo-cli-${targetName}`;
            if (target.includes("windows")) {
                finalName += ".exe";
            }

            await rename(binaryPath, `./dist/${finalName}`);
            await rm(tempDir, { recursive: true, force: true });
            console.log(`Successfully built ${finalName}`);
        } else {
            console.error(`Build successful but artifact path missing for ${target}`);
        }
    }
}

console.log(`Build complete! Binaries are in ./dist`);