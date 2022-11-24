package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lordnr/jwt-authentication/config"
	"github.com/lordnr/jwt-authentication/database"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
