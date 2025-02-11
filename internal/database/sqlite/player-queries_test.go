package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func getTestPlayer(name string, color bool) models.Player {
	return models.Player{
		Name:  name,
		Color: color,
	}
}

func TestCreatePlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	testPlayer := getTestPlayer("Alice", true)
	playerID, err := db.CreatePlayer(testPlayer.Name, testPlayer.Color)
	require.NoError(t, err)
	require.NotZero(t, playerID)
}

func TestReadPlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	testPlayer := getTestPlayer("Alice", true)
	playerID, err := db.CreatePlayer(testPlayer.Name, testPlayer.Color)
	require.NoError(t, err)

	player, err := db.ReadPlayer(playerID)
	require.NoError(t, err)
	require.NotNil(t, player)
	require.Equal(t, testPlayer.Name, player.Name)
	require.Equal(t, testPlayer.Color, player.Color)
}
