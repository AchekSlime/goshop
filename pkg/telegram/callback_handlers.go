package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (b *Bot) handleCallBack(query *tgbotapi.CallbackQuery) error {
	callback := tgbotapi.NewCallback(query.ID, query.Data)
	if _, err := b.bot.AnswerCallbackQuery(callback); err != nil {
		panic(err)
	}

	splitCommand := strings.Split(query.Data, " ")
	command := splitCommand[0]

	switch command {
	case next:
		categoryId, _ := strconv.Atoi(splitCommand[1])
		lastId, _ := strconv.Atoi(splitCommand[2])
		return b.nextInCatalog(query, categoryId, lastId)
	case backToCatalog:
		return b.backToCatalog(query)
	default:
		return nil
	}
}

func (b *Bot) backToCatalog(query *tgbotapi.CallbackQuery) error {
	b.deleteMessage(query.Message.Chat.ID, query.Message.MessageID)
	err := b.catalogMessage(query.Message)
	return err
}

func (b *Bot) nextInCatalog(query *tgbotapi.CallbackQuery, categoryId, lastId int) error {
	category, err := b.storage.CategoryRepositoryInterface.GetById(categoryId)
	if err != nil {
		return err
	}

	if lastId == 0 {
		startCatalogMsg := tgbotapi.NewMessage(query.Message.Chat.ID, startOfCatalogMessage+category.Name+":")
		if _, err := b.bot.Send(startCatalogMsg); err != nil {
			return err
		}
	}

	products, err := b.storage.ProductRepositoryInterface.GetAfterByCategory(categoryId, lastId, limit)
	if err != nil {
		return err
	}

	b.deleteMessage(query.Message.Chat.ID, query.Message.MessageID)

	endMsg := tgbotapi.NewMessage(query.Message.Chat.ID, endOfCatalogResponse)
	endMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", backToCatalog),
		),
	)

	if len(products) == 0 {
		if _, err := b.bot.Send(endMsg); err != nil {
			return err
		}
		return nil
	}

	for _, v := range products {
		var msg tgbotapi.MessageConfig
		msg = formCatalogItemMessage(query.Message.Chat.ID, v)
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}

	if len(products) < limit {
		if _, err := b.bot.Send(endMsg); err != nil {
			return err
		}
	} else {
		packMsg := formEndOfPackMessage(query.Message.Chat.ID, categoryId, products[len(products)-1].Id)
		if _, err := b.bot.Send(packMsg); err != nil {
			return err
		}
	}
	return nil
}
