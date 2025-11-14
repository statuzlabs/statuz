package repository

import (
	"context"
	"time"
)

type Monitor struct {
	ID             string    `db:"id"`
	Name           string    `db:"name"`
	Type           string    `db:"type"`
	URL            string    `db:"url"`
	IntervalSec    int       `db:"interval_sec"`
	DegradedThresh int       `db:"degraded_thresh_ms"`
	Enabled        bool      `db:"enabled"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type MonitorRepository interface {
	Create(ctx context.Context, m *Monitor) error
	GetByID(ctx context.Context, id string) (*Monitor, error)
	List(ctx context.Context) ([]Monitor, error)
	Update(ctx context.Context, m *Monitor) error
	Delete(ctx context.Context, id string) error
}
