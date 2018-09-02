import m from 'mithril'

export chatView = (vnode) ->
    chat.node = vnode
    chatbox.node = vnode

    m '.pane-chat', [
        m chat
        m chatbox
        m about
    ]

chat =
    node: null
    lines: []

    append: (line) ->
        chat.lines.push line
        m.redraw

    view: ->
        # todo: perf, is this rerendering everything all the time?
        m '.chat.scroller', chat.lines.map (line) -> formatLine line

formatLine = (line) ->
    l = if line.system? then '.line.line-system' else '.line'
    return m l, [
        m 'span.line-time', "#{line.time.getHours()}:#{line.time.getMinutes()}" unless line.system?
        m 'span.line-author', line.author unless line.system?
        m 'span.line-content', line.content
    ]
    

newChatLine = (author, content) -> 
    return {
        time: new Date
        author: author
        content: content
    }

newSystemLine = (content) ->
    return {
        system: true
        content: content
    }

chatbox =
    node: null

    view: ->
        m 'textarea.chatbox', { onkeypress: chatbox.onkey }
    
    onkey: (e) ->
        return unless e.which == 13 and !e.shiftkey

        input = e.target.value
        e.target.value = ""

        # TODO: sweeter command parsing
        if input.startsWith '/room '
            arg = input.substring 6
            m.route.set "/room/#{arg}"
            chat.append newSystemLine "moved to room #{arg}."
            chatbox.node.state.api.changeRoom()
        else if input.startsWith '/connect'
            chat.append newSystemLine "connecting:)"
            console.log chatbox.node.state
            chatbox.node.state.api.connect()
        else
            chat.append newChatLine 'anonymous', input
        
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