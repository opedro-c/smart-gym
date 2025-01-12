import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'
import { computed, ref } from 'vue'
import { Userdata } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  const token = useLocalStorage('token', '')

  const user = ref<Userdata | null>(null)
  const isAuthenticated = computed(() => user.value !== null)

  const isAdmin = ref(false);

  function logout() {
    token.value = ''
    user.value = null

    // redirect to Login page
    window.location.href = '/'
  }

  return { token, user, isAuthenticated, logout }
})
