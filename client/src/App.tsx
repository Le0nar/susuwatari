import { useState } from "react"
import { initWS } from "./utils/init-ws";
import { changePosition } from "./utils/change-position";

function App() {
  const [name, setName] = useState("")
  const [isStarted, setIsStarted] = useState(false)


  return (
    <div>
      {!isStarted && (
      <>
        <input type="text" value={name} onChange={(event) => setName(event.target.value)} />
        <button onClick={() => {
          setIsStarted(true)
          const ws = initWS(name)
          document.addEventListener("keypress", (event) => changePosition(event, ws, name));
        }}>
          Старт
        </button>
      </>
      )}

    </div>
  )
}

export default App
