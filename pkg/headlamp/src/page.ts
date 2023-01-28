import { executeScripElements } from './utils'

interface Message {
  t: 'e'
  i: string
  file?: {
    name: string
    size: number
    type: string
    index: number
    total: number
  }
  s?: boolean
  d?: Record<string, string>
  vm?: string[]
}

type DiffPartsKey = keyof typeof diffParts
export type DiffParts = typeof diffParts[DiffPartsKey]
const diffParts = {
  DiffType: 1,
  Root: 2,
  Path: 3,
  ContentType: 4,
  Content: 5,
} as const

type MessagePartsKey = keyof typeof messageParts
export type MessageParts = typeof messageParts[MessagePartsKey]
const messageParts = {
  Type: 0,
} as const

type DiffTypeKey = keyof typeof diffTypes
export type DiffType = typeof diffTypes[DiffTypeKey]
const diffTypes = {
  Create : 'c',
  Update : 'u',
  Delete : 'd',
} as const

type ElementFunc = (el: HTMLElement, msg?: Message) => Message | undefined
type EventFunc = (el: Event, msg?: Message) => Message | undefined
type PluginFunc = (page: Page) => void

export class Page {
  debug = true
  reconnectLimit = 0
  reconnectCount = 0
  initSyncDone = false
  sessionID = 1
  connection!: WebSocket
  afterMessage = new Map<string, () => void>()
  beforeRemoveEventHandlers = new Map<string, ElementFunc>()
  afterRemoveEventHandlers = new Map<string, ElementFunc>()
  beforeSendEvent = new Map<string, EventFunc>()
  beforeProcessMessage = new Map<string, ElementFunc>()

  constructor(...plugins: PluginFunc[]) {
    // TODO: fix wails runtime
    // if (window?.runtime !== undefined) {
    //   this.connectWails()
    //   return
    // }

    if (window?.WebSocket) {
      this.connect()
    } else {
      // TODO: do something better?
      alert('Your browser does not support WebSockets')
    }

    plugins.forEach((p) => p(this))
  }

  private log(msg: string) {
    if (this.debug) {
      console.group('gogoracer')
      console.debug(msg)
      console.groupEnd()
    }
  }

  private connect() {
    let ws = 'ws'
    if (location.protocol === 'https:') {
      ws = 'wss'
    }

    // Add Session ID to query params
    const { search } = window.location
    let q = `${search ? `${search}&` : '?'}hlive=${this.sessionID}`

    const hhash = document.documentElement.getAttribute('data-hlive-hash')
    if (hhash) q += `&hhash=${hhash}`

    this.connection = new WebSocket(
      ws + '://' + window.location.host + window.location.pathname + q,
    )

    this.connection.onopen = (evt) => {
      this.log('connection: open')
      this.reconnectCount = 0
    }

    this.connection.onmessage = (evt) => {
      this.processMsg(evt)
      this.postMessage()
    }

    this.connection.onclose = (evt) => {
      this.log(
        'con: closed: ' +
          evt.reason +
          ' (' +
          evt.code +
          ') Clean: ' +
          evt.wasClean,
      )

      if (this.reconnectCount < this.reconnectLimit) {
        this.log('connection: reconnect')
        this.reconnectCount++
        this.connect()
        return
      }

      let cover = document.createElement('div')
      const s =
        'position:fixed;top:0;left:0;background:rgba(0,0,0,0.4);z-index:1000;width:100%;height:100%;'
      cover.setAttribute('style', s)

      document.getElementsByTagName('body')[0].appendChild(cover)
    }
  }

  eventHandler(e: Event) {
    const eventTarget = e.currentTarget as HTMLElement
    if (!eventTarget?.getAttribute) return

    const pairs = eventTarget.getAttribute('hon')?.split(',') || []
    for (let i = 0; i < pairs.length; i++) {
      const parts = pairs[i].split('|')
      if (parts[1].toLowerCase() === e.type.toLowerCase()) {
        this.eventHandlerHelper(e, parts[0], false)
      }
    }
  }

  removeEventHandlers(el: HTMLElement) {
    this.beforeRemoveEventHandlers.forEach((fn: ElementFunc) => fn(el))
    this.removeHLiveEventHandlers(el)
    this.afterRemoveEventHandlers.forEach((fn: ElementFunc) => fn(el))
  }

