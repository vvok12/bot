package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	token := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")
	b, err := bot.New(token, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	log.Default().Println(token)

	b.Start(ctx, &bot.StartParams{
		AllowedUpdates: []string{"message", "message_reaction"},
	})
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Default().Println("got some message")

	if update.Message == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
