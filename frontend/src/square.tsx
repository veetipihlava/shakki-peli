import React, { memo } from "react";
import { SquareProps } from "./types";
import { ChessPiece } from "./piece";

function Square(props: SquareProps) {
  const { id, rank, file, isWhiteSquare, piece, handlePieceMove } = props;

  const handleDragStart = (e: React.DragEvent) => {
    if (piece) {
      // TODO: limit props sent to id, piece
      e.dataTransfer.setData("sourceSquare", JSON.stringify({ props }));
    }
  };

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    const sourceSquare = JSON.parse(e.dataTransfer.getData("sourceSquare"));
    console.log(sourceSquare.props)
    console.log(props)

    // TODO: limit props sent to id, piece
    handlePieceMove(sourceSquare.props, props);  // Pass the dragged piece and the new square id
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
  };

  function renderSquare() {
    console.log(piece)
    return (
      <div
        draggable={!!piece}  // Only make the square draggable if it has a piece
        className={isWhiteSquare ? "white" : "black"}
        id={id}
        key={id}
        onDragStart={handleDragStart}
        onDrop={handleDrop}
        onDragOver={handleDragOver}
      >
        {piece && <ChessPiece {...piece} />}
        {file + rank}
      </div>
    );
  }

  return renderSquare();
}


export default memo(Square, (prevProps, nextProps) => {
  // Only re-render if any relevant props have changed
  return (
    prevProps.id === nextProps.id &&
    prevProps.rank === nextProps.rank &&
    prevProps.file === nextProps.file &&
    prevProps.isWhiteSquare === nextProps.isWhiteSquare &&
    prevProps.piece === nextProps.piece
  );
});