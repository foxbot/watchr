const m = require("mithril");
const Chat = require("./views/Chat");
const Media = require("./views/Media");

const app = {
    view() {
        return m(".app-mount", [
            m(Media),
            m(Chat)
        ])
    }
}

m.mount(document.body, app)