  removeHLiveEventHandlers = (el: Element) => {
    if (!el.getAttribute) return

    const val = el.getAttribute('hon')

    if (!!!val) return

    const pairs = val.split(',')
    for (let i = 0; i < pairs.length; i++) {
      const parts = pairs[i].split('|')
      const eventName = parts[1].toLowerCase()
      el.removeEventListener(eventName, this.eventHandler)
    }
  }

  eventHandlerHelper(e: Event, handlerID: string, isInitial: boolean) {
    // TODO: this is terrible, fix it ASAP
    const el = e.currentTarget as any

    let msg: Message | undefined = {
      t: 'e',
      i: handlerID,
    }

    let d: Record<string, string> = {}

    const inputEl = el as HTMLInputElement
    if (inputEl?.value) {
      d.value = String(inputEl.value)
      if (isInitial) {
        d.init = 'true'
      }
    }

    if (el?.selected) {
      msg.s = true
    }

    if (inputEl?.checked) {
      msg.s = true
    }

    if (el.selectedOptions && el.selectedOptions.length !== 0) {
      msg.vm = []
      for (let i = 0; i < el.selectedOptions.length; i++) {
        const opt = el.selectedOptions[i]
        msg.vm.push(opt.value || opt.text)
      }
    }

    // TODO: this is terrible, fix it ASAP
    const eAny = e as any
    if (eAny.key !== undefined) {
      d.key = eAny.key
      d.charCode = String(eAny.charCode)
      d.keyCode = String(eAny.keyCode)
      d.shiftKey = String(eAny.shiftKey)
      d.altKey = String(eAny.altKey)
      d.ctrlKey = String(eAny.ctrlKey)
    }

    if (d) msg.d = d

    // File
    const fileInput = el as HTMLInputElement
    // TODO: move out to plugin
    if (fileInput.files) {
      // No files
      msg.file = {
        name: '',
        size: 0,
        type: '',
        index: 0,
        total: 0,
      }
      // Single file
      if (fileInput.files.length === 1) {
        const { name, size, type } = fileInput.files[0]
        const index = 0
        const total = 1
        msg.file = { name, size, type, index, total }
      }
      // Multiple files
      // Need to send multiple messages
      if (fileInput.files.length > 1) {
        for (let i = 0; i < fileInput.files.length; i++) {
          const { name, size, type } = fileInput.files[i]
          const index = i
          const total = fileInput.files.length

          msg.file = { name, size, type, index, total }

          // TODO: Really?? not outside the loop?
          this.sendMsg(msg)
        }

        return
      }
    }

    this.beforeSendEvent.forEach((fn) => (msg = fn(e, msg)))

    this.sendMsg(msg)
  }

  sendMsg(msg: Message) {
    queueMicrotask(() => {
      // https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState
      // TODO: maybe add to a retry queue?
      if (this.connection.readyState === 1) {
        this.connection.send(JSON.stringify(msg))
      }
    })
  }

  setEventHandlers() {
    document.querySelectorAll('[hon]').forEach((el) => {
      const pairs = el.getAttribute('hon')?.split(',') || []
      for (let i = 0; i < pairs.length; i++) {
        const parts = pairs[i].split('|')
        el.addEventListener(parts[1].toLowerCase(), this.eventHandler)
      }
    })
  }

  // Sync Initial Input Values
  // Looks at the current value of the input and if needed triggers events to sync that value to the backend
  // TODO: move out as plugin?
  initSync() {
    document.querySelectorAll('[hon]').forEach((el) => {
      // Radio
      const elType = el.getAttribute('type')

      if (elType === 'radio' || elType === 'checkbox') {
        const inputEl = el as HTMLInputElement
        // No sync needed
        if (inputEl.checked === inputEl.hasAttribute('checked')) {
          return
        }
      } else if (elType === 'select') {
        const selectEl = el as HTMLSelectElement
        let hit = false
        for (let i = 0; i < selectEl.selectedOptions.length; i++) {
          if (
            selectEl.selectedOptions[i].selected &&
            selectEl.hasAttribute('checked') === false
          ) {
            hit = true
            break
          }
        }
        // No sync needed
        if (hit === false) {
          return
        }
      } else {
        const inputEl = el as HTMLInputElement
        // TODO: Why the smiley face? :D
        // How'd this happen :)
        if (!inputEl?.value) {
          return
        }
        // At default
        const aV = el.getAttribute('value')
        // Empty
        if (aV === null && inputEl.value === '') {
          return
        }
        // Match
        if (aV === inputEl.value) {
          return
        }
      }

      const pairs = el?.getAttribute('hon')?.split(',') || []
      for (let i = 0; i < pairs.length; i++) {
        const parts = pairs[i].split('|')
        const name = parts[1].toLowerCase()

        if (
          name === 'keyup' ||
          name === 'keydown' ||
          name === 'keypress' ||
          name === 'input' ||
          name === 'change'
        ) {
          // TODO: this is a hack, my eyes are bleeding
          const evt = {
            currentTarget: el,
          }
          // TODO: this is terrible, fix it ASAP
          this.eventHandlerHelper(evt as unknown as Event, parts[0], true)
        }
      }
    })
  }

