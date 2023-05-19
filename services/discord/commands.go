package discord

import (
	"bahno_bot/domain"
	"bahno_bot/repository"
	"bahno_bot/usecase"
	"context"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (d *Service) BahnoCommand() error {
	command := &discordgo.ApplicationCommand{
		Name:        "bahno",
		Description: "Bahno ?/",
	}
	_, err := d.discord.Session.ApplicationCommandCreate(
		"1108445580407095329",
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
		data := i.ApplicationCommandData()
		switch data.Name {
		case command.Name:
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
	})
	return nil
}

func (d *Service) BahnakCommand(db *mongo.Database) error {
	command := &discordgo.ApplicationCommand{
		Name:        "bahnak",
		Description: "Tvuj ucet bahnaka",
	}
	_, err := d.discord.Session.ApplicationCommandCreate(
		"1108445580407095329",
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
		data := i.ApplicationCommandData()
		switch data.Name {
		case command.Name:

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
			}

			err = SendInteractionResponse(s, i, "Tvoje preferovana substance je "+profile.PreferredSubstance)

			if err != nil {
				// Handle the error
			}

		}
	})
	return nil

}

func (d *Service) SubstanceCommand(db *mongo.Database) error {
	command := &discordgo.ApplicationCommand{
		Name:        "substance",
		Description: "Set your preferred substance (bahno, mate, zeli, nikotin, kofein )",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "substance",
				Description: "Zmen svoji defaultni substanci",
				//Required:     true,
				Autocomplete: true,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name: "bahno",
					},
					{
						Name: "mate	",
					},
				},
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Bahno",
						Value: "bahno",
					},
					{
						Name:  "maté",
						Value: "mate",
					},
					{
						Name:  "Dablovo zeli",
						Value: "zeli",
					},
					{
						Name:  "Nikotin",
						Value: "nikotin",
					},
					{
						Name:  "Kofein",
						Value: "kofein",
					},
				},
			},
		},
	}
	_, err := d.discord.Session.ApplicationCommandCreate(
		"1108445580407095329",
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
		data := i.ApplicationCommandData()
		switch data.Name {
		case command.Name:

			repo := repository.NewUserRepository(*db, "users")

			userUseCase := usecase.NewUserUseCase(repo, time.Duration(time.Second*10))

			userId := i.Member.User.ID

			profile, err := userUseCase.GetProfileByID(context.Background(), userId)

			if err != nil {
				err = SendInteractionResponse(s, i, "Nemas bahnici ucet :warning:")

				if err != nil {
					// Handle the error
				}
				return
			}
			if i.Type.String() == profile.PreferredSubstance {
				err = SendInteractionResponse(s, i, "Uz mas tuto sunstanci :warning:")
				return
			}

			switch i.Type.String() {
			case "bahno":
				userUseCase.SetPreferredSubstance(context.Background(), userId, i.Type.String())
			case "mate":
				userUseCase.SetPreferredSubstance(context.Background(), userId, i.Type.String())
			case "zeli":
				userUseCase.SetPreferredSubstance(context.Background(), userId, i.Type.String())
			case "nikotin":
				userUseCase.SetPreferredSubstance(context.Background(), userId, i.Type.String())
			case "kofein":
				userUseCase.SetPreferredSubstance(context.Background(), userId, i.Type.String())

			default:
				err = SendInteractionResponse(s, i, "Netrepej somary :warning: ")
				if err != nil {
					// Handle the error
				}
				return
			}
			err = SendInteractionResponse(s, i, "Substance updatnuta :white_check_mark: ")
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
