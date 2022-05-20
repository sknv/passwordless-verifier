package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestSetVerificationChatParams_TypedID(t *testing.T) {
	type fields struct {
		id string
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
				id: id.String(),
			},
			want: id,
		},
		{
			name: "when invalid id is provided it returns an empty uuid",
			fields: fields{
				id: "invalid-uuid",
			},
			want: uuid.UUID{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := SetVerificationChatParams{
				ID: tt.fields.id,
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
		store Store
	}
	type args struct {
		params *SetVerificationChatParams
	}

	verificationID, chatID := uuid.New(), int64(1)

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
					store: &StoreMock{
						FindVerificationByIDFunc: func(_ context.Context, id uuid.UUID) (*model.Verification, error) {
							return &model.Verification{ID: id}, nil
						},
						UpdateVerificationFunc: func(ctx context.Context, verification *model.Verification) error {
							in := &model.Verification{
								ID:     verificationID,
								ChatID: chatID,
							}

							assert.Equalf(t, in, verification, "store.UpdateVerification(%v)", verification)
							return nil
						},
					},
				}
			},
			args: args{
				params: &SetVerificationChatParams{
					ID:     verificationID.String(),
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
				Store: fields.store,
			}
			err := u.SetVerificationChat(context.Background(), tt.args.params)
			assert.Equalf(t, tt.wantErr, err != nil, "SetVerificationChat(ctx, %v)", tt.args.params)
		})
	}
}
