<script setup lang="ts">
import { isAddress } from '@ethersproject/address'
import { commify, formatEther } from '@ethersproject/units'

const { t } = useI18n()
const { address, setupWallet } = useWallet()
const balanceStore = useBalanceStore()
const wsStore = useWebSocketStore()
const addressToCheck = ref(address ?? '')
async function checkBalance() {
  if (!isAddress(address.value))
    return
  await balanceStore.getBalances(address.value)
}
onMounted(async () => {
  await setupWallet()
  wsStore.connect()
})
</script>

<template>
  <p my10 text-3xl>
    {{ t('balance.title') }}
  </p>

  <TheInput
    v-model="addressToCheck"
    my10
    w150
    :placeholder="t('balance.input-address-to-check')"
    autocomplete="false"
    @keydown.enter="checkBalance"
  />
  <label class="hidden" for="input">{{ t('balance.input-address-to-check') }}</label>

  <div>
    <button
      m-3 text-sm btn
      :disabled="!isAddress(addressToCheck)"
      @click="checkBalance"
    >
      {{ t('button.check-balance') }}
    </button>
  </div>

  <div v-if="!balanceStore.balance.startsWith('-')" class="w-150 overflow-x-auto shadow-md sm:rounded-lg">
    <table class="w-full text-left text-sm text-gray-500 dark:text-gray-400">
      <caption class="bg-white p-5 text-left text-lg font-semibold text-gray-900 dark:bg-gray-800 dark:text-white">
        {{ t('balance.table.title') }}
        <div>
          <p pull-right text-sm font-normal text-gray-500 dark:text-gray-400>
            {{ t('balance.table.current') }}:
            <span font-mono>
              {{ commify(formatEther(balanceStore.balance)) }} ETH
            </span>
          </p>
          <p text-sm font-normal text-gray-500 dark:text-gray-400>
            {{ t('balance.table.sub-title') }}
          </p>
        </div>
      </caption>
      <thead class="bg-gray-50 text-xs uppercase text-gray-700 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3">
            {{ t('balance.table.block-number') }}
          </th>
          <th scope="col" class="px-6 py-3">
            {{ t('balance.table.balance') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="balance in balanceStore.balances" class="border-b bg-white dark:border-gray-700 dark:bg-gray-800">
          <th :key="balance.blockNumber" scope="row" class="whitespace-nowrap px-6 py-4 font-mono font-medium text-gray-900 dark:text-white">
            {{ balance.blockNumber }}
          </th>
          <td class="px-6 py-4" font-mono>
            {{ balance.balance }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
