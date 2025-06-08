import { ChakraProvider } from '@chakra-ui/react';
import { type FC, type PropsWithChildren } from 'react';
import { I18nextProvider } from 'react-i18next';

// themes
import defaultTheme from '@/theme';

// types
import type { IProps } from './types';

const AppProvider: FC<PropsWithChildren<IProps>> = ({ children, i18n, theme = defaultTheme }) => {
  const innerProviders = (
    <ChakraProvider value={theme}>
      {children}
    </ChakraProvider>
  );

  if (!i18n) {
    return innerProviders;
  }

  return (
    <I18nextProvider i18n={i18n}>
      {innerProviders}
    </I18nextProvider>
  );
};

export default AppProvider;
