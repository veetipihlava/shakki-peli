import { useNavigate } from 'react-router-dom';

const CreateGameButton = () => {
  const navigate = useNavigate();

  const handleCreateGame = async () => {
    try {
      const response = await fetch('/game', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        console.error('failed to create game:');
        return;
      }

      const data = await response.json();
      const gameID = data.game_id;

      navigate(`/game/${gameID}`);
    } catch (error) {
      console.error('error:', error);
    }
  };

  return (
    <button onClick={handleCreateGame}>
      Create New Game
    </button>
  );
};

export default CreateGameButton;
