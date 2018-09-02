import m from 'mithril'
import { Component } from './application.coffee'

console.log "watchr v0 / mithril #{m.version} / foxbot.me"

document.addEventListener 'DOMContentLoaded', ->
    root = document.getElementById 'root'

    if !window.WebSocket?
        m.render root, 'This application requires WebSocket support, sorry!'
        return
    
    main = new Component
    m.route root, "/room/lobby", {
        "/room/:id": main,
    }

window.mithril = m

if module.hot
    module.hot.accept ->
        console.log 'hotload?'
        m.redraw()