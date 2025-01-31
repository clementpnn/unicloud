/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_TOKEN: string
  readonly VITE_GROUP: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}