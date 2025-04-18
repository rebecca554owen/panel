<script setup lang="ts">
defineOptions({
  name: 'apps-gitea-index'
})

import Editor from '@guolao/vue-monaco-editor'
import { NButton, NPopconfirm } from 'naive-ui'

import gitea from '@/api/apps/gitea'
import systemctl from '@/api/panel/systemctl'

const currentTab = ref('status')
const status = ref(false)
const isEnabled = ref(false)
const config = ref('')

const statusStr = computed(() => {
  return status.value ? '正常运行中' : '已停止运行'
})

const getStatus = async () => {
  status.value = await systemctl.status('gitea')
}

const getIsEnabled = async () => {
  isEnabled.value = await systemctl.isEnabled('gitea')
}

const getConfig = async () => {
  config.value = await gitea.config()
}

const handleSaveConfig = () => {
  useRequest(gitea.saveConfig(config.value)).onSuccess(() => {
    window.$message.success('保存成功')
  })
}

const handleStart = async () => {
  await systemctl.start('gitea')
  window.$message.success('启动成功')
  await getStatus()
}

const handleStop = async () => {
  await systemctl.stop('gitea')
  window.$message.success('停止成功')
  await getStatus()
}

const handleRestart = async () => {
  await systemctl.restart('gitea')
  window.$message.success('重启成功')
  await getStatus()
}

const handleIsEnabled = async () => {
  if (isEnabled.value) {
    await systemctl.enable('gitea')
    window.$message.success('开启自启动成功')
  } else {
    await systemctl.disable('gitea')
    window.$message.success('禁用自启动成功')
  }
  await getIsEnabled()
}

onMounted(() => {
  getStatus()
  getIsEnabled()
  getConfig()
})
</script>

<template>
  <common-page show-footer>
    <template #action>
      <n-button
        v-if="currentTab == 'config'"
        class="ml-16"
        type="primary"
        @click="handleSaveConfig"
      >
        <TheIcon :size="18" icon="material-symbols:save-outline" />
        保存
      </n-button>
    </template>
    <n-tabs v-model:value="currentTab" type="line" animated>
      <n-tab-pane name="status" tab="运行状态">
        <n-card title="运行状态">
          <template #header-extra>
            <n-switch v-model:value="isEnabled" @update:value="handleIsEnabled">
              <template #checked> 自启动开 </template>
              <template #unchecked> 自启动关 </template>
            </n-switch>
          </template>
          <n-space vertical>
            <n-alert :type="status ? 'success' : 'error'">
              {{ statusStr }}
            </n-alert>
            <n-space>
              <n-button type="success" @click="handleStart">
                <TheIcon :size="24" icon="material-symbols:play-arrow-outline-rounded" />
                启动
              </n-button>
              <n-popconfirm @positive-click="handleStop">
                <template #trigger>
                  <n-button type="error">
                    <TheIcon :size="24" icon="material-symbols:stop-outline-rounded" />
                    停止
                  </n-button>
                </template>
                确定要停止 Gitea 吗？
              </n-popconfirm>
              <n-button type="warning" @click="handleRestart">
                <TheIcon :size="18" icon="material-symbols:replay-rounded" />
                重启
              </n-button>
            </n-space>
          </n-space>
        </n-card>
      </n-tab-pane>
      <n-tab-pane name="config" tab="修改配置">
        <n-space vertical>
          <n-alert type="warning">
            此处修改的是 Gitea 配置文件，如果您不了解各参数的含义，请不要随意修改！
          </n-alert>
          <Editor
            v-model:value="config"
            language="ini"
            theme="vs-dark"
            height="60vh"
            mt-8
            :options="{
              automaticLayout: true,
              formatOnType: true,
              formatOnPaste: true
            }"
          />
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="run-log" tab="运行日志">
        <realtime-log service="gitea" />
      </n-tab-pane>
    </n-tabs>
  </common-page>
</template>
