package router

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// Env that we want to pass down to API calls.
type Env struct {
	DB  *pgxpool.Pool
	Log logrus.Logger
}

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
	env := Env{DB: conn}

	router := gin.Default()

	router.GET("/health", healthCheck(&env))
	router.POST("/game", createNewGame(&env))
	router.POST("/resources", createResource(&env))

	return router, conn
}
