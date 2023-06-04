package handlers

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/user"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Get all users
// @Description Gets the basic info about every user
// @Tags user
// @Produce json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [get]
func GetUsers(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("API CALL: GetUsers")

		userResult, err := useCase.GetUsers()

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"data": userResult,
		})
	}
}

// GetUser godoc
// @Summary gets user by id
// @Description Gets the basic user info by their id
// @Tags user
// @Produce json
// @Param id path int true "Account ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id} [get]
func GetUser(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// if userIdString == "" {
		// 	return c.JSON(fiber.Map{
		// 		"error": "No user id provided",
		// 	})
		// }

		userId, err := strconv.ParseUint(c.Params("id"), 10, 64)

		log.Printf("API CALL: GetUser id: %d\n", userId)

		// if userIdString == "" {
		// 	return c.JSON(fiber.Map{
		// 		"error": "Wrong userId format",
		// 	})
		// }

		userResult, err := useCase.GetProfileByID(uint(userId))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"data": userResult,
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
		discord_id := c.Params("discord_id")

		log.Println("API CALL: GetUserDiscord id: " + discord_id)

		if discord_id == "" {
			return c.JSON(fiber.Map{
				"error": "No discord id provided",
			})
		}

		userResult, err := useCase.GetProfileByDiscordID(discord_id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})

		}
		return c.JSON(fiber.Map{
			"data": userResult,
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
			"data": userResult,
		})
	}
}

// GetUserRecords godoc
// @Summary get all records of user
// @Description
// @Tags user
// @Produce json
// @Param id path int true "Account ID"
// @Body user query models.User{} true "new user body"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id}/records [get]
func GetUserRecords(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId, _ := strconv.ParseUint(c.Params("id"), 10, 64)
		records, err := useCase.GetUserRecords(uint(userId))

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

// GetLastUserRecord godoc
// @Summary get last record of given user
// @Description
// @Tags user
// @Param id path int true "Account ID"
// @Produce json
// @Body user query models.User{} true "new user body"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id}/records/last [get]
func GetLastUserRecord(useCase user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId, _ := strconv.ParseUint(c.Params("id"), 10, 64)
		record, err := useCase.GetLastUserRecord(uint(userId))

		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": record,
		})
	}
}
