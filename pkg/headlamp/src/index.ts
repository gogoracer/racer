import { loadAlpineOnWindow } from './alpine'
import { Page } from './page'
import { PreventDefaultPlugin, StopPropagationPlugin } from './plugins'

export * from './page'
export * from './nats'
export * from './alpine'

export const hlive = new Page(StopPropagationPlugin, PreventDefaultPlugin)

    ; (window as any).hlive = hlive


loadAlpineOnWindow()
