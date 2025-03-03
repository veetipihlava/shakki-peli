import React, { useState, useEffect, FormEvent } from 'react';
import { useParams } from 'react-router-dom';

const GameRoom: React.FC = () => {
  const { gameId } = useParams<{ gameId: string }>();
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [message, setMessage] = useState('');

  useEffect(() => {
    const ws = new WebSocket(`ws://localhost:8080/game/${gameId}`);

    ws.onopen = () => {
      console.log("connected");
      setSocket(ws);

      /* const joinMessage = {
        type: 'join',
      };
      ws.send(JSON.stringify(joinMessage)); */
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log(data);
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
  });

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