package db

import "github.com/jmoiron/sqlx"

type Database interface {
	Open() error
	Close() error
	DB() *sqlx.DB
	Ping() error
	Migrate(schema string) error
}
