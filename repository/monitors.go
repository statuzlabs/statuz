package repository

import (
	"context"
	"fmt"

	"github.com/Unfield/Statuz/db"
	"github.com/Unfield/Statuz/utils"
	"github.com/jmoiron/sqlx"
)

type monitorRepo struct {
	db *sqlx.DB
}

func NewMonitorRepository(db db.Database) MonitorRepository {
	return &monitorRepo{db: db.DB()}
}

func (r *monitorRepo) Create(ctx context.Context, m *Monitor) error {
	query := `
        INSERT INTO monitors (id, name, type, url, interval_sec, degraded_thresh_ms, enabled)
        VALUES (:id ,:name, :type, :url, :interval_sec, :degraded_thresh_ms, :enabled)
        RETURNING id, created_at, updated_at;
    `

	if m.ID == "" {
		m.ID = utils.MustGenerateID()
	}

	rows, err := r.db.NamedQueryContext(ctx, query, m)
	if err != nil {
		return fmt.Errorf("insert monitor: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return err
		}
	}
	return nil
}

func (r *monitorRepo) GetByID(ctx context.Context, id string) (*Monitor, error) {
	var m Monitor
	query := `SELECT * FROM monitors WHERE id = ? LIMIT 1`
	if err := r.db.GetContext(ctx, &m, query, id); err != nil {
		return nil, fmt.Errorf("get monitor by id: %w", err)
	}
	return &m, nil
}

func (r *monitorRepo) List(ctx context.Context) ([]Monitor, error) {
	var monitors []Monitor
	query := `SELECT * FROM monitors ORDER BY created_at DESC`
	if err := r.db.SelectContext(ctx, &monitors, query); err != nil {
		return nil, fmt.Errorf("list monitors: %w", err)
	}
	return monitors, nil
}

func (r *monitorRepo) Update(ctx context.Context, m *Monitor) error {
	query := `
        UPDATE monitors
        SET name = :name,
            type = :type,
            url = :url,
            interval_sec = :interval_sec,
            degraded_thresh_ms = :degraded_thresh_ms,
            enabled = :enabled,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = :id;
    `
	_, err := r.db.NamedExecContext(ctx, query, m)
	if err != nil {
		return fmt.Errorf("update monitor: %w", err)
	}
	return nil
}

func (r *monitorRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM monitors WHERE id = ?`
	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete monitor: %w", err)
	}
	return nil
}
