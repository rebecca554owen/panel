<script setup lang="ts">
defineOptions({
  name: 'apps-mysql-index'
})

import Editor from '@guolao/vue-monaco-editor'
import { NButton, NDataTable, NInput } from 'naive-ui'
import { useGettext } from 'vue3-gettext'

import mysql from '@/api/apps/mysql'
import ServiceStatus from '@/components/common/ServiceStatus.vue'

const { $gettext } = useGettext()
const currentTab = ref('status')

const { data: rootPassword } = useRequest(mysql.rootPassword, {
  initialData: ''
})
const { data: config } = useRequest(mysql.config, {
  initialData: ''
})
const { data: slowLog } = useRequest(mysql.slowLog, {
  initialData: ''
})
const { data: load } = useRequest(mysql.load, {
  initialData: []
})

const loadColumns: any = [
  {
    title: $gettext('Property'),
    key: 'name',
    minWidth: 200,
    resizable: true,
    ellipsis: { tooltip: true }
  },
  {
    title: $gettext('Current Value'),
    key: 'value',
    minWidth: 200,
    ellipsis: { tooltip: true }
  }
]

const handleSaveConfig = () => {
  useRequest(mysql.saveConfig(config.value)).onSuccess(() => {
    window.$message.success($gettext('Saved successfully'))
  })
}

const handleClearErrorLog = () => {
  useRequest(mysql.clearErrorLog()).onSuccess(() => {
    window.$message.success($gettext('Cleared successfully'))
  })
}

const handleClearSlowLog = () => {
  useRequest(mysql.clearSlowLog()).onSuccess(() => {
    window.$message.success($gettext('Cleared successfully'))
  })
}

const handleSetRootPassword = async () => {
  await mysql.setRootPassword(rootPassword.value)
  window.$message.success($gettext('Modified successfully'))
}
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
        <the-icon :size="18" icon="material-symbols:save-outline" />
        {{ $gettext('Save') }}
      </n-button>
      <n-button
        v-if="currentTab == 'error-log'"
        class="ml-16"
        type="primary"
        @click="handleClearErrorLog"
      >
        <the-icon :size="18" icon="material-symbols:delete-outline" />
        {{ $gettext('Clear Log') }}
      </n-button>
      <n-button
        v-if="currentTab == 'slow-log'"
        class="ml-16"
        type="primary"
        @click="handleClearSlowLog"
      >
        <the-icon :size="18" icon="material-symbols:delete-outline" />
        {{ $gettext('Clear Slow Log') }}
      </n-button>
    </template>
    <n-tabs v-model:value="currentTab" type="line" animated>
      <n-tab-pane name="status" :tab="$gettext('Running Status')">
        <n-flex vertical>
          <service-status service="mysqld" />
          <n-card :title="$gettext('Root Password')">
            <n-flex vertical>
              <n-input
                v-model:value="rootPassword"
                type="password"
                show-password-on="click"
              ></n-input>
              <n-button type="primary" @click="handleSetRootPassword">{{
                $gettext('Save Changes')
              }}</n-button>
            </n-flex>
          </n-card>
        </n-flex>
      </n-tab-pane>
      <n-tab-pane name="config" :tab="$gettext('Modify Configuration')">
        <n-flex vertical>
          <n-alert type="warning">
            {{
              $gettext(
                'This modifies the MySQL main configuration file. If you do not understand the meaning of each parameter, please do not modify it randomly!'
              )
            }}
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
        </n-flex>
      </n-tab-pane>
      <n-tab-pane name="load" :tab="$gettext('Load Status')">
        <n-data-table
          striped
          remote
          :scroll-x="400"
          :loading="false"
          :columns="loadColumns"
          :data="load"
        />
      </n-tab-pane>
      <n-tab-pane name="run-log" :tab="$gettext('Runtime Logs')">
        <realtime-log service="mysqld" />
      </n-tab-pane>
      <n-tab-pane name="slow-log" :tab="$gettext('Slow Query Log')">
        <realtime-log :path="slowLog" />
      </n-tab-pane>
    </n-tabs>
  </common-page>
</template>
