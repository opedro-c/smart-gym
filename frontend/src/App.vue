<script setup lang="ts">
import { RouterView } from 'vue-router'
import Toaster from '@/components/ui/toast/Toaster.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <RouterView name="Navbar" />
    <RouterView />
    <Toaster />
  </div>
</template>
