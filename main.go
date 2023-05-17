package main

import (
	"bahno_bot/app"
	"os"
	"os/signal"
)

func main() {
	bahnakApp := app.App()

	env := bahnakApp.Env

	bahnakApp.Mongo.Database(env.DBName)
	defer bahnakApp.CloseDBConnection()

	//discord := bahnakApp.Discord

	bahnakApp.Discord.InitCommands()
	err := bahnakApp.Discord.OpenBot()
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
