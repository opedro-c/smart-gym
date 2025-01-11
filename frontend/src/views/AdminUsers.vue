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


const users = ref([
    { admin: false, name: 'Joao', email: 'joao@gmail.com', id: 2, rfid: 'A0B8CCDF' },
    { admin: false, name: 'Claudia', email: 'claudia@gmail.com', id: 3, rfid: 'C4A5B6D7' },
    { admin: false, name: 'Ronaldo', email: 'ronaldo@gmail.com', id: 4, rfid: 'D4C3B2A1' },
    { admin: false, name: 'Vicente', email: 'vicente@gmail.com', id: 5, rfid: '43B2A1D' },
    { admin: false, name: 'Rafael', email: 'rafael@gmail.com', id: 6, rfid: 'F1E2D3C' },
])

</script>

<template>
    <div class="flex justify-center">

    <div class="max-w-2xl">
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Name</TableHead>
                    <TableHead>user_id</TableHead>
                    <TableHead>email</TableHead>
                    <TableHead>RFID</TableHead>
                    <TableHead></TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(user, index) in users" :key="index">
                    <TableCell>{{ user.name}}</TableCell>
                    <TableCell>{{  user.id  }}</TableCell>
                    <TableCell>{{ user.email  }}</TableCell>
                    <TableCell>{{ user.rfid }}</TableCell>
                    <TableCell>
                        <ButtonEditUser v-model="users[index]" />
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>
    </div>
    </div>
</template>