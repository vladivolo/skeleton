package worker

import (
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/config"

	"github.com/vladivolo/skeleton/shared/logger"
)

type Handler struct {
}

func NewHandler(conf *config.Config, log *logger.Logger) (*Handler, error) {
	return nil, nil
}
