<script setup lang="ts">
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import ButtonEditUser from '@/components/ButtonEditUser.vue';
import { ref } from 'vue';
import ButtonEditMachine from '@/components/ButtonEditMachine.vue';


const machines = ref([
    { name: "maquina1", origin_id: 'A0B8CCDF' },
    { name: "maquina2", origin_id: 'C4A5B6D7' },
    { name: "maquina3", origin_id: 'D4C3B2A1' },
    { name: "maquina4", origin_id: '43B2A1D' },
    { name: "maquina5", origin_id: 'F1E2D3C' },
])

</script>

<template>
    <div class="flex justify-center">

    <div class="max-w-2xl">
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Name</TableHead>
                    <TableHead>origin_id</TableHead>
                    <TableHead></TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(machine, index) in machines" :key="index">
                    <TableCell>{{ machine.name}}</TableCell>
                    <TableCell>{{ machine.origin_id }}</TableCell>
                    <TableCell>
                        <ButtonEditMachine v-model="machines[index]" />
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>
    </div>
    </div>
</template>