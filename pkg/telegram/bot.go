package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goshop/pkg/repository"
	"log"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage repository.Storage
}

func NewBot(token string, storage repository.Storage) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &Bot{
		bot:     bot,
		storage: storage,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	if err := b.Route(updates); err != nil {
		return err
	}
	return nil
}

func (b *Bot) Route(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil { // Message

			// Text Messages
			if err := b.handleMessage(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
		} else if update.CallbackQuery != nil { // CallBack
			if err := b.handleCallBack(update.CallbackQuery); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue

		} else { // Other updates
			continue
		}
	}

	return nil
}

func (b *Bot) deleteMessage(chatId int64, messageId int) {
	b.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chatId,
		MessageID: messageId,
	})
}

func (b *Bot) handleError(chatID int64, err error) {
	log.Printf("[LOG] [CHAT_ID='%d']    ...    '%s'", chatID, err.Error())
}
