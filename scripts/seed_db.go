package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func createGameTable(c *pgx.Conn) {
	query := `
		CREATE TABLE game (
			id serial PRIMARY KEY,
			owner serial NOT NULL,
			paused boolean NOT NULL,
			last_call timestamp
		);
	`

	_, err := c.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: game: %v\n", err)
		os.Exit(1)
	}
}

func main() {

	dbURL := "postgres://postgres:postgres@127.0.0.1:5432/aiville"
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	createGameTable(conn)
}
