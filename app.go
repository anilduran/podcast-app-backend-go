package main

import (
	"context"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/routes"
	"example.com/podcast-app-go/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load SDK configs!")
	}

	client := s3.NewFromConfig(cfg)

	utils.InitializePresigner(client)

	err = db.InitializeDB()

	if err != nil {
		panic("Could not connect to the database!")
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")

}
