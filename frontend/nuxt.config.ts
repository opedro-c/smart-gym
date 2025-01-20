// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    'shadcn-nuxt',
    '@nuxtjs/color-mode',
    '@vueuse/nuxt',
    'nuxt-echarts'
  ],
  echarts: {
    ssr: false,
    renderer: ['canvas', 'svg'],
    charts: ['BarChart', 'LineChart', 'LinesChart'],
    components: ['DatasetComponent', 'GridComponent', 'TooltipComponent'],
  },
})