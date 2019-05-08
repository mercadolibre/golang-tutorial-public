package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", calculate)

	r.Run()
}

func calculate(c *gin.Context) {
	if v := c.Query("value"); v != "" {
		if p, error := strconv.Atoi(v); error == nil {
			count, sum := multiples(p)

			c.JSON(http.StatusOK, gin.H{
				"count": count,
				"sum":   sum,
			})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Invalid arguments..., error: %s", error.Error()),
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid arguments..., must specify a value parameter",
		})
	}
}

func multiples(p int) (int, int) {
	count, sum := 0, 0

	for i := 0; i <= p; i++ {
		if i%3 == 0 || i%5 == 0 {
			count++
			sum += i
		}
	}

	return count, sum
}
