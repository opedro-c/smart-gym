import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'
import { computed, ref } from 'vue'

export interface User {
  id: number
  name: string
  email: string
  admin: boolean
}

export const useAuthStore = defineStore('auth', () => {
  const token = useLocalStorage('token', '')

  const user = ref<User | null>(null)
  const isAuthenticated = computed(() => user.value !== null)

  function loadUser() {
    if (!token.value) {
      return
    }

    // TODO
    // decode jwt token
  }

  function logout() {
    token.value = ''
    user.value = null
  }

  return { token, user, isAuthenticated, loadUser, logout }
})
