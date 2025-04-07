import { useState, FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';

const Join = () => {
  const navigate = useNavigate();
  const [gameID, setGameID] = useState<string>("");

  const joinGame = async (e: FormEvent) => {
    e.preventDefault();

    sessionStorage.setItem('gameID', gameID);

    await JoinGame(gameID);
    navigate(`/game`);
  };
  
  return (
    <form onSubmit={joinGame} className="message-form">
        <input
          type="text"
          value={gameID}
          onChange={(e) => setGameID(e.target.value)}
          placeholder="Game ID"
        />
        <button type="submit">Join</button>
    </form>
  );
};

export const JoinGame = async (gameID: string) => {
  const response = await fetch(`/game/${gameID}/join`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error("Failed to join game");
  }

  const data = await response.json();
  const playerID = data.player_id;

  sessionStorage.setItem('playerID', playerID);
};

export default Join;
