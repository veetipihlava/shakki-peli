import CreateGameButton from './CreateGameButton.tsx';
import Join from './Join.tsx';

const Home: React.FC = () => {
  return (
    <div className="home">
      <CreateGameButton />
      <Join />
    </div>
  );
};

export default Home;