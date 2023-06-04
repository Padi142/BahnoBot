package app

import (
	api "bahno_bot/feature/api"
	"bahno_bot/feature/discord"
	"bahno_bot/generic/database"

	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Db = database.NewDatabase(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBName, uint(app.Env.DBPort))

	err := discord.OpenBot(app.Env.DiscordToken, app.Env.AppID, app.Db)
	if err != nil {
		panic(err)
	}

	api.NewApiService(app.Db)

	// recordRepo := record.NewRecordRepository(app.Db)
	// record := models.Record{Amount: 69, SubstanceID: 2, UserID: 1, Time: time.Now()}
	// r, _ := recordRepo.GetLast(1)

	// fmt.Printf("%v", r)

	// userUseCase := user.NewUserUseCase(userRepo, time.Duration(time.Second*10))
	// userUseCase.GetProfileByID()

	return app
}
