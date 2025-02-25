package chess

import (
	"github.com/veetipihlava/shakki-peli/internal/models"
)

// TODO: Castling, en passant, pawn promotion, checkmate

// GetConsumedPiece takes a move string, extracts the destination position, and returns the piece at that position if one exists.
func GetConsumedPiece(move string, pieces []models.Piece) *models.Piece {
	_, _, toFile, toRank, _ := parseMoveFromString(move)
	if toFile == -1 || toRank == -1 {
		return nil
	}

	return getPiece(toFile, toRank, pieces)
}

// ValidateMove validates whether a given move is applicable given the game state.
func ValidateMove(pieces []models.Piece, move string, color bool) (models.ValidationResult, []models.PieceUpdate) {
	var validationResult models.ValidationResult
	fromFile, fromRank, toFile, toRank, pieceName := parseMoveFromString(move)
	if fromFile == -1 || fromRank == -1 || toFile == -1 || toRank == -1 {
		validationResult.IsValidMove = false
		return validationResult, nil
	}

	// Does the position contain a piece?
	piece := getPiece(fromFile, fromRank, pieces)
	if piece == nil {
		validationResult.IsValidMove = false
		return validationResult, nil
	}

	// Is the piece in the position the correct one?
	if piece.Name != pieceName || piece.Color != color {
		validationResult.IsValidMove = false
		return validationResult, nil
	}

	validationResult.IsValidMove = isValidMove(piece, toFile, toRank, pieces)
	// Check that the move is valid
	return validationResult, nil
}

func isValidMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	switch piece.Name {
	case "K":
		return isValidKingMove(piece, toFile, toRank, pieces)
	case "Q":
		return isValidRookMove(piece, toFile, toRank, pieces) || isValidBishopMove(piece, toFile, toRank, pieces)
	case "R":
		return isValidRookMove(piece, toFile, toRank, pieces)
	case "B":
		return isValidBishopMove(piece, toFile, toRank, pieces)
	case "N":
		return isValidKnightMove(piece, toFile, toRank, pieces)
	case "P":
		return isValidPawnMove(piece, toFile, toRank, pieces)
	default:
		return false
	}
}

// The King can move 1 square to any direction
func isValidKingMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank

	if abs(toFile-fromFile) <= 1 && abs(toRank-fromRank) <= 1 {
		return positionNotOccupiedByFriendly(piece, toFile, toRank, pieces)
	}

	return false
}

// Validate the Rook piece consists of checking that only either file or rank change, not both.
// Also need to check Rook does not jump over any pieces, friend or foe
func isValidRookMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank

	if fromFile != toFile && fromRank != toRank {
		return false
	}

	fileStep, rankStep := 0, 0
	if fromFile < toFile {
		fileStep = 1
	} else if fromFile > toFile {
		fileStep = -1
	}

	if fromRank < toRank {
		rankStep = 1
	} else if fromRank > toRank {
		rankStep = -1
	}

	file, rank := fromFile+fileStep, fromRank+rankStep
	for file != toFile || rank != toRank {
		if getPiece(file, rank, pieces) != nil {
			return false
		}
		file += fileStep
		rank += rankStep
	}

	return positionNotOccupiedByFriendly(piece, toFile, toRank, pieces)
}

// The Bishop can move diagonally, but can't jump over pieces
func isValidBishopMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank

	if abs(toFile-fromFile) != abs(toRank-fromRank) {
		return false
	}

	fileStep := 1
	if toFile < fromFile {
		fileStep = -1
	}

	rankStep := 1
	if toRank < fromRank {
		rankStep = -1
	}

	file, rank := fromFile+fileStep, fromRank+rankStep
	for file != toFile || rank != toRank {
		if getPiece(file, rank, pieces) != nil {
			return false
		}
		file += fileStep
		rank += rankStep
	}

	return positionNotOccupiedByFriendly(piece, toFile, toRank, pieces)
}

// The Knight moves according to the offsets, but can't move to a position occupied by friendly.
func isValidKnightMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank

	knightMoves := []struct{ fileDiff, rankDiff int }{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}

	for _, move := range knightMoves {
		if toFile == fromFile+move.fileDiff && toRank == fromRank+move.rankDiff {
			return positionNotOccupiedByFriendly(piece, toFile, toRank, pieces)
		}
	}

	return false
}

// The Pawn generally moves 1 step, but at start can move 2, and can move diagonally to consume another piece
func isValidPawnMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank
	pawnDirection := 1
	if !piece.Color {
		pawnDirection = -1
	}

	if fromFile == toFile {
		if toRank == fromRank+pawnDirection && getPiece(toFile, toRank, pieces) == nil {
			return true
		}

		startingRank := 2
		if !piece.Color {
			startingRank = 7
		}

		if fromRank == startingRank && toRank == fromRank+(2*pawnDirection) {
			if getPiece(toFile, fromRank+pawnDirection, pieces) == nil &&
				getPiece(toFile, toRank, pieces) == nil {
				return true
			}
		}
	}

	if abs(toFile-fromFile) == 1 && toRank == fromRank+pawnDirection {
		targetPiece := getPiece(toFile, toRank, pieces)
		if targetPiece != nil && targetPiece.Color != piece.Color {
			return true
		}
	}

	return false
}

// Parses a move and returns the corresponding coordinates in range 1-8
func parseMoveFromString(move string) (fromFile, fromRank, toFile, toRank int, pieceName string) {
	if len(move) != 5 {
		return -1, -1, -1, -1, ""
	}

	pieceName = string(move[0])
	fromFile = int(move[1]-'a') + 1
	fromRank = int(move[2] - '1')
	toFile = int(move[3]-'a') + 1
	toRank = int(move[4] - '1')

	if fromFile < 1 || fromFile > 8 || toFile < 1 || toFile > 8 ||
		fromRank < 1 || fromRank > 8 || toRank < 1 || toRank > 8 {
		return -1, -1, -1, -1, ""
	}
	return fromFile, fromRank, toFile, toRank, pieceName
}

// Returns the Piece at the given position, or nil if not found.
func getPiece(file, rank int, pieces []models.Piece) *models.Piece {
	for i := range pieces {
		if pieces[i].File == file && pieces[i].Rank == rank {
			return &pieces[i]
		}
	}
	return nil
}

// A function that checks whether a position is not occupied by a Piece with the same color.
func positionNotOccupiedByFriendly(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	destPiece := getPiece(toFile, toRank, pieces)
	if destPiece != nil && destPiece.Color == piece.Color {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
