package app

import (
	"bahno_bot/feature/discord"
	"bahno_bot/generic/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env     *Env
	Mongo   *mongo.Client
	Discord *discord.Service
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Mongo = database.NewMongoDatabase(app.Env)
	app.Discord = discord.CreateDiscord(app.Env.DiscordToken)
	return app
}

func (app *Application) CloseDBConnection() {
	database.CloseMongoDBConnection(app.Mongo)
}
