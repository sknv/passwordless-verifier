package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVerificationMethod_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       VerificationMethod
		wantErr bool
	}{
		{
			name: "when allowed method is provided it returns no error",
			m:    VerificationMethodTelegram,
		},
		{
			name:    "when unknown method is provided it returns an error",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.m.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}

func TestVerification_Create(t *testing.T) {
	type fields struct {
		db DB
		id uuid.UUID
	}
	type args struct {
		deeplinkFormat string
	}

	deeplinkFormat, id := "https://t.me/example_bot?start=%s", uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          func(*fields) *Verification
		wantErr       bool
	}{
		{
			name: "it creates deeplink, set status and returns db call result",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						CreateFunc: func(context.Context, any) (sql.Result, error) { return nil, errors.New("any-error") },
					},
					id: id,
				}
			},
			args: args{
				deeplinkFormat: deeplinkFormat,
			},
			want: func(f *fields) *Verification {
				return &Verification{
					DB: f.db,

					ID:       id,
					Status:   VerificationStatusInProgress,
					Deeplink: fmt.Sprintf(deeplinkFormat, id),
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare fields
			fields := tt.prepareFields()

			v := &Verification{
				DB: fields.db,

				ID: id,
			}
			err := v.Create(context.Background(), tt.args.deeplinkFormat)
			assert.Equalf(t, tt.wantErr, err != nil, "Create(ctx, %s)", tt.args.deeplinkFormat)
			if tt.want != nil {
				assert.Equalf(t, tt.want(fields), v, "Create(ctx, %v)", tt.args.deeplinkFormat)
			}
		})
	}
}

func TestVerification_Update(t *testing.T) {
	type fields struct {
		db DB
	}

	tests := []struct {
		name          string
		prepareFields func() *fields
		wantErr       bool
	}{
		{
			name: "it returns db call result",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						UpdateFunc: func(context.Context, any, ...string) (sql.Result, error) { return nil, errors.New("any-error") },
					},
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare fields
			fields := tt.prepareFields()

			v := &Verification{
				DB: fields.db,
			}
			err := v.Update(context.Background())
			assert.Equalf(t, tt.wantErr, err != nil, "Update(ctx)")
		})
	}
}

func TestVerification_SetChatID(t *testing.T) {
	type args struct {
		chatID int64
	}

	chatID := int64(1)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "it sets chat id sql.NullInt64 value",
			args: args{
				chatID: chatID,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := &Verification{}
			v.SetChatID(tt.args.chatID)
			assert.Equalf(t, chatID, v.ChatID.Int64, "SetChatID(%v)", tt.args.chatID)
			assert.Truef(t, v.ChatID.Valid, "SetChatID(%v)", tt.args.chatID)
		})
	}
}
