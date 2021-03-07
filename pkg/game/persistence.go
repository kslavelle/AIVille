package game

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func dbCreateGame(c *pgxpool.Pool, name string, user int) (int, error) {
	query := `
		INSERT INTO
			game (name, owner, paused, last_call)
		VALUES
			($1, $2, $3, $4)
		RETURNING id;
	`

	gameID := 0
	err := c.QueryRow(context.Background(), query, name, user, false, time.Now()).Scan(
		&gameID,
	)
	return gameID, err
}

func dbCreateGameState(c *pgxpool.Pool, gameid int) error {
	query := `
		INSERT INTO
			gamestate (gameid, workers, co2, power, money, land, water)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		`
	_, err := c.Exec(context.Background(), query, gameid, 2, 0, 0, 5, 10, 10)
	return err

}
