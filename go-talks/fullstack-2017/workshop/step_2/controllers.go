package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func search(c *gin.Context) {
	query := c.Query("query")
	if len(query)==0 {
		c.String(http.StatusBadRequest, "'query' param should be send")
		return
	}
	c.String(http.StatusOK,query)
}






