<script setup lang="ts">
const { t } = useI18n()
const { address, hasMetamask, connectWallet, setupWallet } = useWallet()
const userStore = useUserStore()
async function connect() {
  await connectWallet()
}
async function clickLogIn() {
  await userStore.logIn()
}
async function clickLogOut() {
  await userStore.logOut()
}
onMounted(async () => {
  await setupWallet()
})
</script>

<template>
  <p text-3xl>
    {{ t('account.title') }}
  </p>
  <div v-if="!hasMetamask">
    <span>{{ t('account.need-metamask') }}</span>
  </div>
  <div v-else-if="address">
    <p my-20>
      {{ t("account.address") }}: {{ address }}
    </p>
    <div v-if="userStore.isLogIn">
      <button
        w-full rounded bg-pink-300 px-4 py-2 font-bold text-gray-800
        hover="bg-pink-400"
        @click="clickLogOut"
      >
        {{ t('button.log-out') }}
      </button>
    </div>
    <div v-else>
      <button
        w-full rounded bg-gray-300 px-4 py-2 font-bold text-gray-800
        hover="bg-gray-400"
        @click="clickLogIn"
      >
        {{ t('button.log-in') }}
      </button>
    </div>
  </div>
  <div v-else>
    <button
      class="inline-flex items-center rounded bg-gray-300 px-4 py-2 font-bold text-gray-800 hover:bg-gray-400"
      @click="connect"
    >
      <span>{{ t('button.connect') }}</span>
    </button>
  </div>
</template>
