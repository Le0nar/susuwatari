import { FC } from "react";

interface User {
    name: string,
    position: {
        x: number,
        y: number,
    }
}

export const PlayingField: FC = () => {
    return (
        <div>
            <div style={{
                position: "absolute",
                top: "500px",
                right: "500px"
                }}>
                <div>NAME</div>
                <img src="" alt="aaa" />
            </div>
        </div>
    )
}