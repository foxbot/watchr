import m from 'mithril'
import { Config, Op } from './config.coffee'
import { mainView } from './views/main.coffee'
import { newChatLine, newSystemLine } from './views/chat.coffee'

class State
    constructor: ->
        @room = m.route.param().id
        @name = localStorage.getItem 'name' or ''
        @media = null
    
    connect: ->
        url = "ws://#{Config.Api}/gateway"

        @sock = new WebSocket url

        _this = this
        @sock.onopen = (e) -> _this.onOpen _this, e
        @sock.onmessage = (e) -> _this.onMessage _this, e
        @sock.onclose = (e) -> _this.onClose _this, e
        @sock.onerror = (e) -> _this.onError _this, e
    
    close: ->
        @sock.close 1000

    onOpen: (me, e) ->
        me.onChat newSystemLine 'connected!'

        identify = {
            op: Op.Identify,
            data: {
                room: me.room
                name: me.name
            }
        }
        me.sock.send JSON.stringify identify

    onMessage: (me, e) ->
        data = JSON.parse e.data
        switch data.op
            when Op.Chat then me.onRawChat data.data
            when Op.UserUpdate then me.onUserUpdate data.data
            when Op.RoomUpdate then me.onRoomUpdate data.data
            else console.error 'unknown opcode!', e

    onClose: (me, e) ->
        me.onChat newSystemLine 'disconnected, try refreshing.'

    onError: (me, e) ->
        console.error e

    onRawChat: (data) ->
        this.onChat newChatLine data.author, data.content

    onUserUpdate: (data) ->
        this.setName data.name if data.name? and data.name
        this.setRoom data.room if data.room? and data.room
    
    onRoomUpdate: (data) ->
        @media = {
            media_type: data.media_type
            media: data.media
        }
        m.redraw()
    
    # this is re-set by chat.coffee
    onChat: (line) ->
        console.error 'chat never initialized...'
    
    setName: (name) ->
        @name = name
        localStorage.setItem 'name', name

        this.onChat newSystemLine "you are now known as '#{name}'"
    
    setRoom: (room) ->
        @room = room
        m.route.set "/room/#{room}"

        this.onChat newSystemLine "switched to room '#{room}'."
    
    sendChat: (content) ->
        chat = {
            op: Op.Chat,
            data: {
                content: content
            }
        }
        @sock.send JSON.stringify chat
    
    sendName: (content) ->
        payload = {
            op: Op.UserSet
            data: {
                name: content
            }
        }
        @sock.send JSON.stringify payload
    
    sendRoom: (content) ->
        payload = {
            op: Op.UserSet
            data: {
                room: content
            }
        }
        @sock.send JSON.stringify payload
        

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
    

