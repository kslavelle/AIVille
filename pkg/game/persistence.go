package game

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func dbCreateGame(c *pgxpool.Pool, name string, user int) error {
	query := `
		INSERT INTO
			game (name, owner, paused, last_call)
		VALUES
			($1, $2, $3, $4)
	`
	_, err := c.Exec(context.Background(), query, name, user, false, time.Now())
	return err
}
