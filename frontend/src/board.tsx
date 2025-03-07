import React, { useState, useEffect } from 'react';
import './board.css'; // Optional: to add your styles
import Square from './square';
import { fenToPieces } from './utils';
import { FEN_START } from "./constants";
import { SquareProps, PieceProps } from './types';

function Chessboard() {
  // Initial board state setup
  const [board, setBoard] = useState<SquareProps[]>([]);

  const generateBoard = (pieces: (PieceProps | null)[]): SquareProps[] => {
    const generatedBoard: SquareProps[] = [];
    for (let rank = 0; rank < 8; rank++) {
      for (let col = 0; col < 8; col++) {
        const isWhiteSquare = (rank % 2 === 0 && col % 2 === 0) || (rank % 2 !== 0 && col % 2 !== 0);
        const file = String.fromCharCode(97 + col);
        const piece = pieces[(8 * rank + col)]
        const squareID = (file + rank).toString();

        generatedBoard.push({
          id: squareID,
          rank,
          file,
          isWhiteSquare,
          piece,
          handlePieceMove: movePiece
        });
      }
    }
    return generatedBoard;
  }

  useEffect(() => {
    const pieces = fenToPieces(FEN_START);
    const generatedBoard = generateBoard(pieces)
    setBoard(generatedBoard);
  }, []);


  const movePiece = (sourceSquare: SquareProps, targetSquare: SquareProps) => {
    console.log(`Moving piece: ${sourceSquare.piece?.name} from ${sourceSquare.id} to ${targetSquare.id}`);

    setBoard((prevBoard) => {
      const updatedBoard = prevBoard.map(square => {
        if (square.id === sourceSquare.id) {
          // Empty the source square
          return { ...square, piece: null };
        }
        if (square.id === targetSquare.id) {
          // Put the dragged piece into the target square
          return { ...square, piece: sourceSquare.piece };
        }
        return square;
      });

      return updatedBoard;
    });
  };

  return (
    <div>
      <div id="chessboard" className="chessboard">
        {board.map((square) => (
          <Square
            key={square.id + square.piece}
            id={square.id}
            rank={square.rank}
            file={square.file}
            isWhiteSquare={square.isWhiteSquare}
            piece={square.piece}
            handlePieceMove={movePiece}
          />
        ))}
      </div>
    </div>
  );
}

export default Chessboard;
