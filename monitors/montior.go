package monitors

import (
	"context"
	"time"
)

type Status string

const (
	StatusUp       Status = "up"
	StatusDown     Status = "down"
	StatusDegraded Status = "degraded"
	StatusUnknown  Status = "unknown"
	StatusWarning  Status = "warning"
)

type Result struct {
	Type      string        `json:"type"`
	Status    Status        `json:"status"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
	Message   string        `json:"message,omitempty"`
	Error     string        `json:"error,omitempty"`
	Success   bool          `json:"success"`
	CheckedAt time.Time     `json:"checked_at"`
	Details   any           `json:"details,omitempty"`
}

type Monitor interface {
	Check(ctx context.Context) Result
}
