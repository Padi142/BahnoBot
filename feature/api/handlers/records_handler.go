package handlers

import (
	"bahno_bot/generic/record"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

// GetLastRecord godoc
// @Summary gets last record for given user
// @Description Gets the last record from records for given user id and substance id, if no substance id provided, gets last record for all substances
// @Tags root
// @Produce json
// @Param userId query string true "ID of the user to retrieve last record"
// @Param substance query string false "value of the substance"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /record/last [get]
func GetLastRecord(useCase record.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Query("userId")
		substance := c.Query("substance")

		log.Println("API CALL: GetLastRecord id: " + userId)

		if userId == "" {
			return c.JSON(fiber.Map{
				"error": "No user id provided",
			})
		}
		if substance == "" {
			//TODO
		}

		userResult, err := useCase.GetLatestRecord(context.Background(), userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"record": userResult,
		})
	}
}
