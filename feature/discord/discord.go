package discord

import (
	"bahno_bot/generic/models"
	"log"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type Service struct {
	discord         *models.Discord
	commands        []*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func CreateDiscord(token string) *Service {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil
	}

	log.Println("Discord bot created")

	return &Service{discord: &models.Discord{
		Session: session,
	}}

}

func (d *Service) InitCommands(db gorm.DB, appId int) {

	//Clears old commands
	//d.ClearCommands()

	err := d.BahnoCommand(appId)
	err = d.BahnakCommand(db, appId)
	err = d.SubstanceCommand(db, appId)
	err = d.BahnimCommand(db, appId)
	err = d.LastBahneni(db, appId)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return
	}
	log.Println("Discord commands registered")

}

func (d *Service) OpenBot() error {
	err := d.discord.Session.Open()
	if err != nil {
		return err
	}
	log.Println("Discord bot running")
	return nil
}
