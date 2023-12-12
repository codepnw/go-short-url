package main

import (
	"github.com/codepnw/go-short-url/internal/database"
	"github.com/codepnw/go-short-url/internal/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDB()
}

func main() {
	router.ClientRoutes()
}