  postMessage() {
    this.setEventHandlers()

    if (!this.initSyncDone && document.querySelectorAll('[hon]').length !== 0) {
      this.initSyncDone = true
      this.initSync()
    }

    // Start file upload
    document.querySelectorAll('[data-hlive-upload]').forEach((el) => {
      const ids = this.getEventHandlerIDs(el)

      if (!ids['upload']) {
        return
      }

      const inputEl = el as HTMLInputElement

      if (inputEl?.files?.length) {
        let i = 0
        const file = inputEl.files[0]

        let msg: Message = {
          t: 'e',
          i: '',
          file: {
            name: file.name,
            size: file.size,
            type: file.type,
            index: i,
            total: inputEl.files.length,
          },
        }

        queueMicrotask(() => {
          for (let j = 0; j < ids['upload'].length; j++) {
            msg.i = ids['upload'][j]
            const f = inputEl?.files?.[i]
            if (!f) return
            const blob = new Blob([JSON.stringify(msg) + '\n\n', f], {
              type: f.type,
            })
            this.connection.send(blob)
          }
        })
      }
    })

    this.afterMessage.forEach((fn) => fn())
  }

  getEventHandlerIDs(el: Element) {
    let map: Record<string, string[]> = {}

    const hon = el.getAttribute('hon')
    if (hon) {
      const pairs = hon.split(',')
      for (let i = 0; i < pairs.length; i++) {
        const [eventID, eventNameRaw] = pairs[i].split('|')
        const eventName = eventNameRaw.toLowerCase()
        map[eventName] = [...map[eventName], eventID]
      }
    }

    return map
  }

  findDiffTarget(diff: string) {
    debugger

    const parts = diff.split('|')
    const root = parts[diffParts.Root]

    let docTarget = document
    let target: Element | null = null
    if (root !== 'doc') {
      target = document.querySelector(`[hid="${parts[diffParts.Root]}"]`)

      if (!target) {
        debugger
       throw new Error(`root element not found: ${parts[diffParts.Root]}`)
      }
    }

    // TODO: talk to Sam about this
    const domPath = parts[diffParts.Path]
    const domPathParts = domPath.split('>').map((s) => parseInt(s))

    for (let j = 0; j < domPathParts.length; j++) {
      const diffType = parts[diffParts.DiffType] as DiffType
      const contentType = parts[diffParts.ContentType]

      // Doesn't exist
      if (diffType === diffTypes.Create
         &&
        (contentType === 'h' || contentType === 't') &&
        j === domPathParts.length - 1
      ) {
        continue
      }

      // Happens when we start the path for a new component
      if (isNaN(domPathParts[j])) continue

      // Skip and child nodes found above the head.
      // Often added by browser plugins
      if (target?.tagName === 'HTML') {
        for (let i = 0; i < target.childNodes.length; i++) {
          const child = target.childNodes[i] as HTMLElement
          if (!child || child?.tagName !== 'HEAD') {
            domPathParts[j]++
          } else {
            break
          }
        }
      }

      // TODO: talk to Sam about this
      if (!target?.childNodes.length) {
        debugger
        this.log('no child nodes found at section : ' + j + ' : for: ' + diff)
        target = null
        break
      }

      if (domPathParts[j] >= target.childNodes.length) {
        debugger
        this.log('child not found at section : ' + j + ' : for: ' + diff)
        target = null
        break
      }

      target = target.childNodes[domPathParts[j]] as Element
    }

    return target
  }

