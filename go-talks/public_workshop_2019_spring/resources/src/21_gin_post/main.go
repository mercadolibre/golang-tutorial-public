package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "fmt"

func main() {

	router := gin.Default()

	//curl -X POST -F 'username=jhon' -F 'password=123' localhost:8080/post
	router.POST("/post", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		fmt.Printf("username: %s; password: %s", username, password)

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("<html><body><h2>Datos Enviados -> Usuario: " + username + " -> Password: "+ password +"</h2></body></html>"))

	})

	router.Run(":8080")

}
