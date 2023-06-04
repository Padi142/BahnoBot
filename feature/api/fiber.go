package fiber_feature

import (
	_ "bahno_bot/docs"
	"bahno_bot/feature/api/routes"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/substance_limit"
	"bahno_bot/generic/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

// NewApiService @title BahnoBot api
// @version 1.0
// @description This is a sample server Bahno server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /
// @schemes http
func NewApiService(db *gorm.DB) {
	log.Println("Creating fiber api service")
	app := fiber.New()
	app.Use(logger.New())

	userUseCase := user.NewUserUseCase(db)

	recordUseCase := record.NewRecordUseCase(db)

	substanceUseCase := substance.NewSubstanceUseCase(db)

	substanceLimitRepo := substance_limit.NewSubstanceRepository(db)
	substanceLimitUseCase := substance_limit.NewSubstanceUseCase(substanceLimitRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello bahno world!")
	})

	api := app.Group("/api")
	routes.UserRouter(api, userUseCase)
	routes.RecordsRoute(api, recordUseCase)
	routes.SubstancesRouter(api, substanceUseCase)
	routes.SubstanceLimitsRouter(api, substanceLimitUseCase)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	go ApiListen(app)

}

func ApiListen(app *fiber.App) {
	log.Fatal(app.Listen(":8081"))
}