  processMsg(evt: MessageEvent) {
    let messages = evt.data.split('\n')

    for (let i = 0; i < messages.length; i++) {
      let msg = messages[i]

      if (!msg) continue

      this.beforeProcessMessage.forEach((fn) => (msg = fn(msg)))

      if (msg === '') {
        continue
      }

      const parts = msg.split('|')

      // DOM Diffs
      if (parts[messageParts.Type] === 'd') {
        if (parts.length !== 6) {
          debugger
          this.log('invalid diff message format')
          continue
        }

        const target = this.findDiffTarget(messages[i]) as HTMLElement
        if (!target) return

        const path = parts[diffParts.Path].split('>')

        // Text
        if (parts[diffParts.ContentType] === 't') {
          if (parts[diffParts.DiffType] === 'c') {
            let element = document.createTextNode(
              base64Decode(parts[diffParts.Content]),
            )

            const index = path[path.length - 1]
            if (index < target.childNodes.length) {
              target.insertBefore(
                element.cloneNode(true),
                target.childNodes[index],
              )
            } else {
              target.appendChild(element.cloneNode(true))
            }
          } else {
            target.textContent = base64Decode(parts[diffParts.Content])
          }
        }

        // Tag / HTML
        if (
          parts[diffParts.DiffType] === 'c' &&
          parts[diffParts.ContentType] === 'h'
        ) {
          // Only a single root element is allowed
          let template = document.createElement('template')
          template.innerHTML = base64Decode(parts[diffParts.Content])

          // TODO: talk to Sam about this
          if (!template.content.firstChild) {
            debugger
            this.log('template content is empty')
            continue
          }

          executeScripElements(template.content)

          const index = path[path.length - 1]
          if (index < target.childNodes.length) {
            target.insertBefore(
              template.content.firstChild,
              target.childNodes[index],
            )
          } else {
            target.appendChild(template.content.firstChild)
          }
        } else if (
          parts[diffParts.DiffType] === 'u' &&
          parts[diffParts.ContentType] === 'h'
        ) {
          let template = document.createElement('template')

          // TODO: talk to Sam about this
          if (!template.content.firstChild) {
            debugger
            this.log('template content is empty')
            continue
          }

          template.innerHTML = base64Decode(parts[diffParts.Content])
          target.replaceWith(template.content.firstChild)
        }

        // Attributes
        if (parts[diffParts.ContentType] === 'a') {
          const attrData = base64Decode(parts[diffParts.Content])

          // We strictly control this Attribute data format
          const index = attrData.indexOf('=')
          const attrName = attrData.substring(0, index).trim()
          const attrValue = attrData.substring(index + 2, attrData.length - 1)

          if (
            parts[diffParts.DiffType] === 'c' ||
            parts[diffParts.DiffType] === 'u'
          ) {
            if (attrName === 'hon' && parts[diffParts.DiffType] === 'u') {
              // They'll be set again if only some were removed
              this.removeEventHandlers(target)
            }

            if (attrName === 'value') {
              if (target === document.activeElement && attrValue !== '') {
                // Don't update when someone is typing
              } else {
                debugger
                // TODO: this is terrible, fix it ASAP
                const targetAny = target as any
                targetAny.value = attrValue
              }
            } else {
              target.setAttribute(attrName, attrValue)
            }
          } else if (parts[diffParts.DiffType] === 'd') {
            // They'll be set again if only some were removed
            this.removeEventHandlers(target)

            target.removeAttribute(attrName)
          }
        }

        // Generic delete
        if (
          parts[diffParts.DiffType] === 'd' &&
          parts[diffParts.ContentType] !== 'a'
        ) {
          this.removeEventHandlers(target)
          target.remove()
        }
        // Sessions
      } else if (parts[messageParts.Type] === 's') {
        if (parts.length === 3) {
          this.sessionID = parts[2]
        }
      }
    }
  }

  connectWails() {
    debugger
    console.error(
      'TODO: not sure what to do here yet, seems like a plugin kind of thing',
    )
    // this.connection.conn = {
    //   readyState: 1,
    //   send: function (msg) {
    //     runtime.EventsEmit('out', msg)
    //   },
    // }

    // runtime.EventsOn('in', function (evt) {
    //   hlive.processMsg({ data: evt })
    //   hlive.postMessage()
    // })

    // runtime.EventsEmit('connect', true)
  }
}

// Base64 decode with unicode support
// Ref: https://stackoverflow.com/questions/30106476/using-javascripts-atob-to-decode-base64-doesnt-properly-decode-utf-8-strings
function base64Decode(str: string) {
  // Going backwards: from byte stream, to percent-encoding, to original string.
  return decodeURIComponent(
    window
      .atob(str)
      .split('')
      .map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
      })
      .join(''),
  )
}
