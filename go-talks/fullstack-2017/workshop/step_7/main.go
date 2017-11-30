package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", ping)
	r.GET("/items", search)
	r.GET("/render", render)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
