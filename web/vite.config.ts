import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
import { resolve } from 'node:path';
import tsconfigPaths from 'vite-tsconfig-paths';

export default defineConfig({
  base: '/app',
  build: {
    emptyOutDir: true,
    outDir: resolve(__dirname, 'dist'),
  },
  plugins: [
    react(),
    tsconfigPaths({
      configNames: ['tsconfig.json'],
    }),
  ],
  resolve: {
    alias: {
      // typescript resolutions handled by the vite-tsconfig-paths plugin - assets must be handled explicitly
      '@/fonts': resolve(__dirname, 'src/fonts'),
      '@/styles': resolve(__dirname, 'src/styles'),
    },
  },
  server: {
    port: 8080,
  },
});
