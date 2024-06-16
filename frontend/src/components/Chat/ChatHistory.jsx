

const ChatHistory = ({ chatHistory }) => {

    console.log(chatHistory);

    return (
        <div className="m-2 p-2 border border-red-500">
            <p>Chat History</p>
            {chatHistory.map((message, index) => <p key={index}>{message.data}</p>)}
        </div>
    )
}

export default ChatHistory;