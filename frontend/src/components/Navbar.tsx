//import { useState, FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import './Navbar.css'
const Navbar = () => {
  const navigate = useNavigate();
  //const [gameID, setGameID] = useState<string>("");

  const handleNavigation = async () => {
    const response = await fetch('/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      console.error('Failed to navigate to homepage');
      return;
    }

    //const data = await response.json();

    sessionStorage.setItem('gameID', '');

    navigate(`/`);
  };

  return (
    <nav className='navbar'>
      <a href='/' onClick={handleNavigation}>
        Shakk.io
      </a>
    </nav>
  );
};

export default Navbar;
