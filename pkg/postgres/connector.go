package postgres

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Config is a connection config.
type Config struct {
	URL             string
	MaxOpenConn     int           // maximum number of open connections
	MaxConnLifetime time.Duration // maximum amount of time a connection may be reused
}

// Option configures *sql.DB.
type Option func(*sql.DB)

// Connect opens a db connection.
func Connect(config Config, options ...Option) *bun.DB {
	sqlDB := sql.OpenDB(
		pgdriver.NewConnector(pgdriver.WithDSN(config.URL)),
	)

	// Apply config
	if config.MaxOpenConn != 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConn)
		sqlDB.SetMaxIdleConns(config.MaxOpenConn)
	}

	if config.MaxConnLifetime != 0 {
		sqlDB.SetConnMaxLifetime(config.MaxConnLifetime)
	}

	// Apply options
	for _, opt := range options {
		opt(sqlDB)
	}

	db := bun.NewDB(sqlDB, pgdialect.New())
	WithLogger(db)

	return db
}
