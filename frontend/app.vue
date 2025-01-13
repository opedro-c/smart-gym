<script setup lang="ts">
import Toaster from './components/ui/toast/Toaster.vue';
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <NuxtLayout>
      <NuxtPage />
      <Toaster /> 
    </NuxtLayout>
  </div>
</template>
