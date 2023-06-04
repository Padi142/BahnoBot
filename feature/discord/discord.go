package discord

import (
	commands "bahno_bot/feature/discord/commands"
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"log"
	"reflect"

	"gorm.io/gorm"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	discord         *discordgo.Session
	commands        map[string]*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (s *Service) GetApplicationCommandsMap(appId string) (commandsMap map[string]*discordgo.ApplicationCommand, err error) {
	commands, err := s.discord.ApplicationCommands(appId, "")

	if err != nil {
		return
	}

	commandsMap = map[string]*discordgo.ApplicationCommand{}
	for _, command := range commands {
		commandsMap[command.Name] = command
	}

	return
}

func OpenBot(token string, appId string, db *gorm.DB) error {
	userUseCase := user.NewUserUseCase(db)
	substanceUseCase := substance.NewSubstanceUseCase(db)
	recordUseCase := record.NewRecordUseCase(db)
	service := Service{}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil
	}

	service.discord = session

	oldCommands, _ := service.GetApplicationCommandsMap(appId)
	service.commands = oldCommands

	log.Println("Registering commands")
	err = GenericHandler(&service, appId)
	err = UserHandler(&service, appId, userUseCase, substanceUseCase, recordUseCase)
	err = SubstanceHandler(&service, appId, userUseCase, substanceUseCase)

	if err != nil {
		return err
	}

	// Delete unused commands
	newCommands, _ := service.GetApplicationCommandsMap(appId)
	for name := range oldCommands {
		if val, ok := newCommands[name]; !ok {
			session.ApplicationCommandDelete(appId, "", val.ID)
		}
	}
	service.commands = newCommands

	err = service.discord.Open()
	if err != nil {
		return err
	}
	log.Println("Discord bot running")
	return nil
}

func GenericHandler(s *Service, appId string) error {
	err := RegisterCommand(s, commands.BahnoCommand("bahno"), appId)

	return err
}

func UserHandler(s *Service, appId string, userUseCase user.UseCase, substanceUseCase substance.UseCase, recordUseCase record.UseCase) error {
	err := RegisterCommand(s, commands.BahnakCommand("bahnak", userUseCase), appId)
	err = RegisterCommand(s, commands.BahnimCommand("bahnim", substanceUseCase, userUseCase, recordUseCase), appId)
	err = RegisterCommand(s, commands.LastBahneniCommand("last_bahneni", userUseCase, recordUseCase), appId)

	return err
}

func SubstanceHandler(s *Service, appId string, userUseCase user.UseCase, substanceUseCase substance.UseCase) error {
	err := RegisterCommand(s, commands.SetPreferredSubstanceCommand("set_substance", substanceUseCase, userUseCase), appId)
	err = RegisterCommand(s, commands.PrintAllSubstances("substances", substanceUseCase), appId)

	return err
}

func RegisterCommand(s *Service, command commands.Command, appId string) error {
	if val, ok := s.commands[command.Command.Name]; ok {
		if val.Description == command.Command.Description &&
			reflect.DeepEqual(val.Options, command.Command.Options) {
			return nil
		}
	}

	log.Printf("Registering command - %s\n", command.Command.Name)

	_, err := s.discord.ApplicationCommandCreate(appId, "", &command.Command)
	s.discord.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		command.Handler(s, i)
	})
	return err
}
