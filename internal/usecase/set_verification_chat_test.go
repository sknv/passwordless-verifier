package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestSetVerificationChatParams_TypedID(t *testing.T) {
	type fields struct {
		ID string
	}

	id := uuid.New()

	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "when valid id is provided it returns a filled uuid",
			fields: fields{
				ID: id.String(),
			},
			want: id,
		},
		{
			name: "when invalid id is provided it returns an empty uuid",
			fields: fields{
				ID: "invalid-uuid",
			},
			want: uuid.UUID{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := SetVerificationChatParams{
				ID: tt.fields.ID,
			}
			assert.Equalf(t, tt.want, p.TypedID(), "TypedID()")
		})
	}
}

func TestSetVerificationChatParams_Validate(t *testing.T) {
	type fields struct {
		id string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "when invalid id is provided it returns an error",
			wantErr: true,
		},
		{
			name: "when valid id is provided it returns no error",
			fields: fields{
				id: uuid.New().String(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := SetVerificationChatParams{
				ID: tt.fields.id,
			}
			err := p.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}

func TestUsecase_SetVerificationChat(t *testing.T) {
	type fields struct {
		db DB
	}
	type args struct {
		params *SetVerificationChatParams
	}

	id, chatID := uuid.New(), int64(1)

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		wantErr       bool
	}{
		{
			name:          "when args are not valid it returns an error",
			prepareFields: func() *fields { return &fields{} },
			args: args{
				params: &SetVerificationChatParams{},
			},
			wantErr: true,
		},
		{
			name: "when args are valid it finds a verification, sets chat id and updates a model returning the update error",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						FindFunc: func(_ context.Context, dest any, _ string, _ ...any) error {
							verification, _ := dest.(*model.Verification)
							verification.ID = id

							return nil
						},
						UpdateFunc: func(ctx context.Context, anyVerification any, columns ...string) (sql.Result, error) {
							in := &model.Verification{ID: id}
							in.SetChatID(chatID)

							verification, _ := anyVerification.(*model.Verification)
							verification.DB, verification.UpdatedAt = nil, time.Time{} // ignore fields when compare

							assert.Equalf(t, in, verification, "db.Update(%v)", anyVerification)
							return nil, nil
						},
					},
				}
			},
			args: args{
				params: &SetVerificationChatParams{
					ID:     id.String(),
					ChatID: chatID,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare fields and args
			fields := tt.prepareFields()

			u := &Usecase{
				DB: fields.db,
			}
			err := u.SetVerificationChat(context.Background(), tt.args.params)
			assert.Equalf(t, tt.wantErr, err != nil, "SetVerificationChat(ctx, %v)", tt.args.params)
		})
	}
}
