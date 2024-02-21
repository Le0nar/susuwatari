import { useState } from "react"
import { initWS } from "./utils/init-ws";
import { changePosition } from "./utils/change-position";
import { PlayingField } from "./components/playing-field/playing-field";

function App() {
  const [name, setName] = useState("")
  const [isStarted, setIsStarted] = useState(false)
  const [ws, setWs] = useState<WebSocket | null>(null)

  return (
    <div>
      {isStarted 
        ? <PlayingField ws={ws} /> 
        : (
          <>
            <input type="text" value={name} onChange={(event) => setName(event.target.value)} />
            <button onClick={() => {
              setIsStarted(true)
              
              const localWs = initWS(name)
              document.addEventListener("keypress", (event) => changePosition(event, localWs, name));
              setWs(localWs)
            }}>
              Старт
            </button>
          </>
        )
      }
    </div>
  )
}

export default App
