package app

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv          string `mapstructure:"APP_ENV"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	DiscordToken    string `mapstructure:"DISCORD_TOKEN"`
	ContextTimeout  int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPass          string `mapstructure:"DB_PASS"`
	GenericDBName   string `mapstructure:"GENERIC_DB_NAME"`
	SubstanceDBName string `mapstructure:"SUBSTANCE_DB_NAME"`
	AppID           int    `mapstructure:"APP_ID"`
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
