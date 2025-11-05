package main

import (
	"log"
	"shorten-service/internal/config"
	"shorten-service/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	//"shorten-service/internal/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()

	app := fiber.New()
	db.Connect(cfg.DBUrl)
	db.ConnectRedis()
	//app.post("/shorten", handler.ShortenURL)

	log.Println("Server Started")

	if err := app.Listen(":8001"); err != nil {
		log.Fatal("Failed to start server", err)
	}

}
