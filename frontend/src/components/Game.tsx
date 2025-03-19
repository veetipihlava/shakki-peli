import React, { useState, useEffect, FormEvent } from 'react';
import { useParams } from 'react-router-dom';

const Game: React.FC = () => {
  const { gameID } = useParams();
  const [playerID, setPlayerID] = useState<string | null>(null);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [recievedMessages, setRecievedMessages] = useState<string[]>([]);
  const [userMessage, setUserMessage] = useState('');

  useEffect(() => {
    console.log("getting player with game id", gameID);
    const storedPlayerID = sessionStorage.getItem("player_id");
    if (!storedPlayerID) {
      const getPlayer = async () => {
        try {
          const response = await fetch(`/game/${gameID}/join`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
          });

          if (!response.ok) {
            throw new Error("Failed to join game");
          }

          const data = await response.json();
          sessionStorage.setItem("player_id", data.player_id);
          setPlayerID(data.player_id);
        } catch (error) {
          console.error(error);
        }
      };

      getPlayer();
    } else {
      setPlayerID(storedPlayerID);
    }
  }, [gameID]);

  useEffect(() => {
    if (!playerID) return;

    const joinGame = async () => {
      const ws = new WebSocket(`ws://localhost:8080/ws/game`);
      setSocket(ws);

      ws.onopen = () => {
        console.log("connected");
        setSocket(ws);
  
        const joinMessage = {
          type: 'join',
          game_id: Number(gameID),
          player_id: Number(playerID),
          Content: "veeti",
        };
        ws.send(JSON.stringify(joinMessage));
      };
  
      ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setRecievedMessages((prevMessages) => [...prevMessages, data.message || JSON.stringify(data)]);
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

    console.log("joining game with player id", playerID);
    joinGame();
  }, [playerID]);

  const sendMessage = (e: FormEvent) => {
    e.preventDefault();
    
    if (socket && socket.readyState === WebSocket.OPEN && userMessage.trim()) {
      const messageObject = {
        type: 'move',
        game_id: Number(gameID),
        player_id: Number(playerID),
        content: userMessage
      };
      socket.send(JSON.stringify(messageObject));
      setUserMessage('');
    } else {
      console.error("WebSocket is not open or message is empty.");
    }
  };

  return (
    <div>
      <div className="messages">
        {recievedMessages.map((text, index) => (
          <div key={index} className="message">
            {text}
          </div>
        ))}
      </div>
      <form onSubmit={sendMessage} className="message-form">
        <input
          type="text"
          value={userMessage}
          onChange={(e) => setUserMessage(e.target.value)}
          placeholder="Type a message..."
        />
        <button type="submit">Send</button>
      </form>
    </div>
  );
};

export default Game;