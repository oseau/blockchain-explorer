import { type ViteSSGContext } from 'vite-ssg'

export type UserModule = (ctx: ViteSSGContext) => void

export type Address = string & { __brand: 'ADDRESS' }

export interface Balance {
  blockNumber: string
  balance: string
}
