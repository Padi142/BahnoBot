package app

import (
	api "bahno_bot/feature/api"
	"bahno_bot/feature/discord"
	"bahno_bot/generic/database"
	"strconv"

	"gorm.io/gorm"
)

type Application struct {
	Env     *Env
	Db      *gorm.DB
	Discord *discord.Service
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	port, _ := strconv.Atoi(app.Env.DBPort)
	app.Db = database.NewDatabase(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBName, uint(port))

	app.Discord = discord.CreateDiscord(app.Env.DiscordToken)

	api.NewApiService(app.Db)

	app.Discord.InitCommands(app.Db, app.Env.AppID)
	err := app.Discord.OpenBot()
	if err != nil {
		panic(err)
	}

	// recordRepo := record.NewRecordRepository(app.Db)
	// record := models.Record{Amount: 69, SubstanceID: 2, UserID: 1, Time: time.Now()}
	// r, _ := recordRepo.GetLast(1)

	// fmt.Printf("%v", r)

	// userUseCase := user.NewUserUseCase(userRepo, time.Duration(time.Second*10))
	// userUseCase.GetProfileByID()

	return app
}
