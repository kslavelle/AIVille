package router

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Env struct {
	DB *pgxpool.Pool
}

// CreateAPI set's up all the routing to the http handlers
func CreateAPI() *gin.Engine {

	dbURL := "postgres://postgres:postgres@127.0.0.1:5432/aiville"
	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	env := Env{DB: conn}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		healthCheck(&env, c)
	})

	return router
}
