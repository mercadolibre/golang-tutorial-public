package main

import (
	"github.com/gin-gonic/gin"
)

// agregar ping
func main() {
	r := gin.Default()
	r.Run(":8080")
}
