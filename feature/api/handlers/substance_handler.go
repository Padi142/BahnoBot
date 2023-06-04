package handlers

import (
	"bahno_bot/generic/substance"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllSubstances godoc
// @Summary gets all substances from db
// @Description Gets all substances that are by default accessible to user
// @Tags substance
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/substances [get]
func GetAllSubstances(useCase substance.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("API CALL: GetAllSubstances")

		substanceResult, err := useCase.GetSubstances()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"data": substanceResult,
		})
	}
}

// GetSubstance godoc
// @Summary get substance by id
// @Tags substance
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/substances/{id} [get]
func GetSubstance(useCase substance.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		substanceId, _ := strconv.Atoi(c.Params("id"))
		substanceResult, err := useCase.GetSubstance(uint(substanceId))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"data": substanceResult,
		})

	}
}
