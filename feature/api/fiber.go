package fiber_feature

import (
	_ "bahno_bot/docs"
	"bahno_bot/feature/api/routes"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// @title Swagger Example API
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
func NewApiService(db *mongo.Database) {
	log.Println("Creating fiber api service")
	app := fiber.New()

	userRepo := user.NewUserRepository(*db, "users")
	userUseCase := user.NewUserUseCase(userRepo, time.Duration(time.Second*10))

	recordRepo := record.NewRecordRepository(*db, "users")
	recordUseCase := record.NewRecordUseCase(recordRepo, time.Duration(time.Second*10))

	substancesRepo := substance.NewSubstanceRepository(*db, "substances")
	substanceUseCase := substance.NewSubstanceUseCase(substancesRepo, time.Duration(time.Second*10))

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
