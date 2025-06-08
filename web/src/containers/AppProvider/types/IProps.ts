import type { SystemContext as ChakraSystemContext } from '@chakra-ui/react';
import type { i18n as I18n } from 'i18next';

interface IProps {
  i18n?: I18n;
  theme?: ChakraSystemContext;
}

export default IProps;
