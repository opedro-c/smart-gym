<script setup lang="ts">
import type { GetUserData } from '~/repositories/api';
import type { UserData } from '~/lib/types';

const model = defineModel<GetUserData>({required: true})

defineEmits(['submit'])

</script>

<template>
  <Dialog>
    <DialogTrigger as-child>
      <Button variant="outline">
        Editar
      </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Editar usuário</DialogTitle>
        <DialogDescription>
          Make changes to your profile here. Click save when you're done.
        </DialogDescription>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="name" class="text-right">
            Name
          </Label>
          <Input id="name" class="col-span-3" v-model="model.data.username" />
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="username" class="text-right">
            Email
          </Label>
          <Input id="username"   class="col-span-3"  v-model="model.data.email"/>
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="password" class="text-right">
            Password
          </Label>
          <Input id="password"   class="col-span-3"  v-model="model.data.password"/>
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="username" class="text-right">
            RFID
          </Label>
          <Input id="username"   class="col-span-3"  v-model="model.rfid"/>
        </div>
      </div>
      <DialogFooter>
        <Button type="submit" @click="$emit('submit')">
          Save changes
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>