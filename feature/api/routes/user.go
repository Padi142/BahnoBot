package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, useCase user.UseCase) {
	app.Get("/user", handlers.GetUserByDiscordId(useCase))
	app.Put("/user", handlers.UpdateUser(useCase))
	//app.Post("/user", handlers.CreateUSer(service))
}
