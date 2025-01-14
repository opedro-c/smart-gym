<script setup lang="ts">
const machines = await useApi().getMachines();
</script>

<template>
    <div class="flex flex-col items-center">
        <div class="max-w-2xl pt-5 flex gap-5 flex-wrap">
            <Card 
                :key="machine.origin_id"
                v-for="machine in machines" 
                :class="machine.origin_id.length > 3 ? 'border-destructive' : 'border-green-700' "
            >
                <CardHeader>
                    <CardTitle>
                        Notifications
                    </CardTitle>
                    <CardDescription class="flex gap-3 items-center">
                        <template v-if="machine.origin_id.length > 3">
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