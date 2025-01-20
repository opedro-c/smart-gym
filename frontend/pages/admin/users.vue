<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'
import type { UserData } from '~/lib/types';
import type { GetUserData } from '~/repositories/user';
const { toast } = useToast()

let users = await useApi().getAllUsers();
const showingUsers = ref(users);

async function updateUser(userIndex: number) {
    try {
        await useApi().updateUser(users[userIndex].id, users[userIndex].data);
        await useApi().updateUserRfids(users[userIndex].id, users[userIndex].rfid);

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

const newUser = ref({ admin: false, rfid: '', data: { email: '', password: '', username: '' } } as GetUserData);
async function createUser() {
    try {
        await useApi().createUser({ ...newUser.value.data, ...newUser.value });
        users = await useApi().getAllUsers();
        showingUsers.value = users;
    } catch (error) {
        toast({
            title: 'Error',
            description: 'Error creating user',
            variant: 'destructive',
        })
    }
}
</script>

<template>
    <div class="flex flex-col items-center justify-center">

        <div class="mt-5">
            <ButtonEditUser v-model="newUser" @submit="createUser">
                Criar user
            </ButtonEditUser>
        </div>

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
                        <ButtonEditUser v-model="users[index]" @submit="updateUser(index)">
                            Editar
                        </ButtonEditUser>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>
    </div>
    </div>
</template>