package handlers

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/substance_limit"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

func CreateSubstanceLimit(useCase substance_limit.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("COCOEOJKFEWJFEKWFFEWKJWEFKWEJF")
		substance_limit := models.SubstanceLimit{}

		log.Printf("COZE %v\n", substance_limit)
		if err := c.BodyParser(&substance_limit); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		log.Printf("COZE %v\n", substance_limit)
		if err := useCase.CreateSubstanceLimit(substance_limit); err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		log.Printf("COZE %v\n", substance_limit)
		return c.Status(fiber.StatusCreated).JSON(&substance_limit)
	}
}
