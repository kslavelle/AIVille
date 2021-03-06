package router

import "github.com/gin-gonic/gin"

func healthCheck(env *Env, c *gin.Context) {
	c.JSON(200, gin.H{
		"detail": "Ok",
	})
}
