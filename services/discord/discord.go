package discord

import (
	"bahno_bot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
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

	fmt.Println("Discord bot created")

	return &Service{discord: &models.Discord{
		Session: session,
	}}

}

func (d *Service) InitCommands() {
	_, err := d.discord.Session.ApplicationCommandBulkOverwrite("1108445580407095329", "1035650956383227995", []*discordgo.ApplicationCommand{
		{
			Name:        "bahno",
			Description: "Bahno ?/",
		},
		{
			Name:        "bahnim",
			Description: "Bahnis ??",
		},
	})
	if err != nil {
		// Handle the error
	}
	d.discord.Session.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "bahno":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "BAHNOOO",
					},
				},
			)
			if err != nil {
				// Handle the error
			}

		case "bahnim":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: ":warning: Pozor mame tu bahnaka :warning:",
					},
				},
			)
			if err != nil {
				// Handle the error
			}
		}
	})
	fmt.Println("Discord commands registered")

}

func (d *Service) OpenBot() error {
	err := d.discord.Session.Open()
	if err != nil {
		return err
	}
	fmt.Println("Discord bot running")
	return nil
}
