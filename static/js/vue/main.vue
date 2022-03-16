<template>
  <el-container>
    <Toast></Toast>
    <router-view></router-view>
    <el-footer>
      <el-row>
        <el-col>
          <el-header>Version</el-header>
          <el-row>Software build {{ Buildnr }}</el-row>
        </el-col>
        <el-col>
          <el-header>Info</el-header>
          <el-row>
            <i class="copyright icon"></i> {{ new Date().getFullYear() }} by
            Invido.it
          </el-row>
        </el-col>
      </el-row>
    </el-footer>
  </el-container>
</template>