package chess

import "github.com/veetipihlava/shakki-peli/internal/models"

// The King can move 1 square to any direction, if not occupied by friendly.
func isValidKingMove(piece *models.Piece, toFile, toRank int, pieces []models.Piece) bool {
	fromFile, fromRank := piece.File, piece.Rank

	if abs(toFile-fromFile) <= 1 && abs(toRank-fromRank) <= 1 {
		return positionNotOccupiedByFriendly(piece, toFile, toRank, pieces)
	}

	return false
}

// Only either file or rank change, not both. Rook does not jump over any pieces. The destination should not contain a friendly.
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

// The Bishop can move diagonally, but can't jump over pieces. The destination should not contain a friendly.
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

// The Knight moves according to the L-shaped offsets, but can't move to a position occupied by friendly.
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

	// Validate scenario where Pawn moves 1 straight
	pawnDirection := 1
	if !piece.Color {
		pawnDirection = -1
	}

	if fromFile == toFile && toRank == fromRank+pawnDirection {
		// Pawn can't move if there is an enemy
		if getPiece(toFile, toRank, pieces) == nil {
			return true
		}
		return false
	}

	// Validate scenario where Pawn moves 2 straight
	startingRank := 2
	if !piece.Color {
		startingRank = 7
	}

	if fromRank == startingRank && toRank == fromRank+(2*pawnDirection) {
		if getPiece(toFile, fromRank+pawnDirection, pieces) == nil &&
			getPiece(toFile, toRank, pieces) == nil {
			return true
		}
		return false
	}

	// Validate scenario where Pawn moves diagonally to capture
	if abs(toFile-fromFile) == 1 && toRank == fromRank+pawnDirection {
		targetPiece := getPiece(toFile, toRank, pieces)
		if targetPiece != nil && targetPiece.Color != piece.Color {
			return true
		}
		return false
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
