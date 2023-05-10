<script setup lang="ts">
const { t } = useI18n()
</script>

<template>
  <div h-screen flex="~ col">
    <TheNav />
    <main
      flex="~ 1"
      items-center justify-center
      text="center gray-700 dark:gray-200"
    >
      <RouterView />
    </main>
    <div mx-auto py-10 text-center text-sm opacity-50>
      [{{ t('intro.name') }}]
    </div>
  </div>
</template>
