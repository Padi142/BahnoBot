package handlers

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/user"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

// GetUser godoc
// @Summary gets user by id
// @Description Gets the basic user info by their id
// @Tags root
// @Produce json
// @Param userId query string true "ID of the user to retrieve"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [get]
func GetUser(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Query("userId")

		log.Println("API CALL: GetUser id: " + userId)

		if userId == "" {
			return c.JSON(fiber.Map{
				"error": "No user id provided",
			})
		}

		userResult, err := useCase.GetProfileByID(context.Background(), userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"user": userResult,
		})
	}
}

// UpdateUser godoc
// @Summary updates user with incoming json struct
// @Description Send new user struct to update user in db
// @Tags root
// @Produce json
// @Body user query models.User{} true "new user body"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [put]
func UpdateUser(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("API CALL: UpdateUser ")

		usr := models.User{}

		err := c.BodyParser(&usr)

		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userResult, err := useCase.SetPreferredSubstance(context.Background(), usr.UserId, usr.PreferredSubstance)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"user": userResult,
		})
	}
}
