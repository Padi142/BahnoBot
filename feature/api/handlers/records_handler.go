package handlers

import (
	"bahno_bot/generic/record"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

// GetLastRecord godoc
// @Summary gets last record for given user
// @Description Gets the last record from records for given user id and substance id, if no substance id provided, gets last record for all substances
// @Tags record
// @Produce json
// @Param userId query string true "ID of the user to retrieve last record"
// @Param substance query string false "value of the substance"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /record/last [get]
func GetLastRecord(useCase record.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdString := c.Query("userId")
		substance := c.Query("substance")

		log.Println("API CALL: GetLastRecord id: " + userIdString)

		if userIdString == "" {
			return c.JSON(fiber.Map{
				"error": "No user id provided",
			})
		}
		if substance == "" {
			//TODO
		}
		userId, err := strconv.ParseUint(userIdString, 10, 32)

		if userIdString == "" {
			return c.JSON(fiber.Map{
				"error": "Wrong userId format",
			})
		}

		userResult, err := useCase.GetLatestRecord(uint(userId))
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
