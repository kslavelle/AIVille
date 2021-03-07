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

	energy := `INSERT INTO resource_types (name) VALUES ('energy') ON CONFLICT DO NOTHING`
	human := `INSERT INTO resource_types (name) VALUES ('human') ON CONFLICT DO NOTHING`
	food := `INSERT INTO resource_types (name) VALUES ('food') ON CONFLICT DO NOTHING`
	water := `INSERT INTO resource_types (name) VALUES ('water') ON CONFLICT DO NOTHING`
	co2 := `INSERT INTO resource_types (name) VALUES ('co2') ON CONFLICT DO NOTHING`

	queries := []string{
		energy, human, food, water, co2,
	}
	for _, query := range queries {
		_, err = c.Exec(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to populate table: resource_types: %v\n", err)
			os.Exit(1)
		}
	}
}

func createResources(c *pgx.Conn) {
	query := `
		CREATE TABLE IF NOT EXISTS resources_actors (
			id serial PRIMARY KEY,
			type serial NOT NULL REFERENCES resource_types(id),
			name text UNIQUE NOT NULL,
			inputs json NOT NULL,
			outputs json NOT NULL,
			constraints json NOT NULL
		);
	`

	_, err := c.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: resources: %v\n", err)
		os.Exit(1)
	}

	coalStation := `INSERT INTO 
		resources_actors (type, name, inputs, outputs, constraints)
		VALUES (
			1, 'Coal Power Station',
			'{"workers": 2}',
			'{"co2": 100, "power": 100}',
			'{"cost": 1000000, "land": 1}'
		) ON CONFLICT DO NOTHING;
		`

	queries := []string{coalStation}

	for _, query := range queries {
		_, err = c.Exec(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to populate table: resources: %v\n", err)
			os.Exit(1)
		}
	}
}

func dropAllTables(c *pgx.Conn) {
	query := `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;

		GRANT ALL ON SCHEMA public TO postgres;
		GRANT ALL ON SCHEMA public TO public;
	`
	_, err := c.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to drop all tables: %v\n", err)
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

	dropAllTables(conn)

	createGameTable(conn)

	createResourceTypes(conn)
	createResources(conn)
}
