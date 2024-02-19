import { MessageType } from "../types/message-type";

export const initWS = (name: string): WebSocket => {
    const ws = new WebSocket("ws://localhost:8080")
    ws.onopen = () => {
        console.log("Successfully Connected");

        const sendingInfo  = {
            name,
            type: MessageType.AddUser,
        }
    
        ws.send(JSON.stringify(sendingInfo))
    };

    ws.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    ws.onerror = error => {
        console.log("Socket Error: ", error);
    }

    return ws
}
