import { Page } from './page'

export function StopPropagationPlugin(p: Page) {
  if (!p.beforeSendEvent.get('hsp')) {
    p.beforeSendEvent.set('hsp', (e, msg) => {
      const target = e.currentTarget as HTMLElement
      if (!target?.getAttribute) return msg
      if (!target.hasAttribute('data-hlive-sp')) return msg
      e.stopPropagation()
      return msg
    })
  }
}

export function PreventDefaultPlugin(p: Page) {
  if (!p.beforeSendEvent.get('hpd')) {
    p.beforeSendEvent.set('hpd', function (e, msg) {
      const target = e.currentTarget as HTMLElement
      if (!target?.getAttribute) return msg
      if (!target.hasAttribute('data-hlive-pd')) return msg
      e.preventDefault()
      return msg
    })
  }
}
