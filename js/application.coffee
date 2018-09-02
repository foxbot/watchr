import m from 'mithril'
import { mainView } from './views/main.coffee'
import { Config, Op } from './config.coffee'

class State
    constructor: ->
        @room = m.route.param().id
    
    connect: ->
        url = "ws://#{Config.Api}/gateway"
        console.log 'connecting on', url

        @sock = new WebSocket url

        _this = this
        @sock.onopen = (e) -> _this.onOpen _this, e
        @sock.onmessage = (e) -> _this.onMessage _this, e
        @sock.onclose = (e) -> _this.onClose _this, e
        @sock.onerror = (e) -> _this.onError _this, e
    
    close: ->
        @sock.close(1000)

    onOpen: (me, e) ->
        console.log e
        identify = {
            Op: Op.Identify,
            Data: {
                Room: me.room
            }
        }
        me.sock.send(identify)

    onMessage: (me, e) ->
        console.log e

    onClose: (me, e) ->
        console.log e
    
    onError: (me, e) ->
        console.error e
    
    onChat: (line) ->
        console.error 'chat never initialized...'
    
    updateRoom: ->
        @room = m.route.param().id
        

export class Component
    oninit: (vnode) ->
        state = new State
        state.connect()

        vnode.state.api = state

        # TODO: remove
        window.state = state
    
    onbeforeremove: (vnode) ->
        vnode.state.api.close()
    
    view: (vnode) ->
        mainView vnode
    

