<template>

<el-table :data="data" max-height="1000"> 
    <el-table-column prop="payload" label="内容"></el-table-column>
    <el-table-column prop="time" label="时间"> </el-table-column>
    <el-table-column label="标签">
        <template #default="scope">
        <el-tag
          :type="scope.row.tag === true ? 'danger' : 'success'"
          disable-transitions
          >
          {{ scope.row.tag === true? "有害" : "正常" }}
          </el-tag
        >
      </template></el-table-column>
</el-table>

</template>

<script setup>
import axios from 'axios'
const { onMounted, ref }=require("vue-demi")

const data = ref()

onMounted(() => {
    const url = "http://127.0.0.1:8888/data"
    axios.get(url)
    .then((response) => {
        console.log(response.data)
        data.value = response.data
    })
})
</script>
