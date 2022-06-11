package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goshop/pkg/entities"
	"strconv"
)

const (
	limit = 5
)

const (
	startResponse = "Привет, это инетрнет магазин \"GoShop\"\n" +
		"Чтобы вызвать подсказку нажмите на кнопку  \"Инструкция\"\n" +
		"чтобы отобразить ленту товаров нажмите на кнопку \"Каталог\""

	unknownResponse = "Простите, я вас не понял\n" +
		"Чтобы вызвать подсказку нажмите на кнопку  \"Инструкция\""

	categoriesResponse = "Выберите раздел, чтобы вывести список товаров:"

	emptyCatalogResponse  = "Каталог пуст"
	endOfCatalogResponse  = "Конец каталога"
	startOfCatalogMessage = "Товары в разделе "
)

var endOfPackResponse string = "Показано " + strconv.Itoa(limit) + " товаров"

const (
	startCommand = "/start"
	startMenu    = "Начало"
	catalogMenu  = "Каталог"
)

const (
	category      = "category"
	next          = "next"
	toCart        = "toCart"
	backToCatalog = "backToCatalog"
)

var MenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Начало"),
		tgbotapi.NewKeyboardButton("Каталог"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Корзина"),
		tgbotapi.NewKeyboardButton("Заказы"),
		tgbotapi.NewKeyboardButton("Настройки"),
	),
)

var ProductInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В корзину", toCart),
	),
)

func formCatalogItemMessage(chatId int64, product entities.Product) tgbotapi.MessageConfig {
	response := strconv.Itoa(product.Id) + "\n" + product.Title + "\n Цена: " + strconv.Itoa(product.Price)
	msg := tgbotapi.NewMessage(chatId, response)
	msg.ReplyMarkup = ProductInlineKeyboard
	return msg
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
