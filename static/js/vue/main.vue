<template>
  <el-container>
    <Toast></Toast>
    <router-view></router-view>
    <el-footer>
      <el-row>
        <el-col>
          <h4 class="ui header">Version</h4>
          <span>Software build {{ Buildnr }}</span>
        </el-col>
        <el-col>
          <h4 class="ui header">Info</h4>
          <span>
            <i class="copyright icon"></i> {{ new Date().getFullYear() }} by
            Invido.it
          </span>
        </el-col>
      </el-row>
    </el-footer>
  </el-container>
</template>