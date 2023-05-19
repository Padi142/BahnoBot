package discord

import (
	"bahno_bot/models"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (d *Service) InitCommands(db *mongo.Database) {

	d.BahnoCommand()
	d.BahnakCommand(db)
	d.SubstanceCommand(db)

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
