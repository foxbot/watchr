import m from 'mithril'
import { Commands } from '../commands.coffee'

export newChatLine = (author, content) -> 
    return {
        time: new Date
        author: author
        content: content
    }

export newSystemLine = (content) ->
    return {
        system: true
        content: content
    }

export chatView = (vnode) ->
    vnode.state.api.onChat = chat.append

    commands = new Commands vnode.state.api, chat
    commands.onChat = chat.append

    chatbox.api = vnode.state.api
    chatbox.commands = commands

    m '.pane-chat', [
        m chat
        m chatbox
        m about
    ]

chat =
    lines: []
    scroll: true

    append: (line) ->

        chat.lines.push line
        m.redraw()

    view: ->
        # todo: perf, is this rerendering everything all the time?
        m '.chat.scroller', chat.lines.map (line) -> formatLine line
    
    onupdate: (vnode) ->
        if chat.scroll
            vnode.dom.scrollTop = vnode.dom.scrollHeight

formatLine = (line) ->
    l = if line.system? then '.line.line-system' else '.line'
    return m l, [
        m 'span.line-time', "#{line.time.getHours()}:#{line.time.getMinutes()}" unless line.system?
        m 'span.line-author', line.author unless line.system?
        m 'span.line-content', line.content
    ]
    

chatbox =
    api: null
    commands: null

    view: ->
        m 'textarea.chatbox', { onkeypress: chatbox.onkey }
    
    onkey: (e) ->
        return unless e.which == 13 and !e.shiftkey

        input = e.target.value
        e.target.value = ""

        if input.startsWith '/'
            chatbox.commands.onCommand input
        else
            chatbox.api.sendChat input
        
        return false

about =
    view: ->
        m '.about', [
            m 'span.about-name', 'watchr'
            m 'span.about-ver', 'v0'
            m icons
        ]

iconList = ->
    return [
        m 'i', 't'
        m 'i', 'gh'
        m 'i', 'w'
    ]
icons =
    view: ->
        m 'span.about-icons', iconList()

# testing
###
chat.append newSystemLine 'welcome to the chat'
chat.append newChatLine 'anon', 'test 1'
chat.append newChatLine 'anon2', 'test 2'
###