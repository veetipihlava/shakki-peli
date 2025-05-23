import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  server: {
    host: '0.0.0.0',
    port: 5173,
    strictPort: true
  },
  plugins: [react()],
  server: {
    proxy: {
      '/game': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
      '/game/': {
        target: 'ws://localhost:8080',
        ws: true,
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
