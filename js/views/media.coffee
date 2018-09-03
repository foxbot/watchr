import m from 'mithril'
import { MediaType } from '../config.coffee'

export mediaView = (vnode) ->
    media = vnode.state.api.media
    if not media?
        return m '.pane-media.media-text'
    
    subview = switch media.media_type
        when MediaType.Text then textView media.media
        else textView '(error) unknown media type.'
    
    m '.pane-media', subview

textView = (content) ->
    m 'section.media-text', content