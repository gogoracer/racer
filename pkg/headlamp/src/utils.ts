// Execute Script Elements https://stackoverflow.com/a/69190644/1269893
export function executeScripElements(containerElement: DocumentFragment) {
  Array.from(containerElement.querySelectorAll('script')).forEach((el) => {
    const clone = document.createElement('script')

    Array.from(el.attributes).forEach((attr) => {
      clone.setAttribute(attr.name, attr.value)
    })

    clone.text = el.text

    if (!el.parentNode) return
    el.parentNode.replaceChild(clone, el)
  })
}
