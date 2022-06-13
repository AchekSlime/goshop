package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) startMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, startResponse)
	msg.ReplyMarkup = MenuKeyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) catalogMessage(message *tgbotapi.Message) error {
	categories, err := b.storage.CategoryRepositoryInterface.GetAll()
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, categoriesResponse)
	keyboard := tgbotapi.NewInlineKeyboardMarkup()
	for _, v := range categories {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(v.Name, next+" "+strconv.Itoa(v.Id)+" 0"),
		))
	}
	msg.ReplyMarkup = keyboard

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (b *Bot) unknownMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, unknownResponse)

	_, err := b.bot.Send(msg)
	return err

}
