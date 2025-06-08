import { createSystem, defineConfig, defaultConfig } from '@chakra-ui/react';

export default createSystem(
  defaultConfig,
  defineConfig({
    globalCss: {
      '::-webkit-scrollbar': {
        display: 'none',
      },
      html: {
        bg: 'bg',
        colorPalette: {
          _dark: 'primaryDark',
          _light: 'primaryLight',
        },
        lineHeight: '1.5',
        scrollbarWidth: 'none',
      },
    },
    theme: {
      tokens: {
        colors: {
          algorand: {
            50:  { value: '#000000' },
            100: { value: '#000000' },
            200: { value: '#000000' },
            300: { value: '#000000' },
            400: { value: '#1a1a1a' },
            500: { value: '#333333' },
            600: { value: '#4d4d4d' },
            700: { value: '#666666' },
            800: { value: '#ffffff' },
            900: { value: '#ffffff' },
          },
          primaryDark: {
            50:  { value: '#F6E9FF' },
            100: { value: '#F2DEFF' },
            200: { value: '#EED3FF' },
            300: { value: '#E9C8FF' },
            400: { value: '#E5BDFF' },
            500: { value: '#E0B0FF' }, // mauve
            600: { value: '#C875FF' },
            700: { value: '#AF37FF' },
            800: { value: '#9500F8' },
            900: { value: '#6F00BA' },
          },
          primaryLight: {
            50:  { value: '#F59CFD' },
            100: { value: '#F16AFD' },
            200: { value: '#EC39FC' },
            300: { value: '#E707FB' },
            400: { value: '#BC03CD' },
            500: { value: '#8D029B' }, // mauveine
            600: { value: '#7B0285' },
            700: { value: '#66026F' },
            800: { value: '#520159' },
            900: { value: '#3D0143' },
          },
          voi: {
            50:  { value: '#d9c7f7' },
            100: { value: '#bb9af1' },
            200: { value: '#ac84ee' },
            300: { value: '#9d6deb' },
            400: { value: '#8e57e8' },
            500: { value: '#702ae2' },
            600: { value: '#591abf' },
            700: { value: '#4e17a9' },
            800: { value: '#441492' },
            900: { value: '#2f0e65' },
          },
        },
        fonts: {
          heading: {
            value: '"Nunito", sans-serif',
          },
          body: {
            value: '"Nunito", sans-serif',
          },
          mono: {
            value: '"SourceCodePro", monospace',
          },
        },
      },
    },
  })
);
