import { computed } from 'vue'
import type { UserData } from '~/lib/types'

export const useAuthUser = () => {
  const user = useState('user', () => null as UserData | null)
 
  const isAuthenticated = computed(() => user.value !== null)

  function logout() {
    user.value = null

    // redirect to Login page
    window.location.href = '/'
  }

  return { user, isAuthenticated, logout }
}
