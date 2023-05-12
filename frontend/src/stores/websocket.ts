import { acceptHMRUpdate, defineStore } from 'pinia'

export const useWebSocketStore = defineStore('websocket', () => {
  const balanceStore = useBalanceStore()
  const connect = () => {
    useWebSocket(`ws://${location.host}/ws/`, {
      autoClose: true,
      autoReconnect: { retries: 3 },
      onConnected: () => {
      },
      onDisconnected: () => {
      },
      onMessage: (ws: WebSocket, event: MessageEvent) => {
        console.log(event)
        const msg = JSON.parse(event.data)
        console.log(msg)
        switch (msg.action) {
          case 'new-balances':
            balanceStore.addNewBalances(msg.balances)
            break
        }
      },
    })
  }
  return { connect }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useWebSocketStore as any, import.meta.hot))
