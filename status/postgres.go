package status

import (
	"context"
	"database/sql"

	"github.com/vishal1132/bootstrap-go-web/utils"
	"gorm.io/gorm"
)

type pgMonitorOption func(*Monitor)

type Monitor struct {
	name string
	conn *sql.DB
}

// Monitor for deep health checks
func NewPGMonitor(name string, conn *gorm.DB, opts ...pgMonitorOption) *Monitor {
	m := &Monitor{name: name, conn: utils.Must(conn.DB())}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Monitor) GetName() string {
	return m.name
}

func (m *Monitor) Ping(ctx context.Context) error {
	return logDependencyError("Postgres", m.name, m.conn.PingContext(ctx))
}
