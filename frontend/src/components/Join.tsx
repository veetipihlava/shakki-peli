import { useState, FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';

const CreateGameButton = () => {
  const navigate = useNavigate();
  const [gameID, setGameID] = useState<string>("");

  const joinGame = (e: FormEvent) => {
    e.preventDefault();
    navigate(`/game/${gameID}`);
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

export default CreateGameButton;
