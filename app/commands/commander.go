package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/igluko/go-test-bot/internal/services/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Servise
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Servise,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
