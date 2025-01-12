<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'

definePageMeta({
  layout: false,
})
const { user } = useAuthUser()

const { toast, dismiss } = useToast()

async function onSubmit(values: { email: string, password: string }) {
  toast({
    title: 'Entrando...',
    description: 'Aguarde um momento.',
    duration: 5000,
    variant: 'default',
  })

  try {
    const response = await useApi().login(values);
    user.value = response
    console.log('user', response)

    if (response.admin) {
      navigateTo('/admin/machines')
    } else {
      navigateTo('/machines')
    }

  } catch (error) {
    console.error(error)
    toast({
      title: 'Erro em login...',
      duration: 5000,
      variant: 'destructive',
    })
  }
}

onUnmounted(() => {
  dismiss()
})
</script>

<template>
 <FormLogin @submit="onSubmit">
   <h1 class="text-2xl font-semibold">Login</h1>
   <small class="text-sm text-muted-foreground">
     Entre no sistema para poder desfrutar das funcionalidades.
    </small>
  </FormLogin>
</template>
