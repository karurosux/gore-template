import {defineConfig} from "vite"
import react from "@vitejs/plugin-react-swc"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  // base: "/static/",
  build: {
    rollupOptions: {
      output: {
        dir: "../dist/static",
        entryFileNames: "assets/[name]-[hash].js",
        assetFileNames: "assets/[name]-[hash].[extname]",
        chunkFileNames: "assets/[name]-[hash].js",
      },
    },
  },
  resolve: {
    alias: {
      "@src": "/src",
      "@components": "/src/components",
      "@containers": "/src/containers",
      "@constants": "/src/constants",
      "@hooks": "/src/hooks",
      "@slices": "/src/slices",
      "@api": "/src/api",
      "@selectors": "/src/selectors",
      "@assets": "/src/assets",
    },
  },
  server: {
    port: 8080,
    proxy: {
      "/api": {
        target: "http://localhost:8081",
        changeOrigin: true,
        rewrite: (path) => {
          return "http://localhost:8081/" + path
        },
      },
    },
  },
})
