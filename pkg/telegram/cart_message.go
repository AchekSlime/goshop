package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) toCart(query *tgbotapi.CallbackQuery, productId int) error {
	userId := query.From.ID
	product, err := b.storage.ProductRepositoryInterface.GetById(productId)
	if err != nil {
		return err
	}

	err = b.storage.CartRepositoryInterface.Create(product.Id, userId)
	if err != nil {
		return err
	}

	// ToDo кинуть сообщение что продукт добавден в корзину
	return nil
}

func (b *Bot) cartMessage(message *tgbotapi.Message) error {
	userId := message.From.ID
	pId, err := b.storage.CartRepositoryInterface.GetProductsId(userId)
	if err != nil {
		return err
	}

	var msg tgbotapi.MessageConfig
	if len(pId) == 0 {
		msg = tgbotapi.NewMessage(message.Chat.ID, emptyCartResponse)
		msg.ReplyMarkup = BackInlineKeyboard
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, startOfCartMessage)
	}
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	for i, v := range pId {
		product, err := b.storage.ProductRepositoryInterface.GetById(v)
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

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Удалить", deleteFromCart+" "+strconv.Itoa(product.Id)),
			),
		)

		if i == len(pId)-1 {
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Удалить", deleteFromCart+" "+strconv.Itoa(product.Id)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Заказать", order),
				),
			)
		}

		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) deleteFromCart(query *tgbotapi.CallbackQuery, productId int) error {
	userId := query.From.ID
	if err := b.storage.CartRepositoryInterface.Delete(productId, userId); err != nil {
		return err
	}

	b.deleteMessage(query.Message.Chat.ID, query.Message.MessageID)

	//product, err := b.storage.ProductRepositoryInterface.GetById(productId)
	//if err != nil {
	//	return err
	//}
	//msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Товар\n\n\""+product.Title+"\"\n\n"+cartItemDeleted)
	//if _, err := b.bot.Send(msg); err != nil {
	//	return err
	//}

	return nil
}
