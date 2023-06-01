package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/user"
	"github.com/gofiber/fiber/v2"
)

func UserRoter(app fiber.Router, useCase user.UseCase) {
	app.Get("/user", handlers.GetUser(useCase))
	//app.Post("/user", handlers.CreateUSer(service))
}
