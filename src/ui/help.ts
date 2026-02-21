import { usage } from "../utils/usage-help";
import { DodoCliLogo } from "./logo";


export const renderHelp = (): void => {
  console.log(DodoCliLogo);

  for (const [category, commands] of Object.entries(usage)) {
    console.log(`Category: ${category}`);

    for (const { command, description } of commands) {
      console.log(`dodo ${category} ${command} - ${description}`);
    }

    console.log("");
  }
};