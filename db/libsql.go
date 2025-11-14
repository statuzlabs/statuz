package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type LibSQLDB struct {
	db             *sqlx.DB
	dataSourceName string
}

func NewLibSQLDatabase(dataSourceName string) *LibSQLDB {
	return &LibSQLDB{
		dataSourceName: dataSourceName,
	}
}

func (l *LibSQLDB) Open() error {
	newDB, err := sqlx.Open("libsql", l.dataSourceName)
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

func (l *LibSQLDB) DB() *sqlx.DB {
	return l.db
}

func (l *LibSQLDB) Ping() error {
	if l.db == nil {
		return fmt.Errorf("database is not open")
	}
	return l.db.Ping()
}

func (l *LibSQLDB) Close() error {
	if l.db == nil {
		return nil
	}
	return l.db.Close()
}
