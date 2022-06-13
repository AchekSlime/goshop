package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	emptyCartResponse     = "В корзине пусто ;("
	startOfCartMessage    = "Ваша корзина:"
	cartItemDeleted       = "Удален из вашей корзины"
)

var endOfPackResponse string = "Показано " + strconv.Itoa(limit) + " товаров"

const (
	startCommand = "/start"
	startMenu    = "Начало"
	catalogMenu  = "Каталог"
	cartMenu     = "Корзина"
	storyMenu    = "Заказы"
)

const (
	next           = "next"
	toCart         = "toCart"
	backToCatalog  = "backToCatalog"
	deleteFromCart = "deleteFromCart"
	order          = "order"
)

const (
	imageBasePath = "images/"
)

var MenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Начало"),
		tgbotapi.NewKeyboardButton("Каталог"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Корзина"),
		tgbotapi.NewKeyboardButton("Заказы"),
	),
)

var ProductInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В корзину", toCart),
	),
)

var BackInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", backToCatalog),
	),
)
