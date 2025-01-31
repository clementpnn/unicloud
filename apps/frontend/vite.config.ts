import path from "node:path"
import { defineConfig } from "vite"
import react from "@vitejs/plugin-react-swc"
import svgr from "vite-plugin-svgr"
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), svgr()],
  resolve: {
    alias: {
      // eslint-disable-next-line unicorn/prefer-module
      "@": path.resolve(__dirname, "./src")
    }
  }
})