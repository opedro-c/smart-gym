<script setup lang="ts">
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'

import { ref } from 'vue'

const machines = ref([
    {  name: "maquina1", origin_id: 'A0B8CCDF', occupied: true },
    {  name: "maquina2", origin_id: 'C4A5B6D7', occupied: false },
    {  name: "maquina3", origin_id: 'D4C3B2A1', occupied: true },
    {  name: "maquina4", origin_id: '43B2A1D', occupied: false },
    {  name: "maquina5", origin_id: 'F1E2D3C', occupied: true },
    {  name: "maquina6", origin_id: 'F1E2D3C', occupied: true },
    {  name: "maquina7", origin_id: 'F1E2D3C', occupied: false },
    {  name: "maquina8", origin_id: 'F1E2D3C', occupied: false },
    {  name: "maquina9", origin_id: 'F1E2D3C', occupied: true },
])
</script>

<template>
    <div class="flex flex-col items-center">
        <div class="max-w-2xl pt-5 flex gap-5 flex-wrap">
            <Card 
                :key="machine.origin_id"
                v-for="machine in machines" 
                :class="machine.occupied ? 'border-destructive' : 'border-green-700' "
            >
                <CardHeader>
                    <CardTitle>
                        Notifications
                    </CardTitle>
                    <CardDescription class="flex gap-3 items-center">
                        <template v-if="machine.occupied">
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