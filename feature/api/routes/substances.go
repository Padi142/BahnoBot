package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/substance"
	"github.com/gofiber/fiber/v2"
)

func SubstancesRouter(app fiber.Router, useCase substance.UseCase) {
	app.Get("/substance/all", handlers.GetAllSubstances(useCase))
}
