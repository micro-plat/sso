import ReconnectingWebSocket from 'reconnecting-websocket';
const wsURL = process.env.service.ws

export function ws() {
    const options = {
        connectionTimeout: 1000,
    };
    const rws = new ReconnectingWebSocket(wsURL, [], options);
    return rws
}


export function wssend(ws, message) {
    this.waitForExecute = function(callback) {
        if (ws && ws.readyState && ws.readyState === 1) {
            callback();
        } else {
            var that = this;
            setTimeout(function() {
                that.waitForExecute(callback);
            }, 1000);
        }
    }

    this.waitForExecute(function() {
        console.log("send:", message)
        ws.send(message);
    });
};