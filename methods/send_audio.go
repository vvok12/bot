package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#sendaudio

type SendAudioParams struct {
	ChatID                   string                 `json:"chat_id"`
	Audio                    models.InputFile       `json:"audio"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	Performer                string                 `json:"performer,omitempty"`
	Title                    string                 `json:"title,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

func (p SendAudioParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.Audio == nil {
		return bot.ErrEmptyAudio
	}
	return nil
}

func SendAudio(ctx context.Context, b *bot.Bot, params *SendAudioParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendAudio", params, result)

	return result, err
}