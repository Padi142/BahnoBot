package main

import (
	"bahno_bot/feature/api/routes"
	"bahno_bot/generic/user"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func apiService(db *mongo.Database) {
	app := fiber.New()

	userRepo := user.NewUserRepository(*db, "users")

	userUseCase := user.NewUserUseCase(userRepo, time.Duration(time.Second*10))

	api := app.Group("/api")
	routes.UserRoter(api, userUseCase)

	defer cancel()

	log.Fatal(app.Listen(":8081"))
}
