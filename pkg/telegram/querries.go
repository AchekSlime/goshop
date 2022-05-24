package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goshop"
	"strconv"
)

const (
	showFeed = "showFeed"
	next     = "showFeed"
)

func (b *Bot) handleCallBack(query *tgbotapi.CallbackQuery) error {
	switch query.Data {
	case showFeed:
		return b.handleShowFeedQuery(query)
	default:
		return b.handleUnknownQuery(query)
	}
}

func (b *Bot) handleShowFeedQuery(query *tgbotapi.CallbackQuery) error {
	// получить список товаров из бд
	products, err := b.storage.ProductRepository.GetFirst(10)
	if err != nil {
		return err
	}

	// ToDo зачем?
	callback := tgbotapi.NewCallback(query.ID, query.Data)
	if _, err := b.bot.AnswerCallbackQuery(callback); err != nil {
		panic(err)
	}

	// отправить каждый товар отдельным сообщением
	for i, v := range products {
		msg := formMessage(v, query)
		if i == len(products)-1 {
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("В корзину", ""),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Дальше", "/next/"+strconv.Itoa(v.Id)),
				),
			)
		}
		// ToDo обработать ошибку
		b.bot.Send(msg)
	}

	return nil
}

func formMessage(product goshop.Product, query *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	response := strconv.Itoa(product.Id) + "\n" + product.Title + "\n Цена: " + strconv.Itoa(product.Price)
	msg := tgbotapi.NewMessage(query.Message.Chat.ID, response)
	msg.ReplyMarkup = ProductInlineKeyboard.InlineKeyboard
	return msg
}

func (b *Bot) handleUnknownQuery(query *tgbotapi.CallbackQuery) error {
	//TODO implement me
	panic("implement me")
}
