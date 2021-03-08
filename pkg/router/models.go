package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// CreateGameModel is the expected body when creating
// a new game
type CreateGameModel struct {
	Name string `json:"name"`
}

// GetGameStatsModel returns all of the state for the
// specified game
type GetGameStatsModel struct {
	GameID int `json:"game_id"`
}

// Env that we want to pass down to API calls.
type Env struct {
	DB  *pgxpool.Pool
	Log logrus.Logger
}

// GetUser gets the current user
func (e *Env) GetUser(c *gin.Context) int {
	// write a query here to select a user from the database
	// if they're not found we should instead insert one
	// into the database and then return their ID
	// for now it will be faked

	return 0
}
