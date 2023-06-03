package main

import (
	"bahno_bot/app"
	"os"
	"os/signal"
)

func main() {
	app.App()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
