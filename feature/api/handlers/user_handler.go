package handlers

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/user"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary gets user by id
// @Description Gets the basic user info by their id
// @Tags user
// @Produce json
// @Param userId query string true "ID of the user to retrieve"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [get]
func GetUser(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdString := c.Query("userId")

		log.Println("API CALL: GetUser id: " + userIdString)

		if userIdString == "" {
			return c.JSON(fiber.Map{
				"error": "No user id provided",
			})
		}

		userId, err := strconv.ParseUint(userIdString, 10, 32)

		if userIdString == "" {
			return c.JSON(fiber.Map{
				"error": "Wrong userId format",
			})
		}

		userResult, err := useCase.GetProfileByID(uint(userId))
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

// GetUserDiscord godoc
// @Summary gets user by their discord id
// @Description Gets the basic user info by their discord id
// @Tags user
// @Produce json
// @Param discordId query string true "ID of the user to retrieve"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user_discord [get]
func GetUserDiscord(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdString := c.Query("discordId")

		log.Println("API CALL: GetUserDiscord id: " + userIdString)

		if userIdString == "" {
			return c.JSON(fiber.Map{
				"error": "No discord id provided",
			})
		}

		userResult, err := useCase.GetProfileByDiscordID(userIdString)
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
// @Tags user
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

		userResult, err := useCase.SetPreferredSubstance(usr.ID, usr.PreferredSubstanceID)
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
