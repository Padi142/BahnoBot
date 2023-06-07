package commands

import (
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func SetPreferredSubstanceCommand(name string, substanceUseCase substance.UseCase, userUseCase user.UseCase) Command {

	substances, err := substanceUseCase.GetSubstances()
	if err != nil {
		log.Println(err.Error())
		return Command{}
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: sub.Value,
		}
	}

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Set your preferred substance (bahno, mate, zeli, nikotin, kofein )",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "substance",
				Description: "Zmen svoji defaultni substanci",
				Required:    true,
				Choices:     choices,
			},
		},
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

		userId := i.Member.User.ID

		profile, err := userUseCase.GetProfileByDiscordID(userId)

		if err != nil {
			err = SendInteractionResponse(s, i, "Nemas bahnici ucet :warning:")

			if err != nil {
				log.Println(err)
				return
			}
			return
		}
		if i.ApplicationCommandData().Options == nil {
			err = SendInteractionResponse(s, i, "Napis validni substanci pls ")

		}

		value := i.ApplicationCommandData().Options[0].Value.(string)

		if value == profile.PreferredSubstance.Value {
			err = SendInteractionResponse(s, i, "Uz mas tuto substanci :warning:")
			return
		}
		found := false
		for _, sub := range substances {
			if sub.Value == value {
				_, err = userUseCase.SetPreferredSubstance(profile.ID, sub.ID)
				found = true
			}
		}
		if !found {
			err = SendInteractionResponse(s, i, "Netrepej somary :warning: ")
			if err != nil {
				log.Println(err)
			}
			return
		}
		err = SendInteractionResponse(s, i, "Substance updatnuta :white_check_mark: ")

	}
	return Command{Command: command, Handler: handler}
}

func PrintAllSubstances(name string, substanceUseCase substance.UseCase) Command {

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Prints all available substances",
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

		substances, err := substanceUseCase.GetSubstances()

		if err != nil {
			err = SendInteractionResponse(s, i, "Error pri ziskavani subtanci")

			if err != nil {
				log.Println(err)
				return
			}
			return
		}
		var fields []*discordgo.MessageEmbedField

		for _, sub := range substances {
			field := discordgo.MessageEmbedField{
				Name:   sub.Label,
				Value:  "Doporucena davka: " + strconv.FormatFloat(sub.RecommendedDosageMin, 'f', -1, 64) + " - " + strconv.FormatFloat(sub.RecommendedDosageMax, 'f', -1, 64) + "g \n" + "Doporucena pauza: " + strconv.FormatFloat(sub.RecommendedPauseMin, 'f', -1, 64) + " - " + strconv.FormatFloat(sub.RecommendedPauseMax, 'f', -1, 64) + "dni",
				Inline: false,
			}
			fields = append(fields, &field)
		}

		embed := discordgo.MessageEmbed{
			Title:  "Substance",
			Fields: fields,
			Color:  0x00ff00,
		}

		err = SendInteractionResponseEmbed(s, i, &embed)

	}
	return Command{Command: command, Handler: handler}
}
