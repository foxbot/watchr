const m = require("mithril");
const state = require("../State");

const Media = {
    view() {
        return m(".media", { onclick() { state.increment() } }, [
            m("p", "media testlet - the state he say:" + state.data.v),
            m("p", "the websocket is big: " + state.socket.ws.readyState),
            m("p", "media url he say", state.data.mediaUrl)
        ])
    }
}

module.exports = Media