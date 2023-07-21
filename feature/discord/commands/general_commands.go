package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
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
			log.Println(err)
		}

	}
	return Command{Command: command, Handler: handler}
}

func ApiCommand(name string) Command {
	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Returns link to api",
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
					Content: "https://bahno-api.krejzac.cz/swagger/index.html",
				},
			},
		)
		if err != nil {
			log.Println(err)
		}

	}
	return Command{Command: command, Handler: handler}
}
