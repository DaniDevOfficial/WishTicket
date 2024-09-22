
  export type Colors = keyof typeof Tokens.colors
  export type FontSize = keyof typeof Tokens.fontSizes
  export type Shadows = keyof typeof Tokens.boxShadows

  export type Token = Colors | FontSize | Shadows

  export const Tokens = {
  colors: {
    text: {
      '50': '#16a690',
      '100': '#128775',
      '200': '#0e6759',
      '300': '#0a483e',
      '400': '#062e28',
      '500': '#020f0d',
      '600': '#000000',
      '700': '#000000',
      '800': '#000000',
      '900': '#000000',
      base: '#020f0d',
    },
    background: {
      '50': '#ffffff',
      '100': '#ffffff',
      '200': '#ffffff',
      '300': '#ffffff',
      '400': '#ffffff',
      '500': '#f4fefc',
      '600': '#d3fbf3',
      '700': '#98f6e3',
      '800': '#78f3da',
      '900': '#57f0d1',
      base: '#f4fefc',
    },
    primary: {
      '50': '#b5f6ed',
      '100': '#96f2e5',
      '200': '#76eede',
      '300': '#56ead6',
      '400': '#3de7d0',
      '500': '#1de3c8',
      '600': '#18c4ad',
      '700': '#118b7a',
      '800': '#0d6b5e',
      '900': '#094b42',
      base: '#1de3c8',
    },
    secondary: {
      '50': '#ffffff',
      '100': '#f0f3fd',
      '200': '#d1d8f9',
      '300': '#b1bdf5',
      '400': '#98a7f2',
      '500': '#788cee',
      '600': '#5871ea',
      '700': '#1f40e3',
      '800': '#1936c5',
      '900': '#152da6',
      base: '#788cee',
    },
    accent: {
      '50': '#e6e5fc',
      '100': '#c8c5f8',
      '200': '#aaa5f4',
      '300': '#8b85f0',
      '400': '#736ced',
      '500': '#554ce9',
      '600': '#372ce5',
      '700': '#2017ba',
      '800': '#1b139a',
      '900': '#150f7b',
      base: '#554ce9',
    },
  },
  fontSizes: {},
  fontWeights: {},
  lineHeights: {},
  boxShadows: {},
}
  