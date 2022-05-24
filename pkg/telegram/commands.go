package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	startCommand = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, startResponse)
	msg.ReplyMarkup = StartInlineKeyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, unknownResponse)
	msg.ReplyMarkup = StartInlineKeyboard

	_, err := b.bot.Send(msg)
	return err

}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	return b.handleUnknownCommand(message)
}
