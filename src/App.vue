<template>
  <el-container>
    <el-header class="header" >
      DNS流量检测
    </el-header>

    <el-container>
    <el-aside width="100px">
      <el-menu
      router="true"
      >
      <el-menu-item index="dashboard">导航</el-menu-item>

      <el-menu-item index="data">原始数据</el-menu-item>

      <el-menu-item index="about">关于</el-menu-item>

      <el-menu-item index="test">测试页面</el-menu-item>
      
      </el-menu>
    </el-aside>

    <el-main>
      <router-view></router-view>
    </el-main> 
    </el-container>
  </el-container>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* width: 100%; */
  /* margin: 0; */
}

body {
  margin: 0;
}

.header {
  background-color:    #409EFF
;
  text-align: left;
  color: white;
  font-size: 2em;
  border-bottom: 1px solid var(--el-border-color);
  box-shadow: var(--el-box-shadow);
}
</style>
