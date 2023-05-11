import { defaultWindow } from '@vueuse/core'
import { Web3Provider } from '@ethersproject/providers'
import { getAddress } from '@ethersproject/address'
import { type Address } from '~/types'

const window = defaultWindow

export const getProvider = () => new Web3Provider(window!.ethereum!, 'any')

export function assertAddress(address: string | undefined): asserts address is Address {
  if (!address)
    throw new Error(`"${address}" is not a valid address`)
  try {
    getAddress(address)
  }
  catch (e) {
    throw new Error(`"${address}" is not a valid address`)
  }
}

export function getAddressChecksum(value: any) {
  assertAddress(value)
  return getAddress(value) as Address
}

export const getAddressLowerCase = (a: any) => getAddressChecksum(a).toLowerCase() as Address

export function getAddressChecksumShort(a: any) {
  const addressChecksum = getAddressChecksum(a)
  return `${addressChecksum.substring(0, 4 + 2)}...${addressChecksum.substring(42 - 4)}`
}

function wallet() {
  let address = $ref<Address | undefined>()
  const shortAddress = $computed(() => getAddressChecksumShort(address!))
  // https://github.com/Uniswap/interface  function getIsMetaMask()
  const hasMetamask = $computed(() => window?.ethereum?.isMetaMask ?? false)
  const setAddress = (a: Address) => address = getAddressLowerCase(a)
  const connectWallet = async () => {
    const myAccounts = await getProvider().send('eth_requestAccounts', [])
    setAddress(myAccounts[0])
  }
  const getSignature = async (msg: string) =>
    await getProvider().getSigner(address).signMessage(msg)
  let skipSetup = $ref(false)
  const setupWallet = async () => {
    if (skipSetup)
      return
    skipSetup = true
    if (!hasMetamask)
      return console.error('metamask not installed!')
    const provider = getProvider()
    try {
      const myAccounts = await provider?.send('eth_accounts', [])
      if (myAccounts.length > 0)
        setAddress(myAccounts[0])
    }
    catch (e) {
      console.error(e)
    }
  }
  return $$({
    hasMetamask,
    address,
    shortAddress,
    connectWallet,
    getSignature,
    setupWallet,
  })
}

export const useWallet = createSharedComposable(wallet)
