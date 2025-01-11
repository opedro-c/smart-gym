<script setup lang="ts">
import { useColorMode } from '@vueuse/core'
import { RouterView } from 'vue-router'
import Toaster from '@/components/ui/toast/Toaster.vue'
import { onMounted } from 'vue';

const mode = useColorMode()
onMounted(() => {
  mode.value = 'dark'
})

</script>

<template>
  <div class="min-h-screen flex flex-col">
    <RouterView name="Navbar" />
    <RouterView />
    <Toaster />
  </div>
</template>
