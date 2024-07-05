package service

import (
	"log/slog"

	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/config"
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/server"
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/worker"

	"github.com/vladivolo/skeleton/shared/execute"
	"github.com/vladivolo/skeleton/shared/logger"
)

type service struct {
	cfg        *config.Config
	s          *server.Server
	global_log *logger.Logger
	log        *logger.Logger
}

func New(conf *config.Config) execute.Service {
	logger := logger.New(conf.Log)

	s := service{
		cfg: conf,
	}

	s.global_log = logger.WithField("srv", s.Name())
	s.log = s.global_log.WithField("module", "service")

	s.log.Debug("new")

	return &s
}

func (s *service) Name() string {
	return "skeleton-srv"
}

func (s *service) Init() error {
	s.log.Debug("Init")

	lh, err := worker.NewHandler(s.cfg, s.global_log)
	if err != nil {
		s.log.Error(
			"Init",
			slog.String("error", err.Error()),
		)
		return err
	}

	srv, err := server.NewServer(s.cfg, lh, s.global_log)
	if err != nil {
		s.log.Error(
			"Init",
			slog.String("error", err.Error()),
		)
		return err
	}

	s.s = srv

	return nil
}

func (s *service) Start() error {
	s.log.Debug("start")

	s.s.Start()

	return nil
}

func (s *service) Stop() error {
	s.log.Debug("stop")

	s.s.Stop()

	return nil
}
