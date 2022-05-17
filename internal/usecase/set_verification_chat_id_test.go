package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestUsecase_SetVerificationChatID(t *testing.T) {
	type fields struct {
		db DB
	}
	type args struct {
		verification *model.Verification
		chatID       int64
	}

	id, chatID := uuid.New(), int64(1)

	tests := []struct {
		name          string
		prepareFields func() *fields
		prepareArgs   func(*fields) args
		wantErr       bool
	}{
		{
			name: "it sets chat id and updates a model returning the update error",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						UpdateFunc: func(ctx context.Context, anyVerification any, columns ...string) (sql.Result, error) {
							in := &model.Verification{
								ID: id,
							}
							in.SetChatID(chatID)

							verification, _ := anyVerification.(*model.Verification)
							verification.DB, verification.UpdatedAt = nil, time.Time{} // ignore fields when compare

							assert.Equalf(t, in, verification, "model.Update(%v)", anyVerification)
							return nil, errors.New("any-error")
						},
					},
				}
			},
			prepareArgs: func(f *fields) args {
				return args{
					verification: &model.Verification{
						DB: f.db,

						ID: id,
					},
					chatID: chatID,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare fields and args
			fields := tt.prepareFields()
			args := tt.prepareArgs(fields)

			u := &Usecase{
				DB: fields.db,
			}
			err := u.SetVerificationChatID(context.Background(), args.verification, args.chatID)
			assert.Equalf(t, tt.wantErr, err != nil, "SetVerificationChatID(ctx, %v, %v)", args.verification, args.chatID)
		})
	}
}
