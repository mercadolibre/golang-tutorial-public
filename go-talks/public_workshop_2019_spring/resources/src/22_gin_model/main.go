package main

import "github.com/gin-gonic/gin"
import "net/http"

//START OMIT

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Example curl -X POST -d '{"user": "gus", "password": "123"}' localhost:8080/login
func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var loginInfo Login
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Missing some fields"})
			return
		}

		if loginInfo.User == "gus" && loginInfo.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	router.Run(":8080")
}

//END OMIT
