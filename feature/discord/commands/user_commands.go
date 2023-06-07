package commands

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math"
	"strconv"
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
				Type:        discordgo.ApplicationCommandOptionNumber,
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
		amount := 0.0

		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["substance"]; ok {
			oldSubstance = opt.StringValue()
		}

		if opt, ok := optionMap["amount"]; ok {
			amount = float64(opt.IntValue())
		}

		found := false
		for _, sub := range substances {
			if sub.Value == oldSubstance {
				newRecord := models.Record{
					Substance:   sub,
					SubstanceID: sub.ID,
					CreatedAt:   time.Now(),
					Amount:      amount,
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
				if rec.Substance.Value == "weed" {
					_, err := s.ChannelMessageSend(i.ChannelID, "https://media.discordapp.net/attachments/662506011499823115/876479116180869120/smoka.gif")
					if err != nil {
						log.Println(err)
					}
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

func GetRecordsCommand(name string, userUseCase user.UseCase, recordUseCase record.UseCase) ComplexCommand {
	pageSize := 5
	page := 1
	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Prints your records",
	}

	backButton := discordgo.Button{
		Label:    "<",
		Style:    discordgo.SuccessButton,
		CustomID: "all_records_button_back",
		Disabled: true,
	}
	forwardButton := discordgo.Button{
		Label:    ">",
		Style:    discordgo.SuccessButton,
		CustomID: "all_records_button_forward",
		Disabled: false,
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

		records, count, err := recordUseCase.GetPagedRecords(usr.ID, 1, pageSize)
		if err != nil {
			err = SendInteractionResponse(s, i, err.Error())
			return
		}
		var fields []*discordgo.MessageEmbedField

		for _, rec := range records {
			field := discordgo.MessageEmbedField{
				Name:   rec.Substance.Label,
				Value:  "📅: " + GetTimeStamp(rec.CreatedAt, "f") + "\n ⚖️:" + strconv.FormatFloat(rec.Amount, 'f', -1, 64),
				Inline: false,
			}
			fields = append(fields, &field)
		}

		embed := &discordgo.MessageEmbed{
			Title:  "Zaznamy:",
			Color:  0x00ff00,
			Fields: fields,
			Footer: &discordgo.MessageEmbedFooter{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
			},
		}

		// Create an action row component
		actionRow := discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{&backButton, &forwardButton},
		}

		err = SendInteractionResponseComplex(s, i, embed, []discordgo.MessageComponent{&actionRow})
		if err != nil {
			log.Println("Error executing discord command: " + name)
			log.Println(err.Error())
		}
	}
	forwardButtonHandler := func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		usr, err := userUseCase.GetProfileByDiscordID(i.Member.User.ID)
		records, count, err := recordUseCase.GetPagedRecords(usr.ID, page-1, pageSize)

		if err != nil {
			log.Println(err.Error())
		}
		if i.MessageComponentData().CustomID == forwardButton.CustomID {

			//if page >= int(math.Ceil(float64(count)/float64(pageSize))) {
			//	return
			//}

			page++

			var fields []*discordgo.MessageEmbedField

			for _, rec := range records {

				field := discordgo.MessageEmbedField{
					Name:   rec.Substance.Label,
					Value:  "📅: " + GetTimeStamp(rec.CreatedAt, "f") + "\n ⚖️:" + strconv.FormatFloat(rec.Amount, 'f', -1, 64),
					Inline: false,
				}
				fields = append(fields, &field)
			}

			embed := &discordgo.MessageEmbed{
				Title:  "Zaznamy:",
				Color:  0x00ff00,
				Fields: fields,
				Footer: &discordgo.MessageEmbedFooter{
					Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
				},
			}

			if page == int(math.Ceil(float64(count)/float64(pageSize))) {
				forwardButton.Disabled = true
			}
			if page != 1 {
				backButton.Disabled = false
			}
			actionRow := discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{&backButton, &forwardButton},
			}
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Components: []discordgo.MessageComponent{&actionRow},
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
				},
			})
			if err != nil {
				log.Println("Error executing discord command: " + name)
				log.Println(err.Error())
			}
		}
	}
	backButtonHandler := func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		//log.Println("handler 2")
		//
		//if i.Type == discordgo.InteractionMessageComponent && i.MessageComponentData().CustomID == backButton.CustomID {
		//
		//	if page < 2 {
		//		return
		//	}
		//	page--
		//
		//	records, count, err = recordUseCase.GetPagedRecords(usr.ID, page, pageSize)
		//	if err != nil {
		//		err = SendInteractionResponse(s, i, err.Error())
		//		return
		//	}
		//	var fields []*discordgo.MessageEmbedField
		//
		//	for _, rec := range records {
		//		field := discordgo.MessageEmbedField{
		//			Name:   rec.Substance.Label,
		//			Value:  "📅: " + GetTimeStamp(rec.CreatedAt, "f") + "\n ⚖️:" + strconv.FormatFloat(rec.Amount, 'f', -1, 64),
		//			Inline: false,
		//		}
		//		fields = append(fields, &field)
		//	}
		//
		//	msg.Embed.Fields = fields
		//	msg.Embed.Footer = &discordgo.MessageEmbedFooter{
		//		Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
		//	}
		//
		//	if page < 2 {
		//		backButton.Disabled = true
		//	}
		//	msg.Components = []discordgo.MessageComponent{
		//		&backButton, &forwardButton,
		//	}
		//
		//	_, err = s.ChannelMessageSendComplex(i.ChannelID, msg)
		//	if err != nil {
		//		log.Println("Error executing discord command: " + name)
		//	}
		//}

	}
	return ComplexCommand{Command: command,
		Handler: handler,
		Subhandlers: []func(s *discordgo.Session, i *discordgo.InteractionCreate){
			forwardButtonHandler, backButtonHandler,
		}}
}
