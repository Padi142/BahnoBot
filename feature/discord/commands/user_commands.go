package commands

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func BahnakCommand(name string, userUseCase user.UseCase) Command {
	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Tvuj ucet bahnaka",
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
			log.Println(err.Error())
			return
		}
		if profile == nil {
			newProfile := models.User{ID: 0, DiscordID: userId, Username: i.Member.User.Username, PreferredSubstanceID: 1}
			err = userUseCase.CreateUser(newProfile)
			err = SendInteractionResponse(s, i, "Vytvarim bahnici ucet")

			if err != nil {
				// Handle the error
			}
			return
		}

		err = SendInteractionResponse(s, i, "Tvoje preferovana substance je: "+profile.PreferredSubstance.Label)

		if err != nil {
			log.Println(err.Error())
			return
		}

	}
	return Command{Command: command, Handler: handler}
}

func BahnimCommand(name string, substanceUseCase substance.UseCase, userUseCase user.UseCase, recordUseCase record.UseCase) Command {

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
		Description: "Zacne bahnit",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "substance",
				Description: "Vyber si jakou substanci chceš bahnit",
				Choices:     choices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "amount",
				Description: "Zadej množství v gramech",
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

		usr, err := userUseCase.GetProfileByDiscordID(i.Member.User.ID)

		if err != nil {
			_ = SendInteractionResponse(s, i, "Neexistujes 💀")
			return
		}

		oldSubstance := usr.PreferredSubstance.Value
		amount := 0

		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["substance"]; ok {
			oldSubstance = opt.StringValue()
		}

		if opt, ok := optionMap["amount"]; ok {
			amount = int(opt.IntValue())
		}

		found := false
		for _, sub := range substances {
			if sub.Value == oldSubstance {
				newRecord := models.Record{
					Substance:   sub,
					SubstanceID: sub.ID,
					CreatedAt:   time.Now(),
					Amount:      int(amount),
					UserID:      usr.ID,
				}

				rec, err := recordUseCase.CreateNewRecord(usr.ID, newRecord)
				if err != nil {
					err = SendInteractionResponse(s, i, "Pri bahneni vznikla chyba. Pardor")
					log.Println(err)

					if err != nil {
						log.Println(err)
					}
					return
				}
				//formattedTime := rec.CreatedAt.Format("15:04 02.01.2006")
				//timeStamp := fmt.Sprintf("<t:%d, d>", rec.CreatedAt.Unix())
				err = SendInteractionResponse(s, i, "Pridano bahneni: **"+rec.Substance.Label+"** "+GetTimeStamp(rec.CreatedAt, "R"))
				if err != nil {
					log.Println(err)

					err = SendInteractionResponse(s, i, "Pri bahneni vznikla chyba. Pardor")
					if err != nil {
						log.Println(err)
					}
					return
				}
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

	}
	return Command{Command: command, Handler: handler}
}

func LastBahneniCommand(name string, userUseCase user.UseCase, recordUseCase record.UseCase) Command {

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Prints your last record",
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

		usr, err := userUseCase.GetProfileByDiscordID(i.Member.User.ID)
		if err != nil {
			err = SendInteractionResponse(s, i, err.Error())
			return
		}

		rec, err := recordUseCase.GetLastRecord(usr.ID)
		if err != nil {
			err = SendInteractionResponse(s, i, err.Error())
			return
		}
		embed := &discordgo.MessageEmbed{
			Title: "Posledni bahno:",
			Color: 0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Substance: ",
					Value:  rec.Substance.Value,
					Inline: true,
				},
				{
					Name:   "Date: ",
					Value:  GetTimeStamp(rec.CreatedAt, "F"),
					Inline: true,
				},
				{
					Name:   "Dose: ",
					Value:  fmt.Sprintf("%dg", rec.Amount),
					Inline: true,
				},
			},
		}

		err = SendInteractionResponseEmbed(s, i, embed)
	}
	return Command{Command: command, Handler: handler}
}
