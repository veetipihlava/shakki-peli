import React, { useState } from 'react';

const Input: React.FC = () => {
  const [moveInput, setMoveInput] = useState<string>('');

  // Handle input change
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMoveInput(event.target.value);
  };

  // Handle form submission
  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    setMoveInput(''); // Clear input after submitting
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={moveInput}
          onChange={handleInputChange}
          placeholder="Enter move (e.g., Pa2a3)"
          required
        />
        <button type="submit">Submit Move</button>
      </form>
    </div>
  );
};

export default Input;
