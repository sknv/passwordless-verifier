package telegram

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) verifyContact(_ context.Context, message *tgbotapi.Message) error {
	if message.ReplyToMessage == nil {
		if err := b.reply(message, msgWrongContactShared); err != nil {
			return fmt.Errorf("reply wrong contact shared: %w", err)
		}
		return nil
	}

	// TODO: find the latest verification by chat id and set its session

	if err := b.reply(message, fmt.Sprintf(msgFormatContactVerified, b.Config.CallbackURL)); err != nil {
		return fmt.Errorf("reply contact verified: %w", err)
	}
	return nil
}
