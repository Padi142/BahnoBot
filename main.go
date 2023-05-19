package main

import (
	"bahno_bot/app"
	"os"
	"os/signal"
)

func main() {
	bahnakApp := app.App()

	env := bahnakApp.Env

	dbClient := bahnakApp.Mongo.Database(env.SubstanceDBName)
	defer bahnakApp.CloseDBConnection()

	//discord := bahnakApp.Discord

	bahnakApp.Discord.InitCommands(dbClient)
	err := bahnakApp.Discord.OpenBot()
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
