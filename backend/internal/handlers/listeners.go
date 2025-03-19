package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
)

type JoinMessage struct {
	Name string `json:"content"`
}

// handleJoinRequest processes a join request from a player
func handleJoinRequest(ws *websocket.Conn, request ChessMessage) error {
	player := games.Player{
		Name:       request.Content,
		ID:         request.PlayerID,
		Connection: ws,
	}

	err := games.GameManager.TryAddPlayerToGame(request.GameID, player)
	if err != nil {
		return err
	}

	message := JoinMessage{
		Name: player.Name,
	}

	players, err := games.GameManager.GetPlayers(request.GameID)
	if err != nil {
		log.Printf("Could not read players: %v", err)
	}

	utilities.SendMessageToAllPlayers(players, request.GameID, message)

	return nil
}

type ValidationMessage struct {
	Move             string                 `json:"move"`
	ValidationResult chess.ValidationResult `json:"validation_result"`
}

// handleMoveRequest processes a move request from a player
func handleMoveRequest(request ChessMessage) error {
	/* game, err := utilities.ReadChessGame(db, request.GameID)
	if err != nil {
		return err
	}
	if game == nil {
		return errors.New("game is null")
	} */

	/* color := true // TODO

	validationResult, piecesToUpdate := chess.ValidateMove(*game, request.Content, color)
	for _, pieceUpdate := range piecesToUpdate {
		if pieceUpdate.DeletePiece {
			err = db.DeletePiece(pieceUpdate.Piece.ID)
			if err != nil {
				return errors.New("failed to delete chess piece")
			}
		} else {
			err = db.UpdatePiece(pieceUpdate.Piece)
			if err != nil {
				return errors.New("failed to delete chess piece")
			}
		}
	} */

	message := ValidationMessage{
		Move:             request.Content,
		ValidationResult: chess.ValidationResult{},
	}

	players, err := games.GameManager.GetPlayers(request.GameID)
	if err != nil {
		log.Printf("Could not read players: %v", err)
	}

	utilities.SendMessageToAllPlayers(players, request.GameID, message)

	return nil
}
