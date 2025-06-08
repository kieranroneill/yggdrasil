import { type FC } from 'react';

// containers
import AppProvider from '@/containers/AppProvider';

// types
import type { IProps } from './types';

const App: FC<IProps> = ({ i18n }) => {
  return (
    <AppProvider i18n={i18n}>
      <div>Hello humie1!</div>
    </AppProvider>
  );
};

export default App;
