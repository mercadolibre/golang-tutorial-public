package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	// curl 'localhost:8080/welcome?firstname=Jane&lastname=Doe'
	router.GET("/welcome", func(c *gin.Context) {

		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s \n", firstname, lastname)

	})

	router.Run(":8080")
}
