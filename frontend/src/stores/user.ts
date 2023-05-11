import { acceptHMRUpdate, defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  const validUntil = useStorage('validUntil', 0)
  const timestamp = useTimestamp()
  const isLogIn = computed(() => (timestamp.value / 1000) < validUntil.value)
  const { locale } = useI18n()

  const { address, getSignature } = $(useWallet())
  const getNonce = async () => {
    if (!address)
      return ''
    const { data } = await useFetch(`/api/nonce?address=${address}&lang=${locale}`)
      .get()
      .json()
    return data.value.nonce
  }
  const logIn = async () => {
    const nonce = await getNonce()
    const signature = await getSignature(nonce)
    if (!signature)
      return false
    const { statusCode, data } = await useFetch(`/api/login?address=${address}`)
      .post({ signature })
      .json()
    if (statusCode.value === 200)
      validUntil.value = data.value.validUntil
  }

  return {
    isLogIn,
    logIn,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore as any, import.meta.hot))
