import Alpine from 'alpinejs'
import morph from '@alpinejs/morph'

export function loadAlpineOnWindow() {
  ;(window as any).Alpine = Alpine

  Alpine.plugin(morph)
  Alpine.start()
}
