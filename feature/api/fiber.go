package fiber_feature

import (
	_ "bahno_bot/docs"
	"bahno_bot/feature/api/routes"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
	"log"
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

	userRepo := user.NewUserRepository(db)
	recordRepo := record.NewRecordRepository(db)

	userUseCase := user.NewUserUseCase(userRepo, recordRepo)

	recordUseCase := record.NewRecordUseCase(recordRepo)

	substancesRepo := substance.NewSubstanceRepository(db)
	substanceUseCase := substance.NewSubstanceUseCase(substancesRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello bahno world!")
	})

	api := app.Group("/api")
	routes.UserRouter(api, userUseCase)
	routes.RecordsRoute(api, recordUseCase)
	routes.SubstancesRouter(api, substanceUseCase)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	go ApiListen(app)

}

func ApiListen(app *fiber.App) {
	log.Fatal(app.Listen(":8081"))
}
