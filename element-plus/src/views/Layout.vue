<template>
    <el-container style="height: 100%;">
        <el-header>
            <el-menu :default-active="activeIndex" mode="horizontal" router>
                <el-menu-item>
                    <img src="../assets/element-plus-logo.svg" style="width: 120px; height: 60px;" />
                </el-menu-item>
                <template v-if="username !== ''">
                    <template v-for="route in routes" :key="route.name" :route="route">
                        <el-menu-item :index="route.path">{{ route.meta.displayName }}</el-menu-item>
                    </template>
                    <el-sub-menu style="margin-left: auto;">
                        <template #title>
                            <i class="el-icon-setting"></i>
                            {{ username }}
                        </template>
                        <el-menu-item @click="logoff">
                            <i class="el-icon-right"></i>
                            注销登陆
                        </el-menu-item>
                    </el-sub-menu>
                </template>
                <template v-else>
                    <el-menu-item class="is-active">请登陆后访问!</el-menu-item>
                </template>
            </el-menu>
        </el-header>
        <el-main>
            <slot name="main"></slot>
        </el-main>
        <el-footer height="40px" style="color: #909399; text-align: center;">©2021 Leo xxxxxx.com</el-footer>
    </el-container>
</template>

<script>

import { useRoute, useRouter } from "vue-router";
import { ElMessage } from 'element-plus'

export default {
    name: 'Layout',
    setup() {
        const router = useRouter();
        const route = useRoute();
        return {
            routes: router.getRoutes().filter(value => value.name != 'Login' && value.name != 'Logout'),
            activeIndex: route.path,
        }
    },
    computed: {
        username() {
            return this.$store.state.username || ''
        }
    },
    methods: {
        logoff() {
            ElMessage({
                showClose: true,
                message: `用户 ${this.$store.state.username} 注销成功!`,
                type: 'success',
            })
            this.$store.commit('logoff')
            this.$router.replace({ name: 'Login' })
        }
    }
}
</script>