// First game validates fen-record
// if not valid, notify of broken record and allow starting new game

// If valid_fen then setupGame():
// -> fenToPieces: (PieceProps | null)[] i.e. 1d-array of pieces or nulls if no piece


import { PieceProps, pieceColor } from "./types";


export function fenToPieces(fen: string): (PieceProps | null)[] {
  const pieces: (PieceProps | null)[] = [];

  const parsedRanks =  fen.split(" ")[0].split("/")
  for (let rankIndex = 0; rankIndex < 8; rankIndex++) {
    const rank = parsedRanks[rankIndex];
    let fileIndex = 0;

    for (let i = 0; i < rank.length; i++) {
      const emptySquaresCount = parseInt(rank[i])

      if (!isNaN(emptySquaresCount)) {
        for (let i = 0; i < emptySquaresCount; i++) {
          pieces.push(null)
          fileIndex++
        }
      } else {
        const pieceText = rank[fileIndex]
        const pieceColor = returnPieceColor(pieceText)
        const piece: PieceProps = {id: (8*rankIndex+fileIndex), name: pieceText, color: pieceColor} 
        pieces.push(piece)
        fileIndex++
      }
    } // file loop ends
    fileIndex = 0
  } // rank loop ends
  return pieces;
};


function returnPieceColor(piece: string): pieceColor {
  if(piece.toLowerCase() === piece){
    return "black"
  }
  return "white"
};


export function validateFenRecord(fen: string) {
  const pieces = "kqrnbpKQRNBP1"
  if(!(typeof(fen) === 'string')) return false

  // take only position info
  fen = fen.split(" ")[0]

  // expand empty squares to 1's
  fen = expandEmptySquares(fen)

  const ranks = fen.split('/');

  // check correct amount of ranks
  if (ranks.length !== 8) return false;

  // check each rank
  for (let i = 0; i < 8; i++) {
    if (ranks[i].length !== 8) return false;
    const rank = ranks[i]
    for (let j = 0; j < 8; j++) {
      const elem = rank[j];
      if(!pieces.includes(elem)) return false 
    }
  }
  return true;
}

function expandEmptySquares(fen: string) {
  return fen
    .replace(/8/g, '11111111')
    .replace(/7/g, '1111111')
    .replace(/6/g, '111111')
    .replace(/5/g, '11111')
    .replace(/4/g, '1111')
    .replace(/3/g, '111')
    .replace(/2/g, '11');
}