<template>
  <el-config-provider namespace="ep">
    <el-tabs v-model="editableTabsValue" type="card" class="demo-tabs" closable @tab-remove="removeTab" @tab-click="tabClick">
      <!-- :closable="false" 不可关闭 -->
      <el-tab-pane label="首页" name="main" :closable="false">  
        <el-button type="primary" @click="addTab">添加tab</el-button>
        <el-button type="primary" @click="changeTo2">切换到第二个</el-button>
      </el-tab-pane>
      <el-tab-pane v-for="item in editableTabs" :key="item.name" :label="item.title" :name="item.name">
        <div style="display: flex">
          <div>
            <HelloWorld :msg="item.content" />
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>

  </el-config-provider>
</template>

<script lang="ts" setup>
import { ElMessage, TabsPaneContext } from 'element-plus';
import { ref } from 'vue'

const editableTabsValue = ref('main')
const editableTabs = ref([
  {
    title: 'Tab 1',
    name: '1',
    content: 'Tab 1 content',
  },
  {
    title: 'Tab 2',
    name: '2',
    content: 'Tab 2 content',
  },
])

const tabClick = (pane: TabsPaneContext, ev: Event) => {
  console.log(pane, ev)
  ElMessage.success(`${pane.paneName} clicked!`)
}

let tabIndex = 2

const addTab = (targetName: string) => {
  const newTabName = `${++tabIndex}`
  editableTabs.value.push({
    title: 'New Tab',
    name: newTabName,
    content: 'New Tab content',
  })
}

const changeTo2 = () => {
  editableTabsValue.value = '2'
}

const removeTab = (targetName: string) => {
  const tabs = editableTabs.value
  let activeName = editableTabsValue.value
  if (activeName === targetName) {
    tabs.forEach((tab, index) => {
      if (tab.name === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1]
        if (nextTab) {
          activeName = nextTab.name
        }
      }
    })
  }

  editableTabsValue.value = activeName
  editableTabs.value = tabs.filter((tab) => tab.name !== targetName)
}
</script>

<style>
#app {
  text-align: center;
  color: var(--ep-text-color-primary);
}

.element-plus-logo {
  width: 50%;
}
</style>
