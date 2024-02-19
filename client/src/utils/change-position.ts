import { Direction } from '../types/direction'
import { MessageType } from '../types/message-type'

export const changePosition = (event: KeyboardEvent, ws: WebSocket, name: string): void => {
    const direction = getDirection(event.key)

    if (!direction) {
        return
    }

    const sendingInfo  = {
        name,
        direction: direction,
        type: MessageType.ChangeDirection,
    }

    ws.send(JSON.stringify(sendingInfo))
}


const getDirection = (key: string): Direction | null => {
    switch (key) {
        case "w":
            return Direction.Top
        case "d":
            return Direction.Right
        case "s":
            return Direction.Bottom
        case "a":
            return Direction.Left
    
        default:
            return null
    }
}
