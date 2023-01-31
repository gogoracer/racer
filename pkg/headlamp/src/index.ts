import { loadAlpineOnWindow } from './alpine'
import { Page } from './page'
import { PreventDefaultPlugin, StopPropagationPlugin } from './plugins'

export * from './page'
export * from './nats'
export * from './alpine'

export const hive = new Page(StopPropagationPlugin, PreventDefaultPlugin)

loadAlpineOnWindow()
