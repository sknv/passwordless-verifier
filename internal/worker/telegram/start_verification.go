package telegram

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hashicorp/go-multierror"

	"github.com/sknv/passwordless-verifier/internal/usecase"
)

func (b *Bot) startVerification(ctx context.Context, message *tgbotapi.Message) error {
	_, startUUID, found := strings.Cut(message.Text, "/start ")
	if !found {
		if err := b.reply(message, msgStartInitiatedDirectly); err != nil {
			return fmt.Errorf("reply uuid not found: %w", err)
		}
	}

	// Try to find the verification
	_, err := b.Usecase.GetVerification(ctx, &usecase.GetVerificationParams{ID: startUUID})
	if err != nil {
		if replyErr := b.reply(message, msgStartNotFound); err != nil {
			err = multierror.Append(err, replyErr)
		}

		return err
	}

	// TODO: save telegram chat id for the verification

	if err = b.shareContact(message); err != nil {
		return fmt.Errorf("ask share contact: %w", err)
	}

	return nil
}

func (b *Bot) shareContact(to *tgbotapi.Message) error {
	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButtonContact(btnShareContactText),
	})
	msg := tgbotapi.NewMessage(to.Chat.ID, msgAskShareContact)
	msg.ReplyMarkup = keyboard

	_, err := b.bot.Send(msg)
	return err
}
