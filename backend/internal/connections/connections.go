package connections

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

type ConnectionManager struct {
	mutex sync.RWMutex
	Games map[int64][]PlayerConn
}

type PlayerConn struct {
	Player *models.Player
	Conn   *websocket.Conn
}

var ConnManager = ConnectionManager{
	Games: make(map[int64][]PlayerConn),
}

func (cm *ConnectionManager) AddGame(gameID int64) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if _, exists := cm.Games[gameID]; !exists {
		cm.Games[gameID] = []PlayerConn{}
	}
}

func (cm *ConnectionManager) VerifyGameOngoing(gameID int64) ([]PlayerConn, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	_, exists := cm.Games[gameID]
	if !exists {
		return nil, errors.New("no ongoing game with id")
	}

	return nil, nil
}

func (cm *ConnectionManager) RemoveGame(gameID int64) error {
	cm.mutex.RLock()
	conns, exists := cm.Games[gameID]
	cm.mutex.RUnlock()

	if !exists {
		return errors.New("no ongoing game with id")
	}

	for _, playerConn := range conns {
		_ = cm.RemovePlayer(gameID, playerConn)
	}
	return nil
}

func (cm *ConnectionManager) GetPlayerFromConn(conn *websocket.Conn) (int64, PlayerConn, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	for gameID, conns := range cm.Games {
		for _, pc := range conns {
			if pc.Conn == conn {
				return gameID, pc, nil
			}
		}
	}

	return 0, PlayerConn{}, errors.New("connection not found in any game")
}

func (cm *ConnectionManager) TryAddPlayerToGame(gameID int64, playerConn PlayerConn) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return errors.New("no ongoing game with id")
	}

	for _, pc := range conns {
		if pc.Player.UserID == playerConn.Player.UserID {
			return errors.New("player already in game")
		}
	}

	cm.Games[gameID] = append(conns, playerConn)
	return nil
}

func (cm *ConnectionManager) RemovePlayer(gameID int64, playerConn PlayerConn) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return errors.New("no ongoing game with id")
	}

	for i, pc := range conns {
		if pc.Conn == playerConn.Conn {
			pc.Conn.Close()
			cm.Games[gameID] = append(conns[:i], conns[i+1:]...)

			if len(cm.Games[gameID]) == 0 {
				delete(cm.Games, gameID)
			}
			return nil
		}
	}

	return errors.New("player connection not found in game")
}

func (cm *ConnectionManager) GetGameConnections(gameID int64) ([]PlayerConn, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return nil, errors.New("no ongoing game with id")
	}

	return conns, nil
}

func (cm *ConnectionManager) GetPlayerConnection(gameID int64, player models.Player) (*websocket.Conn, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return nil, errors.New("game does not exist")
	}

	for _, pc := range conns {
		if pc.Player.UserID == player.UserID {
			return pc.Conn, nil
		}
	}

	return nil, errors.New("player not found in game")
}
