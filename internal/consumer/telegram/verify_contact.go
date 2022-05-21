package telegram

import (
	"context"
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

func (b *Bot) VerifyContact(ctx context.Context, message *tgbotapi.Message) error {
	ctx, span := otel.Tracer("").Start(ctx, "telegram.VerifyContact")
	defer span.End()

	verification, err := b.Usecase.VerifyContact(ctx, &usecase.VerifyContactParams{
		ChatID:      message.Chat.ID,
		ContactID:   message.Contact.UserID,
		PhoneNumber: message.Contact.PhoneNumber,
	})
	if err != nil {
		err = fmt.Errorf("verify contact: %w", err)
		if errors.Is(err, usecase.ErrWrongContact) {
			if replyErr := b.reply(message, msgWrongContactShared); replyErr != nil {
				replyErr = fmt.Errorf("reply wrong contact shared: %w", replyErr)
				err = multierror.Append(err, replyErr)
			}
		}

		return err
	}

	if err = b.replyContactVerified(message, verification); err != nil {
		return fmt.Errorf("reply contact verified: %w", err)
	}
	return nil
}

func (b *Bot) replyContactVerified(to *tgbotapi.Message, verification *model.Verification) error {
	replyText := fmt.Sprintf(msgFormatContactVerified, b.formatCallbackURL(verification.ID))
	msg := tgbotapi.NewMessage(to.Chat.ID, replyText)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) formatCallbackURL(verificationID uuid.UUID) string {
	return fmt.Sprintf("%s?verificationId=%s", b.Config.CallbackURL, verificationID)
}
