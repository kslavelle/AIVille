package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	g "github.com/kslavelle/AIVille/pkg/game"
)

func healthCheck(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"detail": "Ok",
		})
	}
}

func gameInfo(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &GetGameStatsModel{}
		err := c.ShouldBindJSON(requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid request body, expected filed `game_id`",
			})
			return
		}

		userID := env.GetUser(c)
		game, err := g.Load(env.DB, requestBody.GameID, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "invalid request body. expected field `game_id`",
			})
		}

		err = game.UpdateGameTime(env.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
		}
		// env.Log.Info(game.GetElapsedGameTime().String())
		fmt.Println(game.GetElapsedGameTime().String())
	}
}

func createNewGame(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &CreateGameModel{}
		err := c.ShouldBindJSON(requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid request body, expected filed `name`",
			})
			return
		}

		userID := env.GetUser(c)
		err = g.CreateGame(env.DB, userID, requestBody.Name)
		if err != nil {
			env.Log.Errorf("error when creating game: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "Failed to create game.",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"detail": "game created",
		})
	}
}
