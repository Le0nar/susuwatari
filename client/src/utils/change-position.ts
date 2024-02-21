import { Direction } from '../enum/direction'
import { MessageType } from '../enum/message-type'

export const changePosition = (event: KeyboardEvent, ws: WebSocket, name: string): void => {
    const direction = getDirection(event.key)

    if (!direction) {
        return
    }

    const sendingInfo  = {
        name,
        direction: direction,
        type: MessageType.ChangePosition,
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
