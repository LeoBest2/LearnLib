<template>
  <el-config-provider namespace="ep">
    <el-menu class="el-menu-demo" mode="horizontal">
      <el-menu-item index="1">Vue JWT Demo</el-menu-item>
      <el-menu-item index="2">当前用户: {{ username }}</el-menu-item>
    </el-menu>
    <!-- <img alt="Vue logo" class="element-plus-logo" src="./assets/logo.png" /> -->
    <el-row style="margin-top: 2rem">
      <el-col :span="8"></el-col>
      <el-col :span="8">
        <template v-if="username === ''">
          <el-form :model="form">
            <el-form-item label="用户">
              <el-input v-model="form.username" placeholder="username" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password" placeholder="password" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" style="width: 100%" @click="onLogin">获取Token</el-button>
            </el-form-item>
          </el-form>
        </template>
        <template v-else>
          <template v-if="result !== ''">
            <h1>请求成功</h1>
            <h2>{{ result }}</h2>
          </template>
          <el-button type="primary" @click="onAPI">请求API</el-button>
          <el-button @click="onLogout">退出登陆</el-button>
        </template>
      </el-col>
      <el-col :span="8"></el-col>
    </el-row>
    <!-- <HelloWorld msg="Hello Vue 3.0 + Element Plus + Vite" /> -->
  </el-config-provider>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus';

const form = reactive({
  username: '',
  password: '',
})

const username = ref(localStorage.getItem('username') || '')
const result = ref('')

const onLogin = () => {
  axios.post('/auth/login', { username: form.username, password: form.password })
    .then(response => {
      if (response.data.code !== 0) {
        ElMessage.warning(`登陆失败: ${response.data.msg}`)
      } else {
        localStorage.setItem('token', response.data.data);
        localStorage.setItem('username', form.username);
        username.value = form.username;
        ElMessage.info(`${form.username}登陆成功!`)
      }
    })
    .catch(err => {
      console.log(err)
      ElMessage.error(err)
    })
}

const onAPI = () => {
  axios.get('/api/test', { headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') } })
    .then(response => {
      if (response.data.code !== 0) {
        ElMessage.warning(`获取: ${response.data.msg}`)
      } else {
        result.value = response.data.data
        ElMessage.info(`请求成功！`)
      }
    })
    .catch(err => {
      console.log(err)
      ElMessage.error(err)
    })
}

const onLogout = () => {
  localStorage.removeItem('username')
  localStorage.removeItem('token')
  username.value = ''
}
</script>

<style>
#app {
  text-align: center;
  color: var(--ep-text-color-primary);
}

/* 
.element-plus-logo {
  width: 50%;
} */
</style>
