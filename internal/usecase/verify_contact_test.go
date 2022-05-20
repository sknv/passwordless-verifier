package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestUsecase_VerifyContact(t *testing.T) {
	type fields struct {
		store Store
	}
	type args struct {
		params *VerifyContactParams
	}

	verificationID, chatID, phone := uuid.New(), int64(1), "+79001002030"

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          *model.Verification
		wantErr       bool
	}{
		{
			name:          "when wrong contact provided it returns an error",
			prepareFields: func() *fields { return &fields{} },
			args: args{
				params: &VerifyContactParams{
					ChatID:    chatID,
					ContactID: chatID + 1,
				},
			},
			wantErr: true,
		},
		{
			name: "when args are valid it finds a verification, sets phone number, updates it and create its session returning an update error",
			prepareFields: func() *fields {
				return &fields{
					store: &StoreMock{
						FindLatestVerificationByChatIDFunc: func(_ context.Context, chatID int64) (*model.Verification, error) {
							return &model.Verification{
								ID:     verificationID,
								ChatID: chatID,
							}, nil
						},
						UpdateVerificationAndCreateSessionFunc: func(ctx context.Context, verification *model.Verification) error {
							in := &model.Verification{
								ID:     verificationID,
								ChatID: chatID,
								Status: model.VerificationStatusCompleted,
								Session: &model.Session{
									VerificationID: verificationID,
									PhoneNumber:    phone,
								},
							}

							verification.Session.ID = uuid.UUID{} // ignore fields when compare

							assert.Equalf(t, in, verification, "store.UpdateVerificationAndCreateSession(%v)", verification)
							return nil
						},
					},
				}
			},
			args: args{
				params: &VerifyContactParams{
					ChatID:      chatID,
					ContactID:   chatID,
					PhoneNumber: phone,
				},
			},
			want: &model.Verification{
				ID:     verificationID,
				ChatID: chatID,
				Status: model.VerificationStatusCompleted,
				Session: &model.Session{
					VerificationID: verificationID,
					PhoneNumber:    phone,
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
			got, err := u.VerifyContact(context.Background(), tt.args.params)
			assert.Equalf(t, tt.wantErr, err != nil, "VerifyContact(ctx, %v)", tt.args.params)
			assert.Equalf(t, tt.want, got, "VerifyContact(ctx, %v)", tt.args.params)
		})
	}
}
