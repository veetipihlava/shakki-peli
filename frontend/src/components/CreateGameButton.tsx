import { useNavigate } from 'react-router-dom';
import { JoinGame } from './Join';

const CreateGameButton = () => {
  const navigate = useNavigate();

  const handleCreateGame = async () => {
    const response = await fetch('/game', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      console.error('Failed to create game');
      return;
    }

    const data = await response.json();
    const gameID = data.game_id;

    sessionStorage.setItem('gameID', gameID);
    
    await JoinGame(gameID);

    navigate(`/game`);
  };

  return (
    <button onClick={handleCreateGame}>
      Create New Game
    </button>
  );
};

export default CreateGameButton;
