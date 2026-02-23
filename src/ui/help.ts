import { usage } from '../utils/usage-help';
import { DodoCliLogo } from './logo';

export const renderHelp = (category?: string): void => {
  console.log(DodoCliLogo);

  if (category && category in usage) {
    const commands = usage[category as keyof typeof usage]!;
    console.log(category);
    commands.forEach(({ command, description }, i) => {
      const isLast = i === commands.length - 1;
      console.log(
        `${isLast ? '└──' : '├──'} dodo ${category} ${command} - ${description}`,
      );
    });
  } else {
    const entries = Object.entries(usage);
    entries.forEach(([categoryName, commands], ci) => {
      const isLastCategory = ci === entries.length - 1;
      console.log(`${isLastCategory ? '└──' : '├──'} ${categoryName}`);
      commands.forEach(({ command, description }, i) => {
        const isLast = i === commands.length - 1;
        const prefix = isLastCategory ? '    ' : '│   ';
        console.log(
          `${prefix}${isLast ? '└──' : '├──'} dodo ${categoryName} ${command} - ${description}`,
        );
      });
      if (!isLastCategory) console.log('│');
    });
  }
};
