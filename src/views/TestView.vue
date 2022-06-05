<template>
  <el-row>
    <el-col :span="8"> <canvas id="myChart"></canvas></el-col>
    <canvas id="to"></canvas>
  </el-row>
</template>

<script setup>
import { onMounted } from "vue-demi";
import getChart from "./chart.js";
import Chart from "chart.js/auto";
import axios from "axios";

onMounted(() => {
  const ctx1 = getChart("myChart");
  //console.log(ctx);

  new Chart(ctx1, {
    type: "bar",
    data: {
      labels: ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
      datasets: [
        {
          label: "# of Votes",
          data: [12, 19, 3, 5, 2, 3],
          backgroundColor: [
            "rgba(255, 99, 132, 0.2)",
            "rgba(54, 162, 235, 0.2)",
            "rgba(255, 206, 86, 0.2)",
            "rgba(75, 192, 192, 0.2)",
            "rgba(153, 102, 255, 0.2)",
            "rgba(255, 159, 64, 0.2)",
          ],
          borderColor: [
            "rgba(255, 99, 132, 1)",
            "rgba(54, 162, 235, 1)",
            "rgba(255, 206, 86, 1)",
            "rgba(75, 192, 192, 1)",
            "rgba(153, 102, 255, 1)",
            "rgba(255, 159, 64, 1)",
          ],
          borderWidth: 1,
        },
      ],
    },
    options: {
      scales: {
        y: {
          beginAtZero: true,
        },
      },
    },
  });

  const ctx2 = getChart("chart2");
  new Chart(ctx2, {
    type: "pie",
    data: {
      datasets: [
        {
          label: "My First Dataset",
          data: [1, 2, 3],
          backgroundColor: ["red", "rgb(54, 162, 235)", "rgb(255, 205, 86)"],
        },
      ],
    },
  });

  console.log("test");
  axios.get("http://127.0.0.1:8888/metric").then((response) => {
    console.log("tet");
    console.log(response);
  });
});
</script>