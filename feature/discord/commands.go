package discord

import (
	"bahno_bot/generic/models"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func (d *Service) BahnoCommand(appId int) error {
	command := &discordgo.ApplicationCommand{
		Name:        "bahno",
		Description: "Bahno ?/",
	}

	_, err := d.discord.Session.ApplicationCommandCreate(
		strconv.Itoa(appId),
		"",
		command,
	)
	if err != nil {
		return err
	}
	d.discord.Session.AddHandler(func(
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

	})
	return nil
}

func (d *Service) ClearCommands() {
	commands, err := d.discord.Session.ApplicationCommands(d.discord.Session.State.User.ID, "")
	if err != nil {
		panic(err)
	}

	// Iterate through the commands and delete them
	for _, command := range commands {
		err = d.discord.Session.ApplicationCommandDelete(d.discord.Session.State.User.ID, "", command.ID)
		if err != nil {
			panic(err)
		}
	}
}
func (d *Service) BahnakCommand(db *gorm.DB, appId int) error {
	command := &discordgo.ApplicationCommand{
		Name:        "bahnak",
		Description: "Tvuj ucet bahnaka",
	}
	_, err := d.discord.Session.ApplicationCommandCreate(
		strconv.Itoa(appId),
		"",
		command,
	)

	if err != nil {
		return err
	}
	d.discord.Session.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		//Only handle this command
		if i.ApplicationCommandData().Name != command.Name {
			return
		}
		LogCommandUse(i.Member.User.Username, command.Name)

		userRepo := user.NewUserRepository(db)
		userUseCase := user.NewUserUseCase(userRepo)

		userId := i.Member.User.ID

		profile, err := userUseCase.GetProfileByDiscordID(userId)

		if err != nil {
			//TODO: handle error
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
			// Handle the error
		}

	})
	return nil

}

func (d *Service) SubstanceCommand(db *gorm.DB, appId int) error {
	userRepo := user.NewUserRepository(db)

	userUseCase := user.NewUserUseCase(userRepo)

	substanceRepo := substance.NewSubstanceRepository(db)

	substanceRepository := substance.NewSubstanceUseCase(substanceRepo)

	substances, err := substanceRepository.GetSubstances()
	if err != nil {
		return err
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: sub.Value,
		}
	}

	command := &discordgo.ApplicationCommand{
		Name:        "substance",
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
	_, err = d.discord.Session.ApplicationCommandCreate(
		strconv.Itoa(appId),
		"",
		command,
	)

	if err != nil {
		return err
	}
	d.discord.Session.AddHandler(func(
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

	})
	return nil

}

func (d *Service) BahnimCommand(db *gorm.DB, appId int) error {
	recordRepo := record.NewRecordRepository(db)

	recordUseCase := record.NewRecordUseCase(recordRepo)

	substanceRepo := substance.NewSubstanceRepository(db)

	substanceRepository := substance.NewSubstanceUseCase(substanceRepo)

	substances, err := substanceRepository.GetSubstances()
	if err != nil {
		return err
	}
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: sub.Value,
		}
	}

	command := &discordgo.ApplicationCommand{
		Name:        "bahnim",
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
	_, err = d.discord.Session.ApplicationCommandCreate(
		strconv.Itoa(appId),
		"",
		command,
	)

	if err != nil {
		return err
	}

	d.discord.Session.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		//Only handle this command
		if i.ApplicationCommandData().Name != command.Name {
			return
		}
		LogCommandUse(i.Member.User.Username, command.Name)

		userRepo := user.NewUserRepository(db)

		userUseCase := user.NewUserUseCase(userRepo)

		usr, err := userUseCase.GetProfileByDiscordID(i.Member.User.ID)

		if err != nil {
			SendInteractionResponse(s, i, "Neexistujes 💀")
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

	})
	return nil

}

func (d *Service) LastBahneni(db *gorm.DB, appId int) error {
	recordRepo := record.NewRecordRepository(db)

	recordUseCase := record.NewRecordUseCase(recordRepo)

	userRepo := user.NewUserRepository(db)

	userUseCase := user.NewUserUseCase(userRepo)

	command := &discordgo.ApplicationCommand{
		Name:        "last_bahneni",
		Description: "Prints your last record",
	}
	_, err := d.discord.Session.ApplicationCommandCreate(
		strconv.Itoa(appId),
		"",
		command,
	)

	if err != nil {
		return err
	}
	d.discord.Session.AddHandler(func(
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

		rec, err := recordUseCase.GetLatestRecord(usr.ID)
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
					Value:  fmt.Sprintf("%dg", rec.Amount),
					Inline: true,
				},
			},
		}

		err = SendInteractionResponseEmbed(s, i, embed)
	})
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
