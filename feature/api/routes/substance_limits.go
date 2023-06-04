package routes

import (
	"bahno_bot/feature/api/handlers"
	"bahno_bot/generic/substance_limit"

	"github.com/gofiber/fiber/v2"
)

func SubstanceLimitsRouter(app fiber.Router, useCase substance_limit.UseCase) {
	app.Get("/substance_limits/:user_id::substance_id", handlers.GetSubstanceLimit(useCase))
	app.Post("/substance_limits/", handlers.CreateSubstanceLimit(useCase))
}
