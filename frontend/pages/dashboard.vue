<script setup lang="ts">
import VueApexCharts from "vue3-apexcharts";

const todayBeginTime = 0
const nowTime = new Date().getTime()

const machines = await useApi().getMachines();
const { user } = useAuthUser();

const workoutData =  await Promise.all(
    machines.map(machine => 
        useApi()
        .getWorkoutData(user.value?.id ?? 1, machine.origin_id, todayBeginTime, nowTime)
        .then(data => ({
            origin_id: machine.origin_id,
            name: machine.name,
            data
        })
    )
));


const formatTime = (date: Date) => {
  const hours = date.getHours().toString().padStart(2, '0');
  const minutes = date.getMinutes().toString().padStart(2, '0');
  const seconds = date.getSeconds().toString().padStart(2, '0');
  return `${hours}:${minutes}:${seconds}`;
};
</script>

<template>
    <div class="flex flex-col gap-7 justify-center items-center w-full">

      <div v-for="workout in workoutData"  :key="workout.origin_id" class="bg-primary w-1/2" >
        <template v-if="workout.data !== null">
          <VueApexCharts
          type="bar"
          :options="{
            xaxis: {
              categories: workout.data.map(data => formatTime(new Date(data.started_at))),
            },
          }"
          :series="[
            {
              name: 'weight',
              data: workout.data.map(data => data.data.weight),
            },
          ]"
        />
        </template>
      </div>
    </div>
</template>