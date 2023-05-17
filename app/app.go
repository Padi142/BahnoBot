package app

import (
	"bahno_bot/services/discord"
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
	app.Mongo = NewMongoDatabase(app.Env)
	app.Discord = discord.CreateDiscord(app.Env.DiscordToken)
	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
