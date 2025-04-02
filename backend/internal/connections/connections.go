package connections

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

type playerConn struct {
	Player *models.Player
	Conn   *websocket.Conn
}

type ConnectionManager struct {
	mutex sync.RWMutex
	Games map[int64][]playerConn
}

var ConnManager = &ConnectionManager{
	Games: make(map[int64][]playerConn),
}

func (cm *ConnectionManager) AddGame(gameID int64) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if _, exists := cm.Games[gameID]; !exists {
		cm.Games[gameID] = []playerConn{}
	}
}

func (cm *ConnectionManager) VerifyGame(gameID int64) error {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	if _, exists := cm.Games[gameID]; !exists {
		return errors.New("no ongoing game with id")
	}
	return nil
}

func (cm *ConnectionManager) RemoveGame(gameID int64) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return errors.New("no ongoing game with id")
	}

	for _, pc := range conns {
		_ = pc.Conn.Close()
	}
	delete(cm.Games, gameID)
	return nil
}

func (cm *ConnectionManager) AddPlayerConnection(gameID int64, player *models.Player, conn *websocket.Conn) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	pc := playerConn{Player: player, Conn: conn}
	conns, exists := cm.Games[gameID]
	if !exists {
		return errors.New("no ongoing game with id")
	}

	for _, existingPC := range conns {
		if existingPC.Player.UserID == pc.Player.UserID {
			return errors.New("player already in game")
		}
	}

	cm.Games[gameID] = append(conns, pc)
	return nil
}

func (cm *ConnectionManager) RemoveConnection(conn *websocket.Conn) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for gameID, conns := range cm.Games {
		for i, pc := range conns {
			if pc.Conn == conn {
				pc.Conn.Close()
				cm.Games[gameID] = append(conns[:i], conns[i+1:]...)
				if len(cm.Games[gameID]) == 0 {
					delete(cm.Games, gameID)
				}
				return nil
			}
		}
	}
	return errors.New("connection not found in any game")
}

func (cm *ConnectionManager) GetConnectionsInGame(gameID int64) ([]*websocket.Conn, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	conns, exists := cm.Games[gameID]
	if !exists {
		return nil, errors.New("no ongoing game with id")
	}

	wsConns := make([]*websocket.Conn, len(conns))
	for i, pc := range conns {
		wsConns[i] = pc.Conn
	}
	return wsConns, nil
}

func (cm *ConnectionManager) GetPlayerByConnection(conn *websocket.Conn) (*models.Player, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	for _, conns := range cm.Games {
		for _, pc := range conns {
			if pc.Conn == conn {
				return pc.Player, nil
			}
		}
	}
	return nil, errors.New("connection not found")
}
