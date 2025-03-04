import React, { useState, useEffect, FormEvent } from 'react';
import { useParams } from 'react-router-dom';

const GameRoom: React.FC = () => {
  const { gameID: gameID } = useParams<{ gameID: string }>();
  const [playerID, setPlayerID] = useState(null);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [message, setMessage] = useState('');

  useEffect(() => {
    const joinGame = async () => {
      const response = await fetch(`/game/${gameID}/join`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });

      if (!response.ok) {
        throw new Error("Failed to join game");
      }

      const data = await response.json();
      setPlayerID(data.player_id);

      const ws = new WebSocket(`ws://localhost:8080/game/${gameID}`);
      setSocket(ws);

      ws.onopen = () => {
        console.log("connected");
        setSocket(ws);
  
        const joinMessage = {
          type: 'join',
          game: gameID,
          player_id: playerID,
          Content: "veeti",
        };
        ws.send(JSON.stringify(joinMessage));
      };
  
      ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setMessages((prevMessages) => [...prevMessages, data.message || JSON.stringify(data)]);
      };
  
      ws.onerror = (error) => {
        console.error("ws error:", error);
      };
  
      ws.onclose = () => {
        console.log("disconnected");
        setSocket(null);
      };
  
      return () => {
        if (ws.readyState === 1) {
          ws.close();
        }
      };
    };

    joinGame();
  }, [gameID]);

  // Function to send a message
  const sendMessage = (e: FormEvent) => {
    e.preventDefault();
    
    if (socket && socket.readyState === WebSocket.OPEN && message.trim()) {
      const messageObject = {
        type: 'move',
        content: message
      };
      socket.send(JSON.stringify(messageObject));
      setMessage('');
    } else {
      console.error("WebSocket is not open or message is empty.");
    }
  };

  return (
    <div>
      <div className="messages">
        {messages.map((text, index) => (
          <div key={index} className="message">
            {text}
          </div>
        ))}
      </div>
      <form onSubmit={sendMessage} className="message-form">
        <input
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Type a message..."
        />
        <button type="submit">Send</button>
      </form>
    </div>
  );
};

export default GameRoom;