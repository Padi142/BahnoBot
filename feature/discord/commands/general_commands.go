package commands

import (
	"github.com/bwmarrin/discordgo"
)

func BahnoCommand(name string) Command {
	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Bahno ?/",
	}

	handler := func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		//Only handle this command
		if i.ApplicationCommandData().Name != command.Name {
			return
		}
		LogCommandUse(i.Member.User.Username, command.Name)

		err := s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "ahoj",
				},
			},
		)
		if err != nil {
			// Handle the error
		}

	}
	return Command{Command: command, Handler: handler}
}
