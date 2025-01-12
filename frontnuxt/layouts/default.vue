<script setup lang="ts">
const router = useRouter()

const handleLogout = (e: MouseEvent) => {
    useAuthUser().logout()
    router.push({ name: '/' })
}

const adminOptions = [
  { title: "Users", link: "/admin/users" },
  { title: "Machines", link: "/admin/machines" },
] as const;

const userOptions = [
  { title: "Dashboard", link: "/dashboards" },
] as const;

// const activeOptions = authStore.user?.admin ? adminOptions : userOptions
const activeOptions = adminOptions 


const STYLE_HOVER_MARK_LINE_EFFECT = "relative after:content-[''] after:absolute after:left-0 after:bottom-0 after:h-[1px] after:w-0 after:bg-primary hover:after:w-[95%] after:transition-all after:ease-in-out after:duration-300"
</script>
<template>
  <nav
    class="sticky backdrop-blur-sm z-20 flex items-center justify-center space-x-5 p-5 border-b border-foreground/10"
  >
    <ul class="flex items-center md:gap-4 gap-8 text-foreground/70 ">
      <li 
        :key="option.title" v-for="option in activeOptions"
        :class="{
          [STYLE_HOVER_MARK_LINE_EFFECT]: true,
          'text-foreground after:w-full': option.link === router.currentRoute.value.path
        }"
      >
        <NuxtLink :to="option.link">
          {{ option.title }}
        </NuxtLink>
      </li>
    </ul>

    <div class="!absolute right-5 flex gap-5 items-center">
      <ToggleDarkMode />
      <p @click="handleLogout" :class="{
        [STYLE_HOVER_MARK_LINE_EFFECT]: true,
        'hover:cursor-pointer': true
      }">Logout</p>
      </div>
  </nav>
  <slot></slot>
</template>