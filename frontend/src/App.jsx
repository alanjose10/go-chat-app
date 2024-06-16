
import Header from "./components/Header";
import Chat from "./components/Chat/Chat";
import { useEffect, useState } from "react";

import { connect } from "@/utils/ws"

function App() {

  const [chatHistory, setChatHistory] = useState([]);

  console.log(chatHistory);

  useEffect(() => {
    connect(
      (msg) => {
        console.log("New message: ", msg);
        setChatHistory(prevState => [...prevState, msg])
      }
    )
  }, []);



  return (
    <>
    <div className="h-screen bg-sky-200">
      <Header/>
      <Chat
        chatHistory={chatHistory}
      />
    </div>
      
    </>
  )
}

export default App
