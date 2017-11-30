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

//Hacer apicall a search usando restclient y obteniendo items id que macheen con el query
//Iterar paginas obteniendo hasta un máximo de 100 items
//Devolver lista de items ids o error en caso de que algo falle
func main() {
	r := setupRouter()
	r.Run(":8080")
}

