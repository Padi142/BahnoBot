package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/record"
	"github.com/gofiber/fiber/v2"
)

func RecordsRoute(app fiber.Router, useCase record.UseCase) {
	app.Get("/record/last", handlers.GetLastRecord(useCase))
	//app.Post("/user", handlers.CreateUSer(service))
}
