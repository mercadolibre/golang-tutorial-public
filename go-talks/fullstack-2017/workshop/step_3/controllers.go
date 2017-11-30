package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func search(c *gin.Context) {
	query := c.Query("query")
	if len(query)==0 {
		c.JSON(http.StatusBadRequest, apiErr{"'query' param should be send",http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK,gin.H{"query":query})
}

//Custom api error
type apiErr struct {
	ErrorMessage string    `json:"message"`
	ErrorStatus  int       `json:"status"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s; Status: %d", e.ErrorMessage, e.ErrorStatus)
}





