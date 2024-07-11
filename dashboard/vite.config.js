import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },

  server: {
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://localhost:8081', // El backend Go debería estar corriendo en este puerto
        changeOrigin: true
      }
    }
  },

  preview: {
    port: 8080
  }
})
