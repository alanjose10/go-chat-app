import { useState } from "react";
import { Label, Textarea, Button } from "flowbite-react";

import { sendMsg } from "@/utils/ws"

const ChatInput = () => {

    const [message, setMessage] = useState("");

    const handleSendMessage = () => {
      console.log("Sending message")
      sendMsg(message);
      setMessage("");
    }
  
    const handleInputChange = e => {
      setMessage(e.target.value);
    }

    return (
        <>
            <div className="max-w-md m-2 p-2 border border-red-500">
                <div className="mb-2 block">
                    <Label htmlFor="comment" value="Your message" />
                </div>
                <Textarea id="comment" placeholder="Leave a comment..." required rows={4} onChange={handleInputChange} value={message}/>
                <div className="my-4">
                    <Button type="submit" onClick={handleSendMessage}>Send</Button>
                </div>
                
            </div>
            
        </>
    )
}

export default ChatInput;