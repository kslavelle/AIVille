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

func dbUpdateResource(c *pgxpool.Pool, gameID int, resource string) error {
	availableMoneyQuery := `
		SELECT money
		FROM stateofgame
		WHERE gameid = $1
		`
	availableMoney := 0
	err := c.QueryRow(context.Background(), availableMoneyQuery, resource).Scan(&availableMoney)
	if err != nil {
		return err
	}

	resourceCostQuery := `
		SELECT constraints -> 'resource_cost' 
		AS money
		outputs -> 'co2'
		AS co2
		FROM resources_actors
		WHERE name = $1
		RETURNING resource_cost, co2
		`
	resourceCost := 0
	// co2 := 0
	err = c.QueryRow(context.Background(), resourceCostQuery, resource).Scan(&resourceCost)
	if err != nil {
		return err
	}

	postPurchaseMoney := availableMoney - resourceCost
	if postPurchaseMoney < 0 {
		return err
	}

	updateQuery := `
		UPDATE stateofgame
		SET $1 = $1 + 1
		SET money = $2
		SET workers = workers + 
		WHERE gameid = $3
		`
	_, err = c.Exec(context.Background(), updateQuery, resource, postPurchaseMoney, gameID)

	if err != nil {
		return err
	}

}
