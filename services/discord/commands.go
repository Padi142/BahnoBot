package discord

import (
	"bahno_bot/domain"
	"bahno_bot/repository"
	"bahno_bot/usecase"
	"context"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"time"
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

func (d *Service) BahnakCommand(db *mongo.Database, appId int) error {
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

		repo := repository.NewUserRepository(*db, "users")

		userUseCase := usecase.NewUserUseCase(repo, time.Duration(time.Second*10))

		userId := i.Member.User.ID

		profile, err := userUseCase.GetProfileByID(context.Background(), userId)

		if err != nil {
			//TODO: handle error
		}
		if profile == nil {
			newProfile := domain.User{UserId: userId, Name: i.Member.User.Username, PreferredSubstance: "bahno", ID: primitive.NewObjectID()}
			userUseCase.CreateUser(context.Background(), newProfile)
			err = SendInteractionResponse(s, i, "Vytvarim bahnici ucet")

			if err != nil {
				// Handle the error
			}
			return
		}

		err = SendInteractionResponse(s, i, "Tvoje preferovana substance je: "+profile.PreferredSubstance)

		if err != nil {
			// Handle the error
		}

	})
	return nil

}

func (d *Service) SubstanceCommand(db *mongo.Database, appId int) error {
	userRepo := repository.NewUserRepository(*db, "users")

	userUseCase := usecase.NewUserUseCase(userRepo, time.Duration(time.Second*10))

	substanceRepo := repository.NewSubstanceRepository(*db, "substances")

	substanceRepository := usecase.NewSubstanceUseCase(substanceRepo, time.Duration(time.Second*10))

	substances, err := substanceRepository.GetSubstances(context.Background())
	if err != nil {
		return err
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, substance := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  substance.Name,
			Value: substance.Value,
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

		userId := i.Member.User.ID

		profile, err := userUseCase.GetProfileByID(context.Background(), userId)

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

		if value == profile.PreferredSubstance {
			err = SendInteractionResponse(s, i, "Uz mas tuto substanci :warning:")
			return
		}
		found := false
		for _, substance := range substances {
			if substance.Value == value {
				_, err = userUseCase.SetPreferredSubstance(context.Background(), userId, value)
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

func (d *Service) BahnimCommand(db *mongo.Database, appId int) error {
	recordRepo := repository.NewRecordRepository(*db, "users")

	recordUseCase := usecase.NewRecordUseCase(recordRepo, time.Duration(time.Second*10))

	substanceRepo := repository.NewSubstanceRepository(*db, "substances")

	substanceRepository := usecase.NewSubstanceUseCase(substanceRepo, time.Duration(time.Second*10))

	substances, err := substanceRepository.GetSubstances(context.Background())
	if err != nil {
		return err
	}
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, substance := range substances {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  substance.Name,
			Value: substance.Value,
		}
	}

	command := &discordgo.ApplicationCommand{
		Name:        "bahnim",
		Description: "Zacne bahnit",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "bahnim",
				Description: "Vyber s jakou substanci chces bahnit",

				Choices: choices,
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
		value := ""
		if i.ApplicationCommandData().Options == nil {
			value = "bahno"
		} else {

			value = i.ApplicationCommandData().Options[0].Value.(string)
		}

		found := false
		for _, substance := range substances {
			if substance.Value == value {

				newRecord := domain.Record{
					ID:        primitive.NewObjectID(),
					Substance: value,
					Time:      time.Now(),
					CreatedAt: time.Now(),
				}

				record, err := recordUseCase.CreateNewRecord(context.Background(), i.Member.User.ID, newRecord)
				if err != nil {
					err = SendInteractionResponse(s, i, "Pri bahneni vznikla chyba. Pardor")
					log.Println(err)

					if err != nil {
						log.Println(err)
					}
					return
				}
				formattedTime := record.CreatedAt.Format("15:04 02.01")
				//timeStamp := fmt.Sprintf("<t:%d, d>", record.CreatedAt.Unix())
				err = SendInteractionResponse(s, i, "Pridano bahneni: **"+record.Substance+"** v: **"+formattedTime+"**")
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
