package main

import (
	"github.com/lordnr/jwt-authentication/config"
	"github.com/lordnr/jwt-authentication/database"
	"github.com/lordnr/jwt-authentication/models"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
