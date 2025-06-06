<script setup lang="ts">
defineOptions({
  name: 'setting-index'
})
import { NButton } from 'naive-ui'
import { useGettext } from 'vue3-gettext'

import setting from '@/api/panel/setting'
import TheIcon from '@/components/custom/TheIcon.vue'
import { useThemeStore } from '@/store'
import CreateModal from '@/views/setting/CreateModal.vue'
import SettingBase from '@/views/setting/SettingBase.vue'
import SettingSafe from '@/views/setting/SettingSafe.vue'
import SettingUser from '@/views/setting/SettingUser.vue'

const { $gettext } = useGettext()
const themeStore = useThemeStore()
const currentTab = ref('base')
const createModal = ref(false)

const { data: model } = useRequest(setting.list, {
  initialData: {
    name: '',
    channel: 'stable',
    locale: 'en',
    port: 8888,
    entrance: '',
    offline_mode: false,
    two_fa: false,
    lifetime: 0,
    bind_domain: [],
    bind_ip: [],
    bind_ua: [],
    website_path: '',
    backup_path: '',
    https: false,
    cert: '',
    key: ''
  }
})

const handleSave = () => {
  useRequest(setting.update(model.value)).onSuccess(() => {
    window.$message.success($gettext('Saved successfully'))
    if (model.value.locale !== themeStore.locale) {
      themeStore.setLocale(model.value.locale)
      window.$message.info($gettext('Panel is restarting, page will refresh in 3 seconds'))
      setTimeout(() => {
        window.location.reload()
      }, 3000)
    }
  })
}

const handleCreate = () => {
  createModal.value = true
}
</script>

<template>
  <common-page show-footer>
    <template #action>
      <n-button v-if="currentTab != 'user'" type="primary" @click="handleSave">
        <the-icon :size="18" icon="material-symbols:save-outline" />
        {{ $gettext('Save') }}
      </n-button>
      <n-button v-if="currentTab == 'user'" type="primary" @click="handleCreate">
        <the-icon :size="18" icon="material-symbols:add" />
        {{ $gettext('Create User') }}
      </n-button>
    </template>
    <n-tabs v-model:value="currentTab" type="line" animated>
      <n-tab-pane name="base" :tab="$gettext('Basic')">
        <setting-base v-model:model="model" />
      </n-tab-pane>
      <n-tab-pane name="safe" :tab="$gettext('Safe')">
        <setting-safe v-model:model="model" />
      </n-tab-pane>
      <n-tab-pane name="user" :tab="$gettext('User')">
        <setting-user />
      </n-tab-pane>
    </n-tabs>
  </common-page>
  <create-modal v-model:show="createModal" />
</template>

<style scoped lang="scss"></style>
