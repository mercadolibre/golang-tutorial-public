package main

import "github.com/gin-gonic/gin"
import "net/http"
//START OMIT
// Binding could be from JSON, XML or form values (foo=bar&boo=baz). 
// Bind will infer binder from Content-Type header

type Login struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func main() {
    router := gin.Default()
    // Example for binding JSON ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var loginInfo Login
        if c.BindJSON(&loginInfo) == nil {
            if loginInfo.User == "gus" && loginInfo.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
    })
    router.Run(":8080")
}
//END OMIT