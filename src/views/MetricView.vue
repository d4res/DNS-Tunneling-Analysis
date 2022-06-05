<template>
  <el-row>
    <el-col :span="12">
      <canvas id="typeChart"></canvas>
    </el-col>

    <el-col :span="12">
      <canvas id="tagChart"></canvas>
    </el-col>
  </el-row>
</template>

<script setup>
import { Chart } from "chart.js";
import getChart from "./chart.js";
const { default: axios } = require("axios");
const { onMounted, ref } = require("vue-demi");

onMounted(() => {
  let data = ref();
  const typeCtx = getChart("typeChart");
  const tagCtx = getChart("tagChart")
  axios.get("http://127.0.0.1:8888/metric").then((response) => {
    data.value = response.data;
    

    new Chart(typeCtx, {
      type: "pie",
      data: {
        labels: ["MX", "TXT", "CNAME"],
        datasets: [
          {
            data: [data.value.MX, data.value.TXT, data.value.CNAME],
            backgroundColor: ["pink", "Aqua", "yellow"],
          },
        ],
      },
      
    });

    new Chart(tagCtx, {
        type: "pie",
        data: {
            labels: ["异常", "正常"],
            datasets: [
                {
                    data: [data.value.Eval, data.value.All - data.value.Eval],
                    backgroundColor: ["HotPink", "Lime"]
                }
            ]
        }
    })
  });
});
</script>