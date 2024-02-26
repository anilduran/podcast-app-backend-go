package main

import (
	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
	}

	err = db.InitializeDB()

	if err != nil {
		panic("Could not connect to the database!")
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")

}
