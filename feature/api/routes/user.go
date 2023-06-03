package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, useCase user.UseCase) {
	app.Get("/users", handlers.GetUsers(useCase))
	app.Get("/users/:id", handlers.GetUser(useCase))
	app.Put("/users", handlers.UpdateUser(useCase))
	app.Get("/users/:id/records", handlers.GetUserRecords(useCase))
	app.Get("/users/:id/records/last", handlers.GetLastUserRecord(useCase))
	app.Get("/user_discord", handlers.GetUserDiscord(useCase))
	//app.Post("/user", handlers.CreateUSer(service))
}
