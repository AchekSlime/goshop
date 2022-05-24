package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	startResponse = "Привет, это инетрнет магазин \"GoShop\"\n" +
		"Чтобы вызвать подсказку нажмите на кнопку  \"Инструкция\"\n" +
		"чтобы отобразить ленту товаров нажмите на кнопку \"Каталог\""
)

var unknownResponse string = "Простите, я вас не понял\n" +
	"Чтобы вызвать подсказку нажмите на кнопку  \"Инструкция\""

var StartInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Инструкция", "/help"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Каталог", "showFeed"),
	),
)

var ProductInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В корзину", ""),
	),
)

var ProductNextInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В корзину", ""),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Дальше", "/next"),
	),
)
