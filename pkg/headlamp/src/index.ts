import { Page } from './page'
import { PreventDefaultPlugin, StopPropagationPlugin } from './plugins'

export * from './page'

export const hive = new Page(StopPropagationPlugin, PreventDefaultPlugin)
