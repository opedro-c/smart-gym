import { createRouter, createWebHistory } from 'vue-router'
import Signin from '@/views/Signin.vue'
import SigninCreatePassword from '@/views/SigninCreatePassword.vue'
import AdminRfid from '@/views/AdminRfid.vue'
import AdminMachines from '@/views/AdminMachines.vue'
import UserDashboard from '@/views/UserDashboard.vue'
import Navbar from '@/views/Navbar.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Signin,
    },
    {
      path: '/new-password',
      name: 'new-password',
      component: SigninCreatePassword,
    },


    {
      path: '/admin/users',
      name: 'users',
      components: {
        default: AdminRfid,
        Navbar: Navbar,
      },
    },

    {
      path: '/admin/rfids',
      name: 'rfids',
      components: {
        default: AdminRfid,
        Navbar: Navbar,
      },
    },

    {
      path: '/admin/machines',
      name: 'machines',
      components: {
        default: AdminMachines,
        Navbar: Navbar,
      },
    },

    {
      path: '/dashboards',
      name: 'dashboards',
      components: {
        default: UserDashboard,
        Navbar: Navbar,
      },
    },
  ],
})

export default router
