package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {

	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/me", func(c *gin.Context) {
		c.String(http.StatusOK, "I'm")
	})

	router.Run(":8080")
}