package model

import (
	"context"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func TestVerifications_FindByID(t *testing.T) {
	type fields struct {
		db *bun.DB
	}
	type args struct {
		id uuid.UUID
	}

	id := uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          func(*fields) *Verification
		wantErr       bool
	}{
		{
			name: "when db returns an error it returns an error",
			prepareFields: func() *fields {
				selectQuery := &bun.SelectQuery{}
				gomonkey.ApplyMethod(selectQuery, "Model", func(*bun.SelectQuery, interface{}) *bun.SelectQuery { return selectQuery })
				gomonkey.ApplyMethod(selectQuery, "Where", func(*bun.SelectQuery, string, ...interface{}) *bun.SelectQuery { return selectQuery })
				gomonkey.ApplyMethod(selectQuery, "Scan", func(*bun.SelectQuery, context.Context, ...interface{}) error { return errors.New("any-error") })

				db := &bun.DB{}
				gomonkey.ApplyMethod(db, "NewSelect", func(*bun.DB) *bun.SelectQuery { return selectQuery })

				return &fields{
					db: db,
				}
			},
			args: args{
				id: id,
			},
			wantErr: true,
		},
		{
			name: "when db call is ok it returns a result with db field has been set",
			prepareFields: func() *fields {
				selectQuery := &bun.SelectQuery{}
				gomonkey.ApplyMethod(selectQuery, "Model", func(_ *bun.SelectQuery, model interface{}) *bun.SelectQuery {
					verification, _ := model.(*Verification)
					verification.ID = id

					return selectQuery
				})
				gomonkey.ApplyMethod(selectQuery, "Where", func(*bun.SelectQuery, string, ...interface{}) *bun.SelectQuery { return selectQuery })
				gomonkey.ApplyMethod(selectQuery, "Scan", func(*bun.SelectQuery, context.Context, ...interface{}) error { return nil })

				db := &bun.DB{}
				gomonkey.ApplyMethod(db, "NewSelect", func(*bun.DB) *bun.SelectQuery { return selectQuery })

				return &fields{
					db: db,
				}
			},
			args: args{
				id: id,
			},
			want: func(f *fields) *Verification {
				return &Verification{
					DB: f.db,

					ID: id,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare fields
			fields := tt.prepareFields()

			v := &Verifications{
				DB: fields.db,
			}

			got, err := v.FindByID(context.Background(), tt.args.id)
			assert.Equalf(t, tt.wantErr, err != nil, "FindByID(ctxx, %v)", tt.args.id)
			if tt.want != nil {
				assert.Equalf(t, tt.want(fields), got, "FindByID(ctx, %v)", tt.args.id)
			}
		})
	}
}
