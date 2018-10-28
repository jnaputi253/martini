package main

import (
	"net/http"

	"martini/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/greet", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Ready for V1 endpoints",
			})
		})

		v1.POST("/new", func(c *gin.Context) {
			var person models.Person

			if err := c.BindJSON(&person); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"user": person,
			})
		})
	}

	router.Run()
}
