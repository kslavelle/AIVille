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
}

// CreateGame inserts a new database with the defaults into the DB
func CreateGame(c *pgxpool.Pool, user int, name string) error {
	return dbCreateGame(c, name, user)
}
