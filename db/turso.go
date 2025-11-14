package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/turso-go"
)

type TursoDB struct {
	db             *sqlx.DB
	dataSourceName string
}

func NewTursoDatabase(dataSourceName string) *TursoDB {
	return &TursoDB{
		dataSourceName: dataSourceName,
	}
}

func (l *TursoDB) Open() error {
	newDB, err := sqlx.Open("turso", l.dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to open libSQL DB: %w", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := newDB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping libSQL DB: %w", err)
	}
	l.db = newDB
	return nil
}

func (l *TursoDB) DB() *sqlx.DB {
	return l.db
}

func (l *TursoDB) Ping() error {
	if l.db == nil {
		return fmt.Errorf("database is not open")
	}
	return l.db.Ping()
}

func (l *TursoDB) Close() error {
	if l.db == nil {
		return nil
	}
	return l.db.Close()
}
