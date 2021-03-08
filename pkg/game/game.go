package game

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Game represents a game that a user has created
type Game struct {
	ID            int
	name          string
	paused        bool
	owner         int
	lastOperation time.Time
	gameState     State
}

// State represents the specific state of a game
type State struct {
	ID      int
	GameID  int
	Workers int
	CO2     int
	Power   int
	Money   int
	Land    int
	Water   int
}

// CreateGame inserts a new database with the defaults into the DB
func CreateGame(c *pgxpool.Pool, user int, name string) error {
	gameID, err := dbCreateGame(c, name, user)
	if err != nil {
		return err
	}

	// create game state with gameID
	err = dbCreateGameState(c, gameID)
	return err
}

// Load fully loads the specified game into memory
func Load(c *pgxpool.Pool, gameID, userID int) (*Game, error) {
	g := Game{}
	err := dbLoadGame(c, gameID, userID, &g)
	if err != nil {
		return &g, err
	}

	err = dbLoadState(c, gameID, &g)
	if err != nil {
		return &g, err
	}

	return &g, err
}

func (g *Game) GetElapsedGameTime() time.Duration {
	return time.Now().Sub(g.lastOperation) * time.Duration(GameTimeMultiplier)
}

func (g *Game) UpdateGameTime(c *pgxpool.Pool) error {
	err := dbUpdateGameTime(c, g.ID)
	return err
}
