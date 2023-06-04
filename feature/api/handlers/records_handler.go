package handlers

import (
	"bahno_bot/generic/record"

	"github.com/gofiber/fiber/v2"
)

// GetLastRecord godoc
// @Summary gets all records
// @Description Gets the last record from records for given user id and substance id, if no substance id provided, gets last record for all substances
// @Tags record
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/records [get]
func GetAllRecords(useCase record.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		records, err := useCase.GetAllRecords()

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
