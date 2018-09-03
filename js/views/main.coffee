import m from 'mithril'
import { chatView } from './chat.coffee'
import { mediaView } from './media.coffee'

export mainView = (vnode) -> 
    [
        mediaView vnode
        chatView vnode
    ]