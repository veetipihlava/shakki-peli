import { useState, FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import "./Home.css"
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
    <form onSubmit={joinGame} className="game-join-form">
      <input
        type="text"
        className='game-join-input'
        value={gameID}
        onChange={(e) => setGameID(e.target.value)}
        placeholder="Game ID"
      />
      <button
        type="submit"
        className='game-join-btn'
      >
        Join
      </button>
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
  const userID = data.user_id;

  sessionStorage.setItem('userID', userID);
};

export default Join;
