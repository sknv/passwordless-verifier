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

		return nil
	}

	// Try to find the verification
	verification, err := b.Usecase.GetVerification(ctx, &usecase.GetVerificationParams{ID: startUUID})
	if err != nil {
		if replyErr := b.reply(message, msgStartNotFound); err != nil {
			err = multierror.Append(err, replyErr)
		}

		return err
	}

	// Save telegram chat id for the verification
	if err = b.Usecase.SetVerificationChatID(ctx, verification, message.Chat.ID); err != nil {
		return fmt.Errorf("update verification: %w", err)
	}

	// Ask a user to share their contact
	if err = b.shareContact(message); err != nil {
		return fmt.Errorf("share contact: %w", err)
	}

	return nil
}

func (b *Bot) shareContact(to *tgbotapi.Message) error {
	keyboard := tgbotapi.NewOneTimeReplyKeyboard([]tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButtonContact(btnShareContactText),
	})
	msg := tgbotapi.NewMessage(to.Chat.ID, msgShareContact)
	msg.ReplyMarkup = keyboard

	_, err := b.bot.Send(msg)
	return err
}
