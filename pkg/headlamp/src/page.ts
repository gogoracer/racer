import { executeScripElements } from './utils'

import {
  ContentType,
  Diff,
  Diffs,
  DiffType,
  FromClient,
  FromClient_File,
  FromClient_Type,
  SessionInfo,
  ToClient,
} from '@gogoracer/gas/protocol'

type ElementFunc = (el: Element, fromClientToMutate?: Diff) => void
type EventFunc = (el: Event, fromClientToMutate?: FromClient) => void
type PluginFunc = (page: Page) => void

export class Page {
  debug = true
  reconnectLimit = 0
  reconnectCount = 0
  initSyncDone = false
  sessionID = 1n
  connection!: WebSocket
  afterMessage = new Map<string, () => void>()
  beforeRemoveEventHandlers = new Map<string, ElementFunc>()
  afterRemoveEventHandlers = new Map<string, ElementFunc>()
  beforeSendEvent = new Map<string, EventFunc>()
  beforeProcessMessage = new Map<string, (fromClientToMutate: Diff) => void>()

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

    this.connection.binaryType = "arraybuffer";

    this.connection.onopen = (evt) => {
      this.log('connection: open')
      this.reconnectCount = 0
    }

    this.connection.onmessage = (evt) => {
      const buf = evt.data as ArrayBuffer
      if (!buf.byteLength)
        throw new Error('MessageEvent does not support arrayBuffer')
      const view = new Uint8Array(evt.data)
      const toClient = ToClient.fromBinary(view)

      switch (toClient.message.oneofKind) {
        case 'diffs':
          this.processDiffsMessage(toClient.message.diffs)
          break
        case 'sessionInfo':
          this.processSessionInfoMessage(toClient.message.sessionInfo)
          break
        default:
          throw new Error('Unknown message type')
      }

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

  eventHandler = (e: Event) => {
    const eventTarget = e.currentTarget as HTMLElement
    if (!eventTarget?.getAttribute) return
    const hon = eventTarget.getAttribute('hon')
    const pairs = hon?.split(',') || []
    for (let i = 0; i < pairs.length; i++) {
      const [handlerID, eventName] = pairs[i].split('|')
      if (eventName.toLowerCase() === e.type.toLowerCase()) {
        this.eventHandlerHelper(e, handlerID, false)
      }
    }
  }

  removeEventHandlers(el: Element) {
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

  eventHandlerHelper(evt: Event, handlerID: string, isInitial: boolean) {
    const el = evt.currentTarget

    const fromClient = FromClient.create({
      type: FromClient_Type.EVENT,
      ids: [handlerID],
    })

    const inputEl = el as HTMLInputElement
    if (el instanceof HTMLInputElement) {
      fromClient.data.value = String(inputEl.value)
      if (isInitial) {
        fromClient.data.init = 'true'
      }
    }

    const anyEl = el as any
    if (anyEl?.selected) {
      debugger
      fromClient.selected = true
    }

    if (inputEl?.checked) {
      fromClient.selected = true
    }

    if (anyEl.selectedOptions && anyEl.selectedOptions.length !== 0) {
      debugger
      // TODO: Looks like the server doesn't use VM at all
      // fromClient.vm = []
      // for (let i = 0; i < el.selectedOptions.length; i++) {
      //   const opt = el.selectedOptions[i]
      //   fromClient.vm.push(opt.value || opt.text)
      // }
    }

    // TODO: this is terrible, needs to be fix it ASAP
    if (evt instanceof KeyboardEvent) {
      fromClient.data.key = evt.key
      fromClient.data.charCode = evt.code
      fromClient.data.keyCode = evt.key
      fromClient.data.shiftKey = String(evt.shiftKey)
      fromClient.data.altKey = String(evt.altKey)
      fromClient.data.ctrlKey = String(evt.ctrlKey)
    }

    // File
    // TODO: move out to plugin
    if (el instanceof HTMLInputElement && el?.files?.length) {
      // No files
      const f = FromClient_File.create()
      fromClient.file = f

      // Single file
      if (el.files.length === 1) {
        const { name, size, type } = el.files[0]
        const index = 0
        const total = 1
        f.name = name
        f.sizeBytes = BigInt(size)
        f.mimeType = type
        f.totalFilesIndex = index
        f.totalFileCount = total
        // TODO: Why aren't we sending the file?
      }
      // Multiple files
      // Need to send multiple messages
      if (el.files.length > 1) {
        for (let i = 0; i < el.files.length; i++) {
          const { name, size, type } = el.files[i]
          const index = i
          const total = el.files.length

          f.name = name
          f.sizeBytes = BigInt(size)
          f.mimeType = type
          f.totalFilesIndex = index
          f.totalFileCount = total

          this.sendMsg(fromClient)
        }

        return
      }
    }

    this.beforeSendEvent.forEach((fn) => fn(evt, fromClient))
    this.sendMsg(fromClient)
  }

  sendMsg(fromClient: FromClient) {
    queueMicrotask(() => {
      // https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState
      // TODO: maybe add to a retry queue?
      if (this.connection.readyState === 1) {
        const buf = FromClient.toBinary(fromClient)
        this.connection.send(buf)
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

        const fromClient = FromClient.create({
          type: FromClient_Type.EVENT,
          file: FromClient_File.create({
            name: file.name,
            sizeBytes: BigInt(file.size),
            mimeType: file.type,
            totalFilesIndex: i,
            totalFileCount: inputEl.files.length,
          }),
        })

        queueMicrotask(() => {
          for (let j = 0; j < ids['upload'].length; j++) {
            fromClient.ids = [ids['upload'][j]]
            const f = inputEl?.files?.[i]
            if (!f) return

            const buf = FromClient.toBinary(fromClient)
            this.connection.send(buf)
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

        if (map[eventName] == undefined) {
          map[eventName] = []
        }

        map[eventName] = [...map[eventName], eventID]
      }
    }

    return map
  }

  findDiffTarget(diff: Diff) {
    let target: Node | null
    switch (diff.root.oneofKind) {
      case 'document':
        // TODO: this is a hack, my eyes are bleeding
        target = document
        break
      case 'elementSelector':
        const possibleTarget = document.querySelector(
          `[hid="${diff.root.elementSelector}"]`,
        )

        if (!possibleTarget) {
          debugger
          throw new Error(
            `root element not found: ${diff.root.elementSelector}`,
          )
        }
        target = possibleTarget
        break
      default:
        throw new Error(`unknown root type: ${diff.root.oneofKind}`)
    }

    for (let j = 0; j < diff.pathIndicies.length; j++) {
      // Doesn't exist
      if (
        diff.diffType === DiffType.CREATE &&
        (diff.contentType === ContentType.HTML ||
          diff.contentType === ContentType.TEXT) &&
        j === diff.pathIndicies.length - 1
      ) {
        continue
      }

      // Happens when we start the path for a new component
      if (j >= diff.pathIndicies.length) continue

      const htmlElementTarget = target as HTMLElement

      // Skip and child nodes found above the head.
      // Often added by browser plugins
      if (htmlElementTarget?.tagName === 'HTML') {
        for (let i = 0; i < target.childNodes.length; i++) {
          const child = target.childNodes[i] as HTMLElement
          if (child?.tagName !== 'HEAD') {
            diff.pathIndicies[j]++
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

      if (diff.pathIndicies[j] >= target.childNodes.length) {
        debugger
        this.log('child not found at section : ' + j + ' : for: ' + diff)
        target = null
        break
      }

      target = target.childNodes[diff.pathIndicies[j]]
    }

    return target
  }

  processSessionInfoMessage(sessionInfo: SessionInfo) {
    this.sessionID = sessionInfo.id
  }

  processDiffsMessage(diffs: Diffs) {
    diffs.values.forEach((diff) => {
      this.beforeProcessMessage.forEach((fn) => fn(diff))

      // DOM Diffs

      const target = this.findDiffTarget(diff) as Element
      if (!target) return

      // Text
      if (diff.contentType === ContentType.TEXT) {
        if (diff.diffType === DiffType.CREATE) {
          let element = document.createTextNode(diff.contents)

          const index = diff.pathIndicies.at(-1)!
          if (index < target.childNodes.length) {
            target.insertBefore(
              element.cloneNode(true),
              target.childNodes[index],
            )
          } else {
            target.appendChild(element.cloneNode(true))
          }
        } else {
          target.textContent = diff.contents
        }
      }

      // Tag / HTML
      if (
        diff.diffType === DiffType.CREATE &&
        diff.contentType === ContentType.HTML
      ) {
        // Only a single root element is allowed
        let template = document.createElement('template')
        template.innerHTML = diff.contents
        // TODO: talk to Sam about this
        // Sam: I've not triggered this yet
        if (!template.content.firstChild) {
          console.log('template content is empty')
          return
          // debugger
          // throw new Error('template content is empty')
        }

        executeScripElements(template.content)

        const index = diff.pathIndicies.at(-1)!
        if (index < target.childNodes.length) {
          target.insertBefore(
            template.content.firstChild,
            target.childNodes[index],
          )
        } else {
          target.appendChild(template.content.firstChild)
        }
      } else if (
        diff.diffType === DiffType.UPDATE &&
        diff.contentType === ContentType.HTML
      ) {
        let template = document.createElement('template')
        template.innerHTML = diff.contents
        executeScripElements(template.content)
        target.replaceWith(template.content)
      }

      // Attributes
      if (diff.contentType === ContentType.ATTRIBUTE) {
        const attrData = diff.contents

        // We strictly control this Attribute data format
        const index = attrData.indexOf('=')
        const attrName = attrData.substring(0, index).trim()
        const attrValue = attrData.substring(index + 2, attrData.length - 1)

        if (
          diff.diffType === DiffType.CREATE ||
          diff.diffType === DiffType.UPDATE
        ) {
          if (attrName === 'hon' && diff.diffType == DiffType.UPDATE) {
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
        } else if (diff.diffType == DiffType.DELETE) {
          this.removeEventHandlers(target) // They'll be set again if only some were removed
          target.removeAttribute(attrName)
        }
      }

      // Generic delete
      if (
        diff.diffType === DiffType.DELETE &&
        diff.contentType !== ContentType.ATTRIBUTE
      ) {
        this.removeEventHandlers(target)
        target.remove()
      }
    })
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
