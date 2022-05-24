package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goshop/pkg/repository"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage repository.Repository
}

func NewBot(token string, storage repository.Repository) *Bot {
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

	for update := range updates {
		if update.Message != nil { // Message

			// Commands
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.handleError(update.Message.Chat.ID, err)
				}

				continue
			}

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
