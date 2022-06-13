package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goshop/pkg/entities"
	"io/ioutil"
	"strconv"
)

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
		//image, err := ReadImage(category.Name + "/" + imageMenuPath + ".jpeg")
		//if err != nil {
		//	return err
		//}
		//
		//msg := tgbotapi.NewPhotoUpload(query.Message.Chat.ID, *image)
		//msg.Caption = startOfCatalogMessage + category.Name + ":"
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, startOfCatalogMessage+category.Name+":")
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}

	products, err := b.storage.ProductRepositoryInterface.GetAfterByCategory(categoryId, lastId, limit)
	if err != nil {
		return err
	}

	b.deleteMessage(query.Message.Chat.ID, query.Message.MessageID)

	endMsg := tgbotapi.NewMessage(query.Message.Chat.ID, endOfCatalogResponse)
	endMsg.ReplyMarkup = BackInlineKeyboard

	if len(products) == 0 {
		if _, err := b.bot.Send(endMsg); err != nil {
			return err
		}
		return nil
	}

	for _, v := range products {
		msg, err := formCatalogItemMessage(query.Message.Chat.ID, category.Name, v)
		if err != nil {
			return err
		}
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

func formCatalogItemMessage(chatId int64, categoryName string, product entities.Product) (*tgbotapi.PhotoConfig, error) {
	//response := strconv.Itoa(product.Id) + "\n" + product.Title + "\n\nЦена: " + strconv.Itoa(product.Price/1000) +
	//	" " + strconv.Itoa(product.Price%1000) + "₽"

	price := ""
	if product.Price > 1000 {
		price = strconv.Itoa(product.Price / 1000)
	}
	sec := strconv.Itoa(product.Price % 1000)
	for len(sec) < 3 {
		sec = "0" + sec
	}
	price = price + " " + sec

	response := product.Title + "\n\nЦена: " + price + "₽"

	//path := strings.Replace(product.Title, " ", "\\ ", -1)
	path := product.Title

	image, err := ReadImage(categoryName + "/" + path + ".jpeg")
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewPhotoUpload(chatId, *image)
	msg.Caption = response
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В корзину", toCart+" "+strconv.Itoa(product.Id)),
		),
	)
	return &msg, nil
}

func ReadImage(path string) (*tgbotapi.FileBytes, error) {
	path = imageBasePath + path
	photoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	return &photoFileBytes, nil
}

func formEndOfPackMessage(chatId int64, categoryId, lastId int) tgbotapi.MessageConfig {
	endMsg := tgbotapi.NewMessage(chatId, endOfPackResponse)
	endMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", backToCatalog),
			tgbotapi.NewInlineKeyboardButtonData("Показать еще", next+" "+strconv.Itoa(categoryId)+" "+strconv.Itoa(lastId)),
		),
	)
	return endMsg
}
