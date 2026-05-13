import { SyntaxStyle } from '@opentui/core';
import { colors } from '../../theme';

let cachedStyle: SyntaxStyle | null = null;

const getMarkdownStyle = (): SyntaxStyle => {
  if (cachedStyle) return cachedStyle;
  cachedStyle = SyntaxStyle.fromStyles({
    default: { fg: colors.textPrimary },

    'markup.heading': { fg: colors.brandLime, bold: true },
    'markup.heading.1': { fg: colors.brandLime, bold: true },
    'markup.heading.2': { fg: colors.accentLime, bold: true },
    'markup.heading.3': { fg: colors.accentSky, bold: true },
    'markup.heading.4': { fg: colors.accentCyan, bold: true },
    'markup.heading.5': { fg: colors.accentMagenta, bold: true },
    'markup.heading.6': { fg: colors.textMuted, bold: true },

    'markup.italic': { fg: colors.textPrimary, italic: true },
    'markup.strong': { fg: colors.textPrimary, bold: true },
    'markup.strikethrough': { fg: colors.textDim, dim: true },

    'markup.raw': { fg: colors.accentLime, bg: colors.brandBlack },
    'markup.raw.block': { fg: colors.textPrimary, bg: colors.brandBlack },

    'markup.link': { fg: colors.info },
    'markup.link.label': { fg: colors.info, underline: true },
    'markup.link.url': { fg: colors.accentSky, underline: true },
    'markup.link.bracket.close': { fg: colors.info },

    'markup.list': { fg: colors.accentLime },
    'markup.list.checked': { fg: colors.success },
    'markup.list.unchecked': { fg: colors.textMuted },

    'markup.quote': { fg: colors.textMuted, italic: true },

    'punctuation.special': { fg: colors.textDim },
    'punctuation.delimiter': { fg: colors.textDim },
    'punctuation.bracket': { fg: colors.textDim },

    keyword: { fg: colors.accentMagenta, bold: true },
    'keyword.function': { fg: colors.accentMagenta, bold: true },
    'keyword.return': { fg: colors.accentMagenta, bold: true },
    'keyword.conditional': { fg: colors.accentMagenta, bold: true },
    'keyword.conditional.ternary': { fg: colors.accentMagenta },
    'keyword.repeat': { fg: colors.accentMagenta, bold: true },
    'keyword.exception': { fg: colors.error },
    'keyword.import': { fg: colors.accentMagenta, bold: true },
    'keyword.modifier': { fg: colors.accentMagenta },
    'keyword.operator': { fg: colors.accentMagenta },
    'keyword.coroutine': { fg: colors.accentMagenta },
    'keyword.type': { fg: colors.accentSky },
    'keyword.directive': { fg: colors.accentMagenta },

    string: { fg: colors.brandLime },
    'string.escape': { fg: colors.accentAmber },
    'string.regexp': { fg: colors.accentAmber },
    'string.special': { fg: colors.accentAmber },
    'string.special.url': { fg: colors.info, underline: true },

    number: { fg: colors.accentAmber },
    boolean: { fg: colors.accentAmber, bold: true },

    function: { fg: colors.accentSky },
    'function.call': { fg: colors.accentSky },
    'function.method': { fg: colors.accentSky },
    'function.method.call': { fg: colors.accentSky },
    'function.builtin': { fg: colors.accentCyan },

    constructor: { fg: colors.accentSky, bold: true },

    constant: { fg: colors.accentAmber },
    'constant.builtin': { fg: colors.accentAmber, bold: true },

    variable: { fg: colors.textPrimary },
    'variable.builtin': { fg: colors.accentCyan, italic: true },
    'variable.member': { fg: colors.textPrimary },
    'variable.parameter': { fg: colors.accentLime, italic: true },

    type: { fg: colors.accentSky },
    'type.builtin': { fg: colors.accentSky, italic: true },

    property: { fg: colors.textPrimary },
    attribute: { fg: colors.accentLime },
    module: { fg: colors.accentCyan },
    'module.builtin': { fg: colors.accentCyan, italic: true },
    label: { fg: colors.accentAmber },
    operator: { fg: colors.textMuted },
    comment: { fg: colors.textDim, italic: true },
    'comment.documentation': { fg: colors.textDim, italic: true },

    'character.special': { fg: colors.accentMagenta },
    embedded: { fg: colors.textPrimary },
  });
  return cachedStyle;
};

export const Markdown = (props: { text: string; streaming?: boolean }) => (
  <markdown
    content={props.text}
    syntaxStyle={getMarkdownStyle()}
    streaming={props.streaming ?? false}
    fg={colors.textPrimary}
  />
);
