const path = require('path')
const { defineConfig } = require('vite')
import dts from 'vite-plugin-dts'
import { visualizer } from 'rollup-plugin-visualizer'

module.exports = defineConfig({
  plugins: [
    dts({
      insertTypesEntry: true,
    }),
    visualizer({
      gzipSize: true,
      open: true,
      brotliSize: true,
      template: 'treemap',
    }),
  ],
  build: {
    target: 'esnext',
    sourcemap: true,
    lib: {
      entry: path.resolve(__dirname, 'src/index.ts'),
      name: 'MyComponentLib',
      fileName: (format: string) => `gogoracer-headlamp-lib.${format}.js`,
    },
    rollupOptions: {
      // make sure to externalize deps that shouldn't be bundled
      // into your library
      external: ['vue'],
      output: {
        // Provide global variables to use in the UMD build
        // for externalized deps
        globals: {
          vue: 'Vue',
        },
      },
    },
  },
})
