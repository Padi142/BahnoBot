package handlers

import (
	"bahno_bot/generic/substance"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

// GetAllSubstances godoc
// @Summary gets all substances from db
// @Description Gets al substances that are by default accessible to user
// @Tags substance
// @Produce json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /substances/all [get]
func GetAllSubstances(useCase substance.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		log.Println("API CALL: GetAllSubstances")

		substanceResult, err := useCase.GetSubstances(context.Background())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"substances": substanceResult,
		})
	}
}
