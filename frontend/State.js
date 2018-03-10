const Socket = require('./Socket');

const StateClass = class {
    constructor() {
        this.data = {
            v: 0,
            d: "<no media url>"
        };
        this.socket = new Socket(this);
    }

    increment() {
        this.data.v = this.data.v + 1;
    }
}

const State = new StateClass();
window.state = State;

module.exports = State