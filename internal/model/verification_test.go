package model

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func TestVerificationMethod_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       VerificationMethod
		wantErr bool
	}{
		{
			name:    "when allowed method is provided it returns no error",
			m:       VerificationMethodTelegram,
			wantErr: false,
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
		db *bun.DB
	}

	tests := []struct {
		name          string
		prepareFields func() *fields
		wantErr       bool
	}{
		{
			name: "it returns db call result",
			prepareFields: func() *fields {
				insertQuery := &bun.InsertQuery{}
				gomonkey.ApplyMethod(insertQuery, "Model", func(*bun.InsertQuery, interface{}) *bun.InsertQuery { return insertQuery })
				gomonkey.ApplyMethod(insertQuery, "Exec", func(*bun.InsertQuery, context.Context, ...interface{}) (sql.Result, error) {
					return nil, errors.New("any-error")
				})

				db := &bun.DB{}
				gomonkey.ApplyMethod(db, "NewInsert", func(*bun.DB) *bun.InsertQuery { return insertQuery })

				return &fields{
					db: db,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare fields
			fields := tt.prepareFields()

			v := &Verification{
				DB: fields.db,
			}
			err := v.Create(context.Background())
			assert.Equalf(t, tt.wantErr, err != nil, "Create()")
		})
	}
}
