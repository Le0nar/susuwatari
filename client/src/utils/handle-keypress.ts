export const handleKeypress = (event: KeyboardEvent,ws: WebSocket): void => {
    ws.send(event.key)
}
