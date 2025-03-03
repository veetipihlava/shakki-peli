package games

import (
	"sync"

	"github.com/gorilla/websocket"
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
	Games map[int64]*Game
	mutex sync.RWMutex
}

func (gm *GameConnectionsManager) GameExists(gameID int64) bool {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()
	_, exists := gm.Games[gameID]
	return exists
}

func (gm *GameConnectionsManager) GameIsFull(gameID int64) bool {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()
	return len(gm.Games[gameID].Players) >= 2
}

func (gm *GameConnectionsManager) CreateGame(gameID int64) {
	gm.mutex.Lock()
	gm.Games[gameID] = &Game{
		ID:      gameID,
		Players: []Player{},
	}
	gm.mutex.Unlock()
}

func (gm *GameConnectionsManager) AddPlayerToGame(gameID int64, player Player) {
	gm.mutex.Lock()
	gm.Games[gameID].Players = append(gm.Games[gameID].Players, player)
	gm.mutex.Unlock()
}

func (gm *GameConnectionsManager) GetPlayers(gameID int64) []Player {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()
	return gm.Games[gameID].Players
}

func (gm *GameConnectionsManager) EndGame(gameID int64) {
	gm.mutex.Lock()
	game := gm.Games[gameID]
	for _, player := range game.Players {
		player.Connection.Close()
	}
	delete(gm.Games, gameID)
	gm.mutex.Unlock()
}
