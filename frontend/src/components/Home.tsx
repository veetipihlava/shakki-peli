import CreateGameButton from './CreateGameButton.tsx';
import Join from './Join.tsx';

const Home: React.FC = () => {
  return (
    <div>
      <CreateGameButton />
      <Join />
    </div>
  );
};

export default Home;