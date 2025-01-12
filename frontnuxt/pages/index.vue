<script setup lang="ts">
import { onUnmounted } from 'vue'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'
import { useToast } from '@/components/ui/toast/use-toast'

definePageMeta({
  layout: false,
})

const formSchema = toTypedSchema(
  z.object({
    email: z
      .string({ required_error: 'Campo email obrigatório' })
      .min(1, { message: 'Esse campo deve ser preenchido.' }),
    password: z
      .string({ required_error: 'Campo senha obrigatório' })
      .min(3, { message: 'Senha deve ter no mínimo 5 caracteres.' }),
  }),
)

const form = useForm({
  validationSchema: formSchema,
})

const { toast, dismiss } = useToast()

const onSubmit = form.handleSubmit(async ({ email, password }) => {
  toast({
    title: 'Entrando...',
    description: 'Aguarde um momento.',
    duration: 5000,
    variant: 'default',
  })

  try {
    const response = await useApi().login({ email, password });
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
})

onUnmounted(() => {
  dismiss()
})
</script>

<template>
  <div class="flex flex-1 flex-col md:flex-row">
    <main class="relative flex flex-col justify-center items-center flex-1 p-5">
      <main>
        <header class="text-center mb-5">
          <h1 class="text-2xl font-semibold">Login</h1>
          <small class="text-sm text-muted-foreground">
            Entre no sistema para poder desfrutar das funcionalidades.
          </small>
        </header>
        <form @submit="onSubmit" class="w-full max-w-[450px]">
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormControl>
                <Input
                  placeholder="Email..."
                  v-bind="componentField"
                  autocomplete="email"
                />
              </FormControl>
              <FormDescription />
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Senha..."
                  v-bind="componentField"
                  autocomplete="current-password"
                />
              </FormControl>
              <FormDescription />
              <FormMessage />
            </FormItem>
          </FormField>
          <Button type="submit" class="w-full">Entrar</Button>
        </form>
      </main>
    </main>
  </div>
</template>
