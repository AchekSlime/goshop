package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case startCommand:
		return b.startMessage(message)
	case startMenu:
		return b.startMessage(message)
	case catalogMenu:
		return b.catalogMessage(message)
	default:
		return b.unknownMessage(message)
	}
}

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

//func (b *Bot) catalogMessage(message *tgbotapi.Message) error {
//	products, err := b.storage.ProductRepositoryInterface.GetFirst(10)
//	if err != nil {
//		return err
//	}
//	if len(products) == 0 {
//		msg := tgbotapi.NewMessage(message.Chat.ID, emptyCatalogResponse)
//		if _, err := b.bot.Send(msg); err != nil {
//			return err
//		}
//		return nil
//	}
//
//	for _, v := range products {
//		var msg tgbotapi.MessageConfig
//		msg = formCatalogItemMessage(message.Chat.ID, v)
//		if _, err := b.bot.Send(msg); err != nil {
//			return err
//		}
//	}
//
//	packMsg := formEndOfPackMessage(message.Chat.ID, products[len(products)-1].Id)
//	if _, err := b.bot.Send(packMsg); err != nil {
//		return err
//	}
//
//	return nil
//}

func (b *Bot) unknownMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, unknownResponse)

	_, err := b.bot.Send(msg)
	return err

}
