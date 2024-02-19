export const initWS = (): WebSocket => {
    const ws = new WebSocket("ws://localhost:8080")
    ws.onopen = () => {
        console.log("Successfully Connected");
    };

    ws.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    ws.onerror = error => {
        console.log("Socket Error: ", error);
    }

    return ws
}
