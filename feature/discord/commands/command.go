package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"time"
)

type Command struct {
	Command discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Name    string
}

func SendInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) error {
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message,
			},
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func SendInteractionResponseEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) error {
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func GetTimeStamp(dateTime time.Time, stampType string) string {
	return "<t:" + strconv.FormatInt(dateTime.Unix(), 10) + ":" + stampType + ">"
}

func LogCommandUse(userName, commandName string) {
	log.Println("User " + userName + " ran command " + commandName)
}