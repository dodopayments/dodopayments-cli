import { standaloneUsage, usage } from '../utils/usage-help';
import { DodoCliLogo } from './logo';

export const renderHelp = (category?: string): void => {
  console.log(DodoCliLogo);

  if (category && category in usage) {
    const commands = usage[category as keyof typeof usage]!;

    console.log(category);
    commands.forEach(({ command, description }, index) => {
      const prefix = index === commands.length - 1 ? '`--' : '|--';
      console.log(`${prefix} dodo ${category} ${command} - ${description}`);
    });

    return;
  }

  standaloneUsage.forEach(({ command, description }) => {
    console.log(`|-- dodo ${command} - ${description}`);
  });

  if (standaloneUsage.length > 0) {
    console.log('|');
  }

  const entries = Object.entries(usage);
  entries.forEach(([categoryName, commands], categoryIndex) => {
    const isLastCategory = categoryIndex === entries.length - 1;

    console.log(`${isLastCategory ? '`--' : '|--'} ${categoryName}`);

    commands.forEach(({ command, description }, commandIndex) => {
      const prefix = isLastCategory ? '    ' : '|   ';
      const branch = commandIndex === commands.length - 1 ? '`--' : '|--';

      console.log(
        `${prefix}${branch} dodo ${categoryName} ${command} - ${description}`,
      );
    });

    if (!isLastCategory) {
      console.log('|');
    }
  });
};
