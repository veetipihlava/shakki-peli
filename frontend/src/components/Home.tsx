import CreateGameButton from './CreateGameButton.tsx';
import Join from './Join.tsx';
import Navbar from './Navbar.tsx';
import './Home.css'

const Home: React.FC = () => {
  return (
    <div>
      <Navbar />
      <div className='options-container'>
        <div className="game-card-container">
          <h2 className="game-card-title">Choose an option to play</h2>
          <div className="game-card-buttons">
            <CreateGameButton />
            <Join />
          </div>
        </div>

      </div>
    </div>
  );
};

export default Home;