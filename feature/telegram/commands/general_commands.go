package telegram_commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BahnoCommand(update tgbotapi.Update, b *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ahoj")
	_, err := b.Send(msg)
	return err
}
