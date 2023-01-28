import { Page } from './page'
import { PreventDefaultPlugin, StopPropagationPlugin } from './plugins'

export * from './page'

export const page = new Page(StopPropagationPlugin, PreventDefaultPlugin)
