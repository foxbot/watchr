const m = require("mithril");
const state = require("../State");

const Chat = {
    view() {
        return m('.chat', { onclick() { state.increment() } }, 'the state he say:' + state.data.v);
    },
    oninit(vnode) {
        state.socket.start();
    }
}

module.exports = Chat