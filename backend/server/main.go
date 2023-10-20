package main

import (
	"log"
	"net/http"

	"github.com/Yash-271120/theatre-grpc/backend/server/response"
	router "github.com/Yash-271120/theatre-grpc/backend/server/router"
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

func initEnv() {
	if err := env.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Environment variables loaded successfully")
}

func main() {
	server := router.NewRouter()
	initEnv()
	server.GET("/ping", func(c *gin.Context) {
		response.SuccessResponse(c, http.StatusOK, "Pong")
	})
	server.Start()
}
