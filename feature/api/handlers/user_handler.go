package handlers

import (
	"bahno_bot/generic/user"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

//// AddBook is handler/controller which creates Books in the BookShop
//func CreateUser(useCase user.UseCase) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		var requestBody entities.Book
//		err := c.BodyParser(&requestBody)
//		if err != nil {
//			c.Status(http.StatusBadRequest)
//			return c.JSON(presenter.BookErrorResponse(err))
//		}
//		if requestBody.Author == "" || requestBody.Title == "" {
//			c.Status(http.StatusInternalServerError)
//			return c.JSON(presenter.BookErrorResponse(errors.New(
//				"Please specify title and author")))
//		}
//		result, err := service.InsertBook(&requestBody)
//		if err != nil {
//			c.Status(http.StatusInternalServerError)
//			return c.JSON(presenter.BookErrorResponse(err))
//		}
//		return c.JSON(presenter.BookSuccessResponse(result))
//	}
//}

// GetBooks is handler/controller which lists all Books from the BookShop
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
