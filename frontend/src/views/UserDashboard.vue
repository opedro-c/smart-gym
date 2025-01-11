<script setup lang="ts">
import Card from '@/components/ui/card/Card.vue';

const series = [
    {
        name: "series1",
        data: [31, 40, 28, 51, 42, 109, 100]
    }
]

const chartOptions = {
    chart: {
        height: 350,
        type: 'line',
        zoom: {
            enabled: false
        }
    },
    dataLabels: {
        enabled: false
    },
    stroke: {
        curve: 'straight'
    },
    title: {
        text: 'Product Trends by Month',
        align: 'left'
    },
    grid: {
        row: {
            colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
            opacity: 0.5
        },
    },
    xaxis: {
        categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul'],
    }
}
</script>

<template>
      <div class="flex flex-col items-center">
        <div class="max-w-2xl pt-5 flex gap-5 flex-wrap">
            <Card class="p-5 bg-white">
                <apexchart type="line" height="350" :options="chartOptions" :series="series"></apexchart>
            </Card>
        </div>
      </div>
</template>