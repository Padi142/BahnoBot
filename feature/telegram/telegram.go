package telegram

import (
	telegram_commands "bahno_bot/feature/telegram/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type telegramBot struct {
	tgbot *tgbotapi.BotAPI
}

func OpenBot(telegramToken string) error {

	tgbot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Printf("Error occured while creating Telegram bot: %s\n", err)
	}
	bot := telegramBot{tgbot: tgbot}

	err = bot.registerGeneralCommandsHandler()
	if err != nil {
		log.Printf("Error registering commands: %s", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() { // ignore any non-Message updates
			continue
		}

		switch update.Message.Command() {
		case "bahno":
			log.Printf("Telegram user: %s ran bahno command", update.Message.Chat.Title)
			if err = telegram_commands.BahnoCommand(update, bot.tgbot); err != nil {
				log.Printf("telegram bahno error: %s", err)
			}
		}

	}

	return nil
}
func (b *telegramBot) registerGeneralCommandsHandler() error {
	genericCommands := []tgbotapi.BotCommand{
		{Command: "/bahno", Description: "Joooo miluju bahnooo mnaaam"},
	}

	scopeEveryone := tgbotapi.NewBotCommandScopeDefault()
	config := tgbotapi.NewSetMyCommandsWithScope(scopeEveryone, genericCommands...)
	_, err := b.tgbot.Request(config)
	if err != nil {
		return err
	}
	return nil
}
