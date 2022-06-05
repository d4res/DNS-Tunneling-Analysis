<template>

<el-table :data="infoData" max-height="1000">
    <el-table-column label="包序列号">
        <template #default="scope">
            {{scope.row.msgheader.packetid}}
        </template>
    </el-table-column>

    <el-table-column label="类型">
         <template #default="scope">
            {{ table[scope.row.msgheader.msgtype]}}
        </template>
    </el-table-column>

    <el-table-column label="会话id">
       <template #default="scope">
          {{ scope.row.msgheader.sessid }}
        </template>
    </el-table-column>

    <el-table-column label="操作">
         <template #default="scope">
           <span v-if="scope.row.msgheader.msgtype==0"> 
               <el-tag type="success">
                   SYN
               </el-tag>
               建立连接:  {{scope.row.payload.name}} 
            </span>

            <span v-if="scope.row.msgheader.msgtype==1">
                <el-tag type="success"
                >CMD</el-tag>
                执行命令: {{scope.row.payload.cmd}}
            </span> 
        </template>
    </el-table-column>
</el-table>


<el-collapse >
    <el-collapse-item title="原始数据" >
        <el-descriptions column="1">
    <el-descriptions-item v-for="item in infoData" :key="item._id" >
       {{item}}
    </el-descriptions-item>
</el-descriptions>
    </el-collapse-item>
</el-collapse>
</template>


<script setup> 
const { default: axios }=require("axios");
const { ref, onMounted }=require("vue-demi");

let infoData = ref()

const table = {
    0: 	"SYN",
	1: "MSG",
	2: "FIN",
}

onMounted(()=>{
    axios.get("http://127.0.0.1:8888/info").then(
(response)=>{
    infoData.value = response.data
    console.log(infoData.value
    )
})
})
</script>