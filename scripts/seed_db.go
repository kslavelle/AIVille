package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func createGameTable(c *pgx.Conn) {
	query := `
		CREATE TABLE IF NOT EXISTS game (
			id serial PRIMARY KEY,
			name text NOT NULL,
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

func createResourceTypes(c *pgx.Conn) {
	query := `
		CREATE TABLE IF NOT EXISTS resource_types (
			id serial PRIMARY KEY,
			name text UNIQUE NOT NULL
		);
	`

	_, err := c.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: resource_types: %v\n", err)
		os.Exit(1)
	}

	query2 := `INSERT INTO resource_types (name) VALUES ('energy') ON CONFLICT DO NOTHING`
	_, err = c.Exec(context.Background(), query2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to populate table: resource_types: %v\n", err)
		os.Exit(1)
	}
}

func createResources(c *pgx.Conn) {
	query := `
		CREATE TABLE IF NOT EXISTS resources (
			id serial PRIMARY KEY,
			type serial NOT NULL REFERENCES resource_types(id),
			name text UNIQUE NOT NULL,
			properties json NOT NULL
		);
	`

	_, err := c.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: resources: %v\n", err)
		os.Exit(1)
	}

	query2 := `INSERT INTO resources (type, name, properties) VALUES (
		1,
		'Coal Power Station',
		'{"cost": 1000000, "co2/hr": 100, "workers": 2}'
	) ON CONFLICT DO NOTHING`
	_, err = c.Exec(context.Background(), query2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to populate table: resources: %v\n", err)
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

	createResourceTypes(conn)
	createResources(conn)
}
