<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'
import type { MachineData } from '~/lib/types';
const { toast } = useToast()

let machines = await useApi().getMachines();
const showingMachines = ref(machines);

async function updateMachine(machineIndex: number) {
    try {
        await useApi().updateMachine(machines[machineIndex].id, machines[machineIndex]);
        machines = await useApi().getMachines();
        showingMachines.value = machines;
    } catch (error) {
        toast({
            title: 'Error',
            description: 'Error updating user',
            variant: 'destructive',
        })
    }
}

const newMachine = ref({} as MachineData);
async function createMachine() {
    try {
        await useApi().createMachine(newMachine.value);
        machines = await useApi().getMachines();
        showingMachines.value = machines;
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
            <ButtonEditMachine v-model="newMachine" @submit="createMachine">
                Criar maquina
            </ButtonEditMachine>
        </div>

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
                <TableRow v-for="(machine, index) in showingMachines" :key="index">
                    <TableCell>{{ machine.name}}</TableCell>
                    <TableCell>{{ machine.origin_id }}</TableCell>
                    <TableCell>
                        <ButtonEditMachine v-model="machines[index]" @submit="updateMachine(index)">
                            Editar
                        </ButtonEditMachine>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>
    </div>
    </div>
</template>