package handlers

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/substance_limit"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetSubstanceLimit godoc
// @Summary get substance limit by id
// @Tags substance_limit
// @Produce json
// @Param user_id path int true "User ID"
// @Param substance_id path int true "Substance ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/substance_limits/{user_id}:{substance_id} [get]
func GetSubstanceLimit(useCase substance_limit.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("user_id"))
		substanceId, _ := strconv.Atoi(c.Params("substance_id"))
		substanceResult, err := useCase.GetSubstanceLimit(uint(userId), uint(substanceId))
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

// CreateSubstanceLimit godoc
// @Summary create new substance limit
// @Tags substance_limit
// @Produce json
// @Body subsatnce_limit query models.SubstanceLimit{} true "new substance limit body"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/substance_limits [post]
func CreateSubstanceLimit(useCase substance_limit.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		substance_limit := models.SubstanceLimit{}

		if err := c.BodyParser(&substance_limit); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := useCase.CreateSubstanceLimit(substance_limit); err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&substance_limit)
	}
}
