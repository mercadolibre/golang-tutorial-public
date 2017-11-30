package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/items", search)
	return r
}

//Hacer api call a items api y completar item info con titulo, precio y ubicacion (lat,long)
func main() {
	r := setupRouter()
	r.Run(":8080")
}

