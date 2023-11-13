package telegrambot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Notify(data int64) error {
	bot, err := tgbotapi.NewBotAPI("6445599205:AAGPIKb-pUOMD8z8un5x4usZuUjXA7BlOrc")
	if err != nil {
		return err
	}

	message := ``
	message += fmt.Sprintf("Order: %v\n", data)

	_, err = bot.Send(tgbotapi.NewMessage(-4095952517, message))
	return err
}
