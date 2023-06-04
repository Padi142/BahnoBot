package discord

import (
	commands "bahno_bot/feature/discord/commands"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"gorm.io/gorm"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	discord         *discordgo.Session
	commands        []*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func OpenBot(token string, appId int, db *gorm.DB) error {
	userUseCase := user.NewUserUseCase(db)
	substanceUseCase := substance.NewSubstanceUseCase(db)
	recordUseCase := record.NewRecordUseCase(db)
	service := Service{}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil
	}

	service.discord = session

	ClearCommands(service.discord, appId)

	log.Println("Registering commands")
	err = GenericHandler(service.discord, appId)
	err = UserHandler(service.discord, appId, userUseCase, substanceUseCase, recordUseCase)
	err = SubstanceHandler(service.discord, appId, userUseCase, substanceUseCase)

	if err != nil {
		return err
	}

	err = service.discord.Open()
	if err != nil {
		return err
	}
	log.Println("Discord bot running")
	return nil
}

func GenericHandler(s *discordgo.Session, appId int) error {
	err := RegisterCommand(s, commands.BahnoCommand("bahno"), appId)

	return err
}

func UserHandler(s *discordgo.Session, appId int, userUseCase user.UseCase, substanceUseCase substance.UseCase, recordUseCase record.UseCase) error {
	err := RegisterCommand(s, commands.BahnakCommand("bahnak", userUseCase), appId)
	err = RegisterCommand(s, commands.BahnimCommand("bahnim", substanceUseCase, userUseCase, recordUseCase), appId)
	err = RegisterCommand(s, commands.LastBahneniCommand("last_bahneni", userUseCase, recordUseCase), appId)

	return err
}

func SubstanceHandler(s *discordgo.Session, appId int, userUseCase user.UseCase, substanceUseCase substance.UseCase) error {
	err := RegisterCommand(s, commands.SetPreferredSubstanceCommand("set_substance", substanceUseCase, userUseCase), appId)
	err = RegisterCommand(s, commands.PrintAllSubstances("substances", substanceUseCase), appId)

	return err
}

func RegisterCommand(s *discordgo.Session, command commands.Command, appId int) error {

	_, err := s.ApplicationCommandCreate(strconv.Itoa(appId), "", &command.Command)
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		command.Handler(s, i)
	})
	return err
}

func ClearCommands(s *discordgo.Session, appId int) {
	log.Println("Discord clearing old commands")

	cmd, err := s.ApplicationCommands(strconv.Itoa(appId), "")
	if err != nil {
		log.Println(err.Error())
	}

	// Iterate through the commands and delete them
	for _, command := range cmd {
		err = s.ApplicationCommandDelete(strconv.Itoa(appId), "", command.ID)
		if err != nil {
			log.Println(err.Error())

		}
	}
}
