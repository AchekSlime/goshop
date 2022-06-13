package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b Bot) story(message *tgbotapi.Message) error {
	userId := message.From.ID

	sId, dates, err := b.storage.StoryRepositoryInterface.GetAllStoryId(userId)
	if err != nil {
		return err
	}
	for i, v := range sId {
		layout := "2006-01-02"
		date := (dates[i])[:len(layout)]

		msg := tgbotapi.NewMessage(message.Chat.ID, "Заказ "+strconv.Itoa(v)+"\nДата: "+date)
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}

		pId, err := b.storage.StoryRepositoryInterface.GetProductsId(v)
		if err != nil {
			return err
		}

		for _, k := range pId {
			product, err := b.storage.ProductRepositoryInterface.GetById(k)
			if err != nil {
				return err
			}
			category, err := b.storage.CategoryRepositoryInterface.GetById(product.CategoryId)
			if err != nil {
				return err
			}
			msg, err := formCatalogItemMessage(message.Chat.ID, category.Name, *product)
			if err != nil {
				return err
			}
			msg.ReplyMarkup = nil

			if _, err := b.bot.Send(msg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Bot) order(query *tgbotapi.CallbackQuery) error {
	userId := query.From.ID
	pId, err := b.storage.CartRepositoryInterface.GetProductsId(userId)
	if err != nil {
		return err
	}

	if _, err = b.storage.StoryRepositoryInterface.Create(pId, userId); err != nil {
		return err
	}

	for _, v := range pId {
		if err := b.storage.CartRepositoryInterface.Delete(v, userId); err != nil {
			return err
		}
	}

	return nil
}
