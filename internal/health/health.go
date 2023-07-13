package health

import (
	mysqldb "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"sync"
)

type IHealthCheck interface {
	Check() error
	IsReady() bool
}

type healthCheckImpl struct {
	db      mysqldb.IMySqlDB
	logger  logging.Logger
	mu      sync.Mutex
	isReady bool
}

func NewHealthCheck(db mysqldb.IMySqlDB, logger logging.Logger) IHealthCheck {
	return &healthCheckImpl{
		db:      db,
		logger:  logger,
		isReady: false,
	}
}

func (h *healthCheckImpl) Check() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if err := h.db.Ping(); err != nil {
		h.isReady = false
		h.logger.Errorf("HealthCheck db failed: %s", err.Error())
		return err
	}

	h.isReady = true
	return nil
}

func (h *healthCheckImpl) IsReady() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.isReady
}
