import m from 'mithril'
import { chatView } from './chat.coffee'

export mainView = (vnode) -> 
    [
        m '.pane-media'
        chatView vnode
    ]