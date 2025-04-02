import { useState, useCallback, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const Game: React.FC = () => {
    const playerID: Number = Number(sessionStorage.getItem("playerID"));
    const gameID: Number = Number(sessionStorage.getItem("gameID"));

    const socketUrl = 'ws://localhost:8080/ws/game';
    const [messageHistory, setMessageHistory] = useState<MessageEvent<any>[]>([]);
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
        shouldReconnect: () => true,
        reconnectAttempts: 10,
        reconnectInterval: 3000,
    });
    const [userMessage, setUserMessage] = useState('');

    useEffect(() => {
        if (readyState === ReadyState.OPEN) {
            const joinMessage = {
              type: 'join',
              game_id: Number(gameID),
              player_id: Number(playerID),
              Content: "placeholder_name",
            };

            const json = JSON.stringify(joinMessage);
            console.log("Sending: ", json);

            sendMessage(json);
        }
    }, [readyState]);

    useEffect(() => {
        if (lastMessage !== null) {
            setMessageHistory((prev) => prev.concat(lastMessage));
        }
    }, [lastMessage]);

    const handleClickSendMessage = useCallback((event: React.FormEvent) => {
        event.preventDefault();
        const messageObject = {
            type: 'move',
            game_id: gameID,
            player_id: playerID,
            content: userMessage
          };

        const json = JSON.stringify(messageObject)
        console.log("Sending: ", json);

        sendMessage(json);
        setUserMessage('');
    }, [userMessage]);
    
    return (
        <div>
            <h1>Chess game</h1>
            <div>Game ID: {String(gameID)}</div>
            <form onSubmit={handleClickSendMessage}>
                <input
                    type="text"
                    value={userMessage}
                    onChange={(e) => setUserMessage(e.target.value)}
                    placeholder="Type a message..."
                />
                <button type="submit" disabled={readyState !== ReadyState.OPEN}>
                    Send
                </button>
            </form>
            {lastMessage ? <div>Last message: {lastMessage.data}</div> : null}
            <h2>Message history</h2>
            <ul>
                {messageHistory.map((message, idx) => (
                    <li key={idx}>{message ? message.data : null}</li>
                ))}
            </ul>
        </div>
    );
};

export default Game;