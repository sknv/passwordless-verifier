package telegram

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hashicorp/go-multierror"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/usecase"
)

func (b *Bot) StartVerification(ctx context.Context, message *tgbotapi.Message) error {
	ctx, span := otel.Tracer("").Start(ctx, "telegram.StartVerification")
	defer span.End()

	_, startUUID, found := strings.Cut(message.Text, "/start ")
	if !found {
		if err := b.reply(message, msgStartInitiatedDirectly); err != nil {
			return fmt.Errorf("reply uuid not found: %w", err)
		}
		return nil
	}

	// Save telegram chat id for the verification
	err := b.Usecase.SetVerificationChat(ctx, &usecase.SetVerificationChatParams{
		ID:     startUUID,
		ChatID: message.Chat.ID,
	})
	if err != nil {
		err = fmt.Errorf("set verification chat: %w", err)
		if replyErr := b.reply(message, msgStartNotFound); replyErr != nil {
			replyErr = fmt.Errorf("reply start not found: %w", replyErr)
			err = multierror.Append(err, replyErr)
		}

		return err
	}

	// Ask a user to share their contact
	if err = b.replyShareContact(message); err != nil {
		return fmt.Errorf("reply share contact: %w", err)
	}
	return nil
}

func (b *Bot) replyShareContact(to *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(to.Chat.ID, msgShareContact)
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard([]tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButtonContact(btnShareContactText),
	})

	_, err := b.bot.Send(msg)
	return err
}
