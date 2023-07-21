package commands

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
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
			amount = opt.FloatValue()
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
					Value:  rec.Substance.Label,
					Inline: true,
				},
				{
					Name:   "Date: ",
					Value:  GetTimeStamp(rec.CreatedAt, "F"),
					Inline: true,
				},
				{
					Name:   "Dose: ",
					Value:  fmt.Sprintf("%.2fg", rec.Amount),
					Inline: true,
				},
			},
		}

		err = SendInteractionResponseEmbed(s, i, embed)
	}
	return Command{Command: command, Handler: handler}
}

func GetRecordsCommand(name string, userUseCase user.UseCase, recordUseCase record.UseCase, substanceUseCase substance.UseCase) ComplexCommand {
	substances, err := substanceUseCase.GetSubstances()
	if err != nil {
		log.Println(err.Error())
		return ComplexCommand{}
	}
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: sub.Value,
		}
	}

	pageSize := 5
	page := 1
	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Prints your records",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "substance",
				Description: "Vyber si pro jakou substanci chces zobrazit historii",
				Choices:     choices,
				Required:    false,
			},
		},
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
		page = 0
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
		var records []models.Record
		footer := discordgo.MessageEmbedFooter{}

		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["substance"]; ok {
			substance := opt.StringValue()
			sub, err := substanceUseCase.GetSubstanceByValue(substance)
			if err != nil {
				err = SendInteractionResponse(s, i, err.Error())
				return
			}
			records, _, err = recordUseCase.GetPagedRecordsForSubstance(usr.ID, sub.ID, page, pageSize)
			footer = discordgo.MessageEmbedFooter{
				IconURL: GetUserAvatarUrl(i.Member.User),
				Text:    i.Member.User.Username + " - " + i.Member.User.ID + " ~ " + sub.Value,
			}
		} else {
			records, _, err = recordUseCase.GetPagedRecords(usr.ID, page, pageSize)
			footer = discordgo.MessageEmbedFooter{
				IconURL: GetUserAvatarUrl(i.Member.User),
				Text:    i.Member.User.Username + " - " + i.Member.User.ID,
			}
		}

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

		if len(records) == 0 {
			field := discordgo.MessageEmbedField{
				Name:  "Zadne zaznamy == spatny bahnak 😪",
				Value: "",
			}
			fields = append(fields, &field)
		}

		embed := &discordgo.MessageEmbed{
			Title:  "Zaznamy:",
			Color:  0x00ff00,
			Fields: fields,
			Footer: &footer,
			//Footer: &discordgo.MessageEmbedFooter{
			//	Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
			//},
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
		if i.MessageComponentData().CustomID == forwardButton.CustomID {
			message, err := s.ChannelMessage(i.ChannelID, i.Message.ID)
			splitFooter := strings.Split(i.Message.Embeds[0].Footer.Text, " - ")
			if !strings.Contains(splitFooter[len(splitFooter)-1], message.Interaction.User.ID) {

				return
			}
			page++
			if err != nil {
				err = SendInteractionResponse(s, i, err.Error())
				return
			}
			usr, err := userUseCase.GetProfileByDiscordID(message.Interaction.User.ID)

			var records []models.Record
			splitSubstance := strings.Split(i.Message.Embeds[0].Footer.Text, " ~ ")

			if len(splitSubstance) != 0 {
				sub, _ := substanceUseCase.GetSubstanceByValue(splitSubstance[len(splitSubstance)-1])
				records, _, err = recordUseCase.GetPagedRecordsForSubstance(usr.ID, sub.ID, page, pageSize)
			} else {
				records, _, err = recordUseCase.GetPagedRecords(usr.ID, page, pageSize)
			}

			if err != nil {
				log.Println(err.Error())
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

			if len(records) == 0 {
				field := discordgo.MessageEmbedField{
					Name:  "Zadne dalsi zaznamy",
					Value: "",
				}
				fields = append(fields, &field)
				forwardButton.Disabled = true
			}
			backButton.Disabled = false

			embed := &discordgo.MessageEmbed{
				Title:  "Zaznamy:",
				Color:  0x00ff00,
				Fields: fields,
				Footer: i.Message.Embeds[0].Footer,
				//Footer: &discordgo.MessageEmbedFooter{
				//	Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
				//},
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
		if i.MessageComponentData().CustomID == backButton.CustomID {
			message, err := s.ChannelMessage(i.ChannelID, i.Message.ID)
			splitFooter := strings.Split(i.Message.Embeds[0].Footer.Text, " - ")
			if !strings.Contains(splitFooter[len(splitFooter)-1], message.Interaction.User.ID) {

				return
			}
			page--

			if err != nil {
				err = SendInteractionResponse(s, i, err.Error())
				return
			}
			usr, err := userUseCase.GetProfileByDiscordID(message.Interaction.User.ID)
			var records []models.Record
			splitSubstance := strings.Split(i.Message.Embeds[0].Footer.Text, " ~ ")

			if len(splitSubstance) != 0 {
				sub, _ := substanceUseCase.GetSubstanceByValue(splitSubstance[len(splitSubstance)-1])
				records, _, err = recordUseCase.GetPagedRecordsForSubstance(usr.ID, sub.ID, page, pageSize)
			} else {
				records, _, err = recordUseCase.GetPagedRecords(usr.ID, page, pageSize)
			}

			if err != nil {
				log.Println(err.Error())
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

			if len(records) == 0 {
				field := discordgo.MessageEmbedField{
					Name:  "Zadne zaznamy == spatny bahnak 😪",
					Value: "",
				}
				fields = append(fields, &field)
				forwardButton.Disabled = true
				backButton.Disabled = false

			}

			embed := &discordgo.MessageEmbed{
				Title:  "Zaznamy:",
				Color:  0x00ff00,
				Fields: fields,
				Footer: i.Message.Embeds[0].Footer,
				//Footer: &discordgo.MessageEmbedFooter{
				//	Text: strconv.Itoa(page) + "/" + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageSize)))),
				//},
			}

			if page == 0 {
				backButton.Disabled = true
			}
			forwardButton.Disabled = false
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
	return ComplexCommand{Command: command,
		Handler: handler,
		Subhandlers: []func(s *discordgo.Session, i *discordgo.InteractionCreate){
			forwardButtonHandler, backButtonHandler,
		}}
}
func MuzuBahnit(name string, userUseCase user.UseCase, recordUseCase record.UseCase, substanceUseCase substance.UseCase) Command {
	substances, err := substanceUseCase.GetSubstances()
	if err != nil {
		log.Println(err.Error())
		return Command{}
	}

	substanceChoices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))
	substancesMap := make(map[string]models.Substance)

	for i, sub := range substances {
		substanceChoices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: sub.Value,
		}

		substancesMap[sub.Value] = sub
	}

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Rekne ti jestli muzes bahnit",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "substance",
				Description: "Vyber si jakou substanci chceš bahnit",
				Choices:     substanceChoices,
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
			err = SendInteractionResponse(s, i, err.Error())
			return
		}

		chosenSubstance := usr.PreferredSubstance.Value

		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["substance"]; ok {
			chosenSubstance = opt.StringValue()
		}

		// Invalid value from discord picker (shouldn't happen)
		sub, ok := substancesMap[chosenSubstance]
		if !ok {
			log.Println()
			return
		}

		title := "Muzes bahnit :thumbsup:"
		color := 0xa0fc7e

		rec, err := recordUseCase.GetLastRecordForSubstance(sub.ID, usr.ID)
		if err != nil {

			// User hasn't started bahnit yet!
			if errors.Is(err, gorm.ErrRecordNotFound) {
				_ = SendInteractionResponseEmbed(s, i, CreateMuzuBahnitEmbed(title, color, 0, rec, sub, false))
				return
			}

			err = SendInteractionResponse(s, i, err.Error())
			return
		}

		timeDiff := time.Now().Sub(rec.CreatedAt)
		hoursDiff := timeDiff.Hours()

		if hoursDiff < sub.RecommendedDosageMin*24 {
			title = "POZOR :warning: bahneni nedoporucujeme..."
			color = 0xf25c5c
		}

		err = SendInteractionResponseEmbed(s, i, CreateMuzuBahnitEmbed(title, color, float32(hoursDiff), rec, sub, true))
	}
	return Command{Command: command, Handler: handler}
}

func CreateMuzuBahnitEmbed(title string, color int, hoursDiff float32, rec models.Record, sub models.Substance, validTimestamp bool) *discordgo.MessageEmbed {
	var lastBahneniStr string
	// Checks if we should display the datetime
	if validTimestamp {
		lastBahneniStr = GetTimeStamp(rec.CreatedAt, "F")
	} else {
		lastBahneniStr = "-"
	}

	return &discordgo.MessageEmbed{
		Title: title,
		Color: color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Doporucena pauza: ",
				Value:  fmt.Sprintf("%.2fh - %.2fh", sub.RecommendedPauseMin*24, sub.RecommendedPauseMax*24),
				Inline: true,
			},
			{
				Name:   "Vase pauza: ",
				Value:  fmt.Sprintf("%.2fh, Posledni bahneni: %s", hoursDiff, lastBahneniStr),
				Inline: true,
			},
		},
	}
}
