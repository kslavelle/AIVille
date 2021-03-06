package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kslavelle/AIVille/pkg/game"
)

func healthCheck(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"detail": "Ok",
		})
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

		hardCodedUser := 0
		err = game.CreateGame(env.DB, hardCodedUser, requestBody.Name)
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
