import { acceptHMRUpdate, defineStore } from 'pinia'
import type { Address, Balance } from '~/types'

export const useBalanceStore = defineStore('balance', () => {
  const balances = ref<Balance[]>([]) // balance changes
  const balance = ref('-0') // current balance
  const addNewBalances = (balancesIn: Balance[]) => {
    if (balancesIn.length === 0)
      return
    const arr = [...new Map([...balancesIn, ...(balances.value)]
      .map(b => [b.blockNumber, b])).values()] // dedup
    arr.sort((a, b) => -(a.blockNumber.localeCompare(b.blockNumber))) // sort
    balances.value = arr.slice(0, 100) // truncate
    balance.value = arr[0].balance // update current balance
  }
  const getBalances = async (address: Address) => {
    const { statusCode, data } = await useFetch(`/api/get-balances?address=${address}`)
      .get()
      .json()
    if (statusCode.value === 200) {
      if (data.value.balance >= 0)
        balance.value = data.value.balance
      balances.value = [] // reset
      addNewBalances(data.value.balances)
    }
  }
  return { getBalances, balance, balances, addNewBalances }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBalanceStore as any, import.meta.hot))
