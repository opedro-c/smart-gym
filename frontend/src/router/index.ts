import { createRouter, createWebHistory } from 'vue-router'
import Signin from '@/views/Signin.vue'
import SigninCreatePassword from '@/views/SigninCreatePassword.vue'
import AdminMachines from '@/views/AdminMachines.vue'
import UserDashboard from '@/views/UserDashboard.vue'
import Navbar from '@/views/Navbar.vue'
import AdminUsers from '@/views/AdminUsers.vue'
import UserAvailableMachines from '@/views/UserAvailableMachines.vue'
import AdminSignin from '@/views/AdminSignin.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Signin,
    },

    {
      path: '/admin',
      name: 'admin-login',
      component: AdminSignin,
    },


    {
      path: '/admin/users',
      name: 'admin-users-managing',
      components: {
        default: AdminUsers,
        Navbar: Navbar,
      },
    },

    {
      path: '/admin/machines',
      name: 'admin-machines',
      components: {
        default: AdminMachines,
        Navbar: Navbar,
      },
    },

    {
      path: '/machines',
      name: 'user-view-machines',
      components: {
        default: UserAvailableMachines,
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
