import { useEffect } from "react"
import { initWS } from "./utils/init-ws";
import { handleKeypress } from "./utils/handle-keypress";

const ws = initWS()

function App() {

  useEffect(() => {
    document.addEventListener("keypress", (event) => handleKeypress(event, ws));
  }, [])

  return (
    <div>
      CONTENT
    </div>
  )
}

export default App
