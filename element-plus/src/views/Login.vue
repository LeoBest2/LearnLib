<template>
    <Layout>
        <template v-slot:main>
            <div style="display: flex; justify-content: center; align-items: center; height: 100%;">
                <el-form
                    ref="loginForm"
                    :model="loginForm"
                    status-icon
                    :rules="rules"
                    label-width="120px"
                    style="width: 400px;"
                >
                    <el-form-item label="用户" prop="username" required>
                        <el-input v-model="loginForm.username" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="userpwd" required>
                        <el-input
                            v-model="loginForm.userpwd"
                            type="password"
                            autocomplete="off"
                            show-password
                        ></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button
                            type="primary"
                            @click="submitForm('loginForm')"
                            style="width: 100%;"
                        >登陆</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </template>
    </Layout>
</template>
<script>

import { ElMessage } from 'element-plus'
import Layout from '../views/Layout.vue';


export default {
    data() {
        return {
            loginForm: {
                username: '',
                userpwd: ''
            },
            rules: {
                username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
                userpwd: [{ required: true, message: '请输入密码', trigger: 'blur' }]
            },
        }
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.$store.commit('login', this.loginForm.username)
                    ElMessage({
                        showClose: true,
                        message: `用户 ${this.loginForm.username} 登陆成功!`,
                        type: 'success',
                    })
                    this.$router.replace({ name: 'Home' })
                } else {
                    return false
                }
            })
        }
    },
    components: {
        Layout
    }
}
</script>
