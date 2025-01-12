<script setup lang="ts">
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'

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

const emit = defineEmits<{
    (event: 'submit', values: { email: string; password: string }): void
}>();

const onSubmit = form.handleSubmit(async (values) => emit('submit', values))
</script>

<template>
  <div class="flex flex-1 flex-col md:flex-row">
    <main class="relative flex flex-col justify-center items-center flex-1 p-5">
      <main>
        <header class="text-center mb-5">
          <slot></slot>
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
