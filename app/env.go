package app

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	DiscordToken string `mapstructure:"DISCORD_TOKEN"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       int    `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	DBName       string `mapstructure:"DB_NAME"`
	AppID        string `mapstructure:"APP_ID"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
