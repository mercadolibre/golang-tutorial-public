package main

import "github.com/gin-gonic/gin"

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached
	r := gin.Default()

	// GET is a shortcut for router.Handle("GET", path, handle)
	r.GET("/ping", func(c *gin.Context) {
		// JSON serializes the given struct as JSON into the response body.
		// It also sets the Content-Type as "application/json".
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
