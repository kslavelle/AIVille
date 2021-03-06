package router

import "github.com/gin-gonic/gin"

// CreateAPI set's up all the routing to the http handlers
func CreateAPI() *gin.Engine {

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"detail": "Ok",
		})
	})

	return router
}
