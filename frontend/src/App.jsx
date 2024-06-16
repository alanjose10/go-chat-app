import { useState } from "react";
import { sendMsg } from "./api"

function App() {

  const [message, setMessage] = useState("");

  const send = () => {
    console.log("Sending message")
    sendMsg(message);
  }

  const handleInputChange = e => {
    setMessage(e.target.value);
  }

  return (
    <>
      <input className="m-4 p-4 border" value={message} onChange={handleInputChange}></input>
      <button className="m-4 p-4 border" onClick={send}>Send</button>
    </>
  )
}

export default App
