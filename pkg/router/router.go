package router

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// CreateAPI set's up all the routing to the http handlers
func CreateAPI() (*gin.Engine, *pgxpool.Pool) {

	// setup the logger
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)

	dbURL := "postgres://postgres:postgres@127.0.0.1:5432/aiville"
	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	env := Env{DB: conn, Log: logrus.Logger{}}

	router := gin.Default()

	router.GET("/health", healthCheck(&env))
	router.POST("/game", createNewGame(&env))

	return router, conn
}
