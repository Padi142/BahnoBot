package commands

import (
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Command discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Name    string
}

type ComplexCommand struct {
	Command     discordgo.ApplicationCommand
	Handler     func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Subhandlers []func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Name        string
}

func DeferInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		},
	)

	if err != nil {
		return err
	}
	return nil
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

func SendInteractionResponseComplex(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
				Components: components,
			},
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func EditInteractionResponseComplex(s *discordgo.Session, i *discordgo.InteractionCreate, edit *discordgo.WebhookEdit) error {
	_, err := s.InteractionResponseEdit(
		i.Interaction,
		edit,
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

func GetUserAvatarUrl(user *discordgo.User) string {
	return "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".png"
}
