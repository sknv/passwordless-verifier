package telegram

import (
	"context"
	"fmt"

	"github.com/google/uuid"

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

	replyText := fmt.Sprintf(msgFormatContactVerified, b.Config.CallbackURL)
	if err := b.reply(message, replyText); err != nil {
		return fmt.Errorf("reply contact verified: %w", err)
	}
	return nil
}

//nolint:unused // TODO: remove
func (b *Bot) formatCallbackURL(verificationID uuid.UUID) string {
	return fmt.Sprintf("%s?verificationId=%s", b.Config.CallbackURL, verificationID)
}
