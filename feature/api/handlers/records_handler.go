package handlers

import (
	"bahno_bot/generic/record"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllRecords godoc
// @Summary gets all records
// @Tags record
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/records [get]
func GetAllRecords(useCase record.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId, err := strconv.ParseUint(c.Params("id"), 10, 64)

		records, err := useCase.GetAllRecords(uint(userId))

		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": records,
		})
	}
}
