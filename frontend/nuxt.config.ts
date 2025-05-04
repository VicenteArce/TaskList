// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  // compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  typescript: {
    strict: true, 
  },
  runtimeConfig: {  
    public: {
      API_BASE_URL: process.env.TODO_LIST_SERVER || 'http://localhost:8080',
    },
  },
  app: {
    head: {
      title: 'TodoList App', 
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Una aplicación simple para gestionar tareas creada con Nuxt 3 y Go.' },
        { name: 'keywords', content: 'nuxt, go, todo, app, tareas' },
        { name: 'author', content: 'Vicente Arce' },
        ]
    },
    
  },
  // Configuración del servidor de desarrollo Nuxt
  devServer: {
    port: 3000, 
  },
  
  modules: [
    '@nuxt/content',
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils',
    '@nuxt/ui'
  ],
  css: ['~/assets/css/main.css'],
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
})