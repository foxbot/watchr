import { newSystemLine } from './views/chat.coffee'

export class Commands
    constructor: (@api, @chat) ->

    onChat: (line) -> # this will be redefined by views/chat.coffee
        console.error 'onChat never redefined!'
    
    onCommand: (input) ->
        input = input.substring 1
        parts = input.split ' '
        command = parts[0].toLowerCase()
        m = '_'+command

        if this[m]?
            this[m](parts[1..])
        else
            this.onChat newSystemLine "command '#{command}' does not exist"
    
    _room: (args) ->
        name = args[0]
        if not name
            this.onChat newSystemLine 'usage: /room <room>'
            return

        @api.sendRoom name

    _name: (args) ->
        if args.length > 1
            this.onChat newSystemLine 'usage: /name <name>'
            return

        name = args[0]
        if not name
            this.onChat newSystemLine 'usage: /name <name>'
            return
        
        @api.sendName name
    
    _clear: (args) ->
        @chat.lines = []