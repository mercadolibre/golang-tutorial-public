package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	return r
}

//agregar servicio GET /items que obtenga el parametro query como querystring
//si no se recibe parámetro devolver un texto de error, caso contrario devolver el query solicitado
func main() {
	r := setupRouter()
	r.Run(":8080")
}

