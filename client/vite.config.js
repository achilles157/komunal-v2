import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  // Tambahkan bagian 'server' ini
  server: {
    proxy: {
      // Setiap permintaan yang dimulai dengan /api
      // akan diteruskan ke server backend Anda
      '/api': {
        target: 'http://localhost:8080', // Sesuaikan port jika backend Anda berbeda
        changeOrigin: true,
      },
    },
  },
})