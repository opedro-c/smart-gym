<script setup lang="ts">
import type { MachineData } from '@/types';

const model = defineModel<MachineData>({required: true})

defineEmits(['submit'])

</script>

<template>
  <Dialog>
    <DialogTrigger as-child>
      <Button variant="outline">
        <slot></slot>
      </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Editar maquina</DialogTitle>
        <DialogDescription>
          Click save when you're done.
        </DialogDescription>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="name" class="text-right">
            Name
          </Label>
          <Input id="name" class="col-span-3" v-model="model.name" />
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="username" class="text-right">
            origin_id
          </Label>
          <Input id="origin_id"   class="col-span-3"  v-model="model.origin_id"/>
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