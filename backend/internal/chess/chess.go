package chess

import (
	"github.com/veetipihlava/shakki-peli/internal/models"
)

// En passant not possible, we need to know history...
// Castling requires a check neither King or Castle have moved.
// Checkmate requires more calculations

// ValidateMove validates whether a given move is applicable given the game state.
func ValidateMove(pieces []models.Piece, move string, color bool) (models.ValidationResult, []models.PieceUpdate) {
	var validationResult models.ValidationResult

	// 1. Get the Piece if move notation is correct and that Piece belongs to the player.
	piece, toFile, toRank := getPieceIfNotationValid(move, pieces, color)
	if piece == nil {
		validationResult.IsValidMove = false
		return validationResult, nil
	}

	// 2. Check if the move is valid => Piece moves correctly and the destination does not contain a friendly Piece.
	validationResult.IsValidMove = isValidMove(piece, toFile, toRank, pieces)
	if !validationResult.IsValidMove {
		return validationResult, nil
	}

	validationResult.Move = move

	var updates []models.PieceUpdate

	// 3. The move is valid, so check if we consumed an enemy Piece.
	consumedUpdate := GetConsumedPiece(toFile, toRank, pieces)
	if consumedUpdate.Piece.ID != 0 {
		updates = append(updates, consumedUpdate)
	}
	// 4. Update the position of the moved Piece
	updatedPiece := GetUpdatedPiece(toFile, toRank, piece)
	updates = append(updates, updatedPiece)

	// 5. If king was consumed, gameover
	if consumedUpdate.Piece.Name == "K" {
		validationResult.GameOver.KingConsumed = true
	}

	//6. Does atleast one Piece have King in check?
	if kingInCheck(pieces, !color) {
		validationResult.KingInCheck = true
	}

	//7. TODO: If atleast 1, then check for Checkmate situation

	return validationResult, updates
}

// This function validates that the move string is correct, the piece exists and belongs to the player. Returns the Piece if piece, player and move format is valid.
func getPieceIfNotationValid(move string, pieces []models.Piece, color bool) (*models.Piece, int, int) {

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
		DeletePiece:    false,
		Piece:          updatedPiece,
		TransformPiece: promoted(updatedPiece),
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

// Return true if any Piece of opposing color can consume this king
func kingInCheck(pieces []models.Piece, color bool) bool {
	var king *models.Piece
	for i := range pieces {
		if pieces[i].Name == "K" && pieces[i].Color == color {
			king = &pieces[i]
			break
		}
	}

	// Should not happen
	if king == nil {
		return false
	}

	for i := range pieces {
		if pieces[i].Color != color {
			if isValidMove(&pieces[i], king.File, king.Rank, pieces) {
				return true
			}
		}
	}

	return false
}

// TODO: is checkmate
func isCheckmate(pieces []models.Piece, color bool, kingInCheck bool) bool {
	if !kingInCheck {
		return false
	}

	return false
}

func promoted(piece models.Piece) string {
	if piece.Name != "P" {
		return ""
	}

	if !piece.Color && piece.Rank == 1 {
		return "Q"
	}

	if piece.Color && piece.Rank == 8 {
		return "Q"
	}

	return ""
}
