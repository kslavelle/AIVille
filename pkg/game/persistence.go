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
			stateofgame (gameid, workers, co2, power, money, land, water)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		`
	_, err := c.Exec(context.Background(), query, gameid, 2, 0, 0, 5, 10, 10)
	return err
}

func dbUpdateGameTime(c *pgxpool.Pool, gameID int) error {
	query := "UPDATE game SET last_call = $1 WHERE id=$2"
	_, err := c.Exec(context.Background(), query, time.Now(), gameID)
	return err
}

func dbLoadGame(c *pgxpool.Pool, gameID, userID int, game *Game) error {
	query := `
		SELECT
			id, name, owner, paused, last_call
		FROM
			game
		WHERE
			id=$1 AND owner=$2
	`

	err := c.QueryRow(context.Background(), query, gameID, userID).Scan(
		&game.ID, &game.name, &game.owner, &game.paused, &game.lastOperation,
	)
	return err
}

func dbLoadState(c *pgxpool.Pool, gameID int, game *Game) error {
	query := `
		SELECT
			id, gameid, workers, co2, power, money, land, water
		FROM
			stateofgame
		WHERE
			gameid=$1
	`

	err := c.QueryRow(context.Background(), query, gameID).Scan(
		&game.gameState.ID, &game.gameState.GameID, &game.gameState.Workers,
		&game.gameState.CO2, &game.gameState.Power, &game.gameState.Money,
		&game.gameState.Land, &game.gameState.Water,
	)
	return err
}
