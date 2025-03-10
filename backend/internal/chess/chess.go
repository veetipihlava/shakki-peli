package chess

import (
	"github.com/veetipihlava/shakki-peli/internal/models"
)

// TODO: Castling, en passant, pawn promotion, checkmate

// ValidateMove validates whether a given move is applicable given the game state.
func ValidateMove(pieces []models.Piece, move string, color bool) (models.ValidationResult, []models.PieceUpdate) {
	var validationResult models.ValidationResult

	// 1. Get the Piece if move notation is correct and that Piece belongs to the player.
	piece, toFile, toRank := getPieceIfValid(move, pieces, color)
	if piece == nil {
		validationResult.IsValidMove = false
		return validationResult, nil
	}

	// 2. Check if the move is valid => Piece moves correctly and the destination does not contain a friendly Piece.
	validationResult.IsValidMove = isValidMove(piece, toFile, toRank, pieces)
	if !validationResult.IsValidMove {
		return validationResult, nil
	}

	var updates []models.PieceUpdate

	// 3. The move is valid, so check if we consumed an enemy Piece.
	consumedUpdate := GetConsumedPiece(toFile, toRank, pieces)
	if consumedUpdate.Piece.ID != 0 {
		updates = append(updates, consumedUpdate)
	}
	// 4. Update the position of the moved Piece
	updatedPiece := GetUpdatedPiece(toFile, toRank, piece)
	updates = append(updates, updatedPiece)

	return validationResult, updates
}

// This function validates that the move string is correct, the piece exists and belongs to the player. Returns the Piece if piece, player and move format is valid.
func getPieceIfValid(move string, pieces []models.Piece, color bool) (*models.Piece, int, int) {

	// 1. Is the move notation a valid format?
	fromFile, fromRank, toFile, toRank, pieceName := parseMoveFromString(move)
	if fromFile == -1 || fromRank == -1 || toFile == -1 || toRank == -1 {
		return nil, -1, -1
	}

	// 2. Does the position contain the piece from the notation (same type and color)?
	piece := getPiece(fromFile, fromRank, pieces)
	if piece == nil || piece.Name != pieceName || piece.Color != color {
		return nil, -1, -1
	}
	// 3. If yes, return the piece and it's next position for validation
	return piece, toFile, toRank
}

// This function either returns a PieceUpdate or an empty PieceUpdate. It assumes that the position is either empty or contains a Piece from the other color
func GetConsumedPiece(toFile int, toRank int, pieces []models.Piece) models.PieceUpdate {
	piece := getPiece(toFile, toRank, pieces)
	if piece != nil {
		return models.PieceUpdate{
			DeletePiece: true,
			Piece:       *piece,
		}
	}
	return models.PieceUpdate{}
}

// This function returns a PieceUpdate for the given piece.
func GetUpdatedPiece(toFile int, toRank int, piece *models.Piece) models.PieceUpdate {
	updatedPiece := *piece
	updatedPiece.File = toFile
	updatedPiece.Rank = toRank
	return models.PieceUpdate{
		DeletePiece: false,
		Piece:       updatedPiece,
	}
}

// This function routes which validation function is applied to the Piece.
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

// Parses a move notation and returns the corresponding coordinates in range 1-8
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

// Returns a Piece at the given position, or nil.
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
