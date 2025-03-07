export type pieceColor = "white" | "black"

export interface PieceProps {
  id: number;
  name: string;
  color: pieceColor
}

export interface SquareProps {
  id: string;
  rank: number; // row
  file: string; // col
  isWhiteSquare: boolean;
  piece: (PieceProps | null)
  handlePieceMove(sourceSquare: SquareProps, targetSquare: SquareProps): void;
}