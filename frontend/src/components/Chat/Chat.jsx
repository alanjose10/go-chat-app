import ChatInput from "@/components/Chat/ChatInput"
import ChatHistory from "@/components/Chat/ChatHistory"

const Chat = ( {chatHistory} ) => {
    return (
        <>
            <div className="">
                <div>
                    <ChatHistory
                        chatHistory={chatHistory}
                    />
                </div>
                <div>
                    <ChatInput/>
                </div>
                
            </div>
        </>

    )
}

export default Chat;