package usecase

import (
	"context"

	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func (u *Usecase) SetVerificationChatID(ctx context.Context, verification *model.Verification, chatID int64) error {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.SetVerificationChatID")
	defer span.End()

	verification.SetChatID(chatID)
	return verification.Update(ctx)
}
