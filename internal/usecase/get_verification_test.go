package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetVerificationParams_TypedID(t *testing.T) {
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

			p := GetVerificationParams{
				ID: tt.fields.ID,
			}
			assert.Equalf(t, tt.want, p.TypedID(), "TypedID()")
		})
	}
}

func TestGetVerificationParams_Validate(t *testing.T) {
	type fields struct {
		id string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "when invalid method is provided it returns an error",
			wantErr: true,
		},
		{
			name: "when valid method is provided it returns no error",
			fields: fields{
				id: uuid.New().String(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := GetVerificationParams{
				ID: tt.fields.id,
			}
			err := p.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}
