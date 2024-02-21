import { FC, useEffect, useMemo, useState } from "react";

interface User {
    name: string,
    position: {
        x: number,
        y: number,
    }
}

interface PlayingFieldProps {
    ws: WebSocket | null
}



export const PlayingField:FC<PlayingFieldProps> = ({ ws }) => {
    const [users, setUsers] = useState<User[]>([])



    useEffect(() => {
        if (!ws) {
            return
        }

        ws.onmessage = (msg) => {
            const users: User[] | unknown = JSON.parse(msg.data)
            if(Array.isArray(users)) {
                setUsers(users)
            }
        }
    }, [ws])

    if (!ws) {
        return "No connection"
    }

    return (
        <div>
            {users.map(({name, position}) => {
                const top = 500 - (position.y * 10)
                const right = 500 - (position.x * 10)
                return (
                    <div style={{
                        position: "absolute",
                        top: `${top}px`,
                        right: `${right}px`,
                        }}>
                        <div>{name}</div>
                        <img style={{width: "10px"}} src="" alt="aaa" />
                    </div>
                )
            })}
        </div>
    )
}