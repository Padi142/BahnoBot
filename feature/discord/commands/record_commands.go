package commands

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
)

func GetLeaderboardCommand(name string, substanceUseCase substance.UseCase, recordUseCase record.UseCase) Command {

	durations := make([]*discordgo.ApplicationCommandOptionChoice, len(models.Durations))
	for i, duration := range models.Durations {
		durations[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  string(duration),
			Value: string(duration),
		}
	}
	substances, err := substanceUseCase.GetSubstances()
	if err != nil {
		log.Println(err.Error())
		return Command{}
	}
	substanceChoices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		substanceChoices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: int(sub.ID),
		}
	}

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Gets leaderboard",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "substance",
				Description: "Vyber si jakou substanci chceš bahnit",
				Choices:     substanceChoices,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "time interval",
				Description: "Vyber v jakem casovem intervalu chces ukazat leaderbard",
				Choices:     durations,
				Required:    false,
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

		duration := string(models.Durations[3])
		options := i.ApplicationCommandData().Options
		substancePresent := false
		chosenSubstance := 0
		chosenSubstanceName := ""

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["time"]; ok {
			duration = opt.StringValue()
		}
		if opt, ok := optionMap["substance"]; ok {
			substancePresent = true
			chosenSubstance = int(opt.IntValue())
			chosenSubstanceName = opt.Name
		}
		var fields []*discordgo.MessageEmbedField
		if substancePresent {
			leaderboard, err := recordUseCase.GetLeaderboardTimeForSubstance(duration, uint(chosenSubstance))
			if err != nil {
				err = SendInteractionResponse(s, i, err.Error())
				return
			}

			for i, usr := range leaderboard {
				field := discordgo.MessageEmbedField{
					Name:   "",
					Value:  strconv.Itoa(i+1) + ") " + usr.User.Username + " " + strconv.Itoa(int(usr.Occurrence)) + " bahnění",
					Inline: false,
				}
				fields = append(fields, &field)
			}

		} else {

			leaderboard, err := recordUseCase.GetLeaderboardTime(duration)
			if err != nil {
				err = SendInteractionResponse(s, i, err.Error())
				return
			}

			for i, usr := range leaderboard {
				field := discordgo.MessageEmbedField{
					Name:   "",
					Value:  strconv.Itoa(i+1) + ") " + usr.User.Username + " " + strconv.Itoa(int(usr.Occurrence)) + " bahnění",
					Inline: false,
				}
				fields = append(fields, &field)
			}
		}
		if err != nil {
			err = SendInteractionResponse(s, i, err.Error())
			return
		}

		embed := discordgo.MessageEmbed{
			Title:  "Nejvetsi bahnaci: ",
			Fields: fields,
			Color:  0x00ff00,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Time: " + string(duration) + " Substance: " + chosenSubstanceName,
			},
		}

		err = SendInteractionResponseEmbed(s, i, &embed)
	}
	return Command{Command: command, Handler: handler}
}
