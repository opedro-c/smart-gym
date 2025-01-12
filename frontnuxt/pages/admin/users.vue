<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'
const { toast } = useToast()

let users = await useApi().getAllUsers();
const showingUsers = ref(users);

async function updateUser(userIndex: number) {
    try {
        await useApi().updateUser(users[userIndex].id, users[userIndex].data);
        users = await useApi().getAllUsers();
        showingUsers.value = users;
    } catch (error) {
        toast({
            title: 'Error',
            description: 'Error updating user',
            variant: 'destructive',
        })
    }
}
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
                <TableRow v-for="(user, index) in showingUsers.sort((a,b) => a.id - b.id)" :key="index">
                    <TableCell>{{ user.data.username }}</TableCell>
                    <TableCell>{{  user.id  }}</TableCell>
                    <TableCell>{{ user.data.email  }}</TableCell>
                    <TableCell>{{ user.rfid }}</TableCell>
                    <TableCell>
                        <ButtonEditUser v-model="users[index]" @submit="updateUser(index)" />
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>
    </div>
    </div>
</template>