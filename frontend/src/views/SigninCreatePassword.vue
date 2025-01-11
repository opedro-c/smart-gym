<script setup lang="ts">
// Zod
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'

// Componentes
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { useToast } from '@/components/ui/toast/use-toast'

// Pinia
import { useAuthStore } from '@/stores/auth'

// Vue imports
const router = useRouter()

// Lifecycle Hooks
import { onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const formSchema = toTypedSchema(
  z.object({
    password: z
      .string({
        required_error: 'Campo senha obrigatório',
      })
      .min(5, { message: 'Senha deve ter no mínimo 5 caracteres.' }),
  }),
)

const form = useForm({
  validationSchema: formSchema,
})

const { toast, dismiss } = useToast()

const { user } = useAuthStore()

const onSubmit = form.handleSubmit(async (values) => {
  toast({
    title: 'Entrando...',
    description: 'Aguarde um momento.',
    duration: 5000,
    variant: 'default',
  })
  console.log(values)
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
          <h1 class="text-2xl font-semibold">Crie uma nova senha</h1>
        </header>
        <form @submit="onSubmit" class="w-full max-w-[450px]">
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
          <Button type="submit" class="w-full">Alterar</Button>
        </form>
      </main>
    </main>
  </div>
</template>
