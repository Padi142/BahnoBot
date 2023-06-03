package main

import (
	"bahno_bot/app"
	"os"
	"os/signal"
)

func main() {
	app.App()

	// env := bahnakApp.Env

	// dbClient := bahnakApp.Db.Database(env.SubstanceDBName)
	// defer bahnakApp.CloseDBConnection()

	//discord := bahnakApp.Discord

	// bahnakApp.Discord.InitCommands(dbClient, env.AppID)
	// err := bahnakApp.Discord.OpenBot()
	// if err != nil {
	// 	panic(err)
	// }

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
