package games

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

var GameManager = GameConnectionsManager{
	Games: make(map[int64]*Game),
}

type Player struct {
	Name       string
	ID         int64
	Connection *websocket.Conn
}

type Game struct {
	ID      int64
	Players []Player
}

type GameConnectionsManager struct {
	DB    *database.DatabaseService
	Games map[int64]*Game
	mutex sync.RWMutex
}

func (gm *GameConnectionsManager) GetGame(gameID int64) (*Game, error) {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()

	game, exists := gm.Games[gameID]
	if !exists {
		return nil, errors.New("game does not exist")
	}

	return game, nil
}

func (gm *GameConnectionsManager) CreateGame(gameID int64) {
	gm.mutex.Lock()
	gm.Games[gameID] = &Game{
		ID:      gameID,
		Players: []Player{},
	}
	gm.mutex.Unlock()
}

func (gm *GameConnectionsManager) TryAddPlayerToGame(gameID int64, player Player) error {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	game, exists := gm.Games[gameID]
	if !exists {
		return errors.New("game does not exist")
	}

	for _, gamePlayer := range game.Players {
		if gamePlayer.ID == player.ID {
			return errors.New("player already joined")
		}
	}

	game.Players = append(game.Players, player)

	return nil
}

func (gm *GameConnectionsManager) GetPlayers(gameID int64) ([]Player, error) {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()

	game, exists := gm.Games[gameID]
	if !exists {
		return nil, errors.New("game does not exist")
	}

	return game.Players, nil
}

/*
func (gm *GameConnectionsManager) EndGame(gameID int64) {
	gm.mutex.Lock()
	game := gm.Games[gameID]
	for _, player := range game.Players {
		if player.Connection != nil {
			player.Connection.Close()
		}
	}
	delete(gm.Games, gameID)
	gm.mutex.Unlock()
} */
