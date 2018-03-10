const Path = "ws://" + window.location.host + "/api";

const Op = {
    Hello: 0,
    Identify: 1,
    ChangeMedia: 2,
    Chat: 3,
    Write: 4
};

const Socket = class {
    constructor(state) {
        this.ws = {
            readyState: WebSocket.CLOSED
        };
        this.state = state;
    }

    start() {
        this.ws = new WebSocket(Path);

        this.ws.addEventListener('open', function(event) {
            console.log('mfer opened my dude');
        });
        this.ws.addEventListener('message', function(event) {
            this.onMessage(event);
        });
        this.ws.addEventListener('close', function(event) {
            console.log('mfer closed my dude');
        });
    }
    
    onMessage(event) {
        console.log(event);
        event = JSON.parse(event.data);
        console.log(event);

        if (!event.hasOwnProperty('op')) {
            console.error('server sent invalid payload', event);
            return;
        }

        switch (event.op) {
        case Op.Hello:
            this.onHello(event);
            break;
        case Op.ChangeMedia:
            this.onChangeMedia(event);
            break;
        default:
            console.log('unhandled opcode', event);
            break;
        }
    }
    onHello(event) {
        this.ws.send({
            op: Op.Identify,
            d: {
                channel: 'home'
            }
        });
    }
    onChangeMedia(event) {
        this.state.data.mediaUrl = event.d.media_url;
    }
}

module.exports = Socket