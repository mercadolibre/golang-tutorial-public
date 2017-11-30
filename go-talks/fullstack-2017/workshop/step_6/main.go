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

//integrar con google maps para mostrar ubicaciones de los items
func main() {
	r := setupRouter()
	r.Run(":8080")
}
