<script setup lang="ts">
import { STATUS_WS_HOST } from '~/lib/config';
import type { MachineStatus } from '~/lib/types';

const machines = await useApi().getMachines();

const machineNames = machines.reduce((acc, machine) => ({ ...acc, [machine.origin_id]: machine.name }), {} as Record<string, string>);

const machineStatus = ref([] as MachineStatus[]);
onMounted(async () => {
    machineStatus.value = await useApi().getActiveMachines();
})

const { data: dataWs } = useWebSocket(`${STATUS_WS_HOST}/ws`, {
    autoReconnect: {
        retries: 3,
        delay: 1000,
        onFailed() {
            alert('Failed to connect WebSocket after 3 retries')
        },
    },
})
watch(dataWs, (message) => {
    if (!message) {
        return
    }

    const newStatus = JSON.parse(message) as MachineStatus;

    const indexMachineStatus = machineStatus.value.findIndex((machine) => machine.origin_id === newStatus?.origin_id);
    if (indexMachineStatus >= 0) {
        machineStatus.value[indexMachineStatus] = newStatus;
    } else {
        machineStatus.value.push(newStatus);
    }
})
</script>

<template>
    <div class="flex flex-col items-center">
        <div class="max-w-2xl pt-5 flex gap-5 flex-wrap">

            <div v-if="machineStatus.length === 0">
                There is no connected machines
            </div>

            <Card 
                :key="machine.origin_id"
                v-for="machine in machineStatus" 
                :class="machine.status ? 'border-destructive' : 'border-green-700' "
            >
                <CardHeader>
                    <CardTitle>
                        {{ machineNames[machine.origin_id] }}
                    </CardTitle>
                    <CardDescription class="flex gap-3 items-center">
                        <template v-if="machine.status">
                            <span class="flex h-3 w-3 rounded-full bg-destructive" />
                            <p class="text-destructive">
                                Ocupada
                            </p>
                        </template>
                        <template v-else>
                            <span class="flex h-3 w-3 rounded-full bg-green-700" />
                            <p class="text-green-700">
                                Livre
                            </p>
                        </template>
                    </CardDescription>
                </CardHeader>
            </Card>
        </div>
    </div>
</template>