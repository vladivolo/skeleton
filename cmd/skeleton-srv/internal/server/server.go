package server

import (
	"log/slog"
	"runtime/debug"
	"time"

	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/config"
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/worker"

	"github.com/vladivolo/skeleton/shared/logger"
)

type Server struct {
	http *HttpServer
	log  *logger.Logger
}

type Vcs struct {
	Revision   string
	LastCommit time.Time
	DirtyBuild bool
}

var (
	vcs *Vcs
)

func NewServer(cfg *config.Config, worker *worker.Handler, log *logger.Logger) (*Server, error) {
	var err error

	ParseVcs()

	http, err := NewHttpServer(cfg, worker, log)
	if err != nil {
		return nil, err
	}

	s := &Server{
		http: http,
		log:  log.WithField("module", "server"),
	}

	s.log.Info(
		"NewServer",
		slog.Bool("success", true),
	)

	return s, nil
}

func (s *Server) Start() error {
	var err error

	s.log.Debug("Start")

	err = s.http.Start()
	if err != nil {
		return err
	}

	s.log.Info(
		"Start",
		slog.Bool("success", true),
	)

	return nil
}

func (s *Server) Stop() {
	s.http.Stop()
}

func ParseVcs() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	vcs = &Vcs{}

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			vcs.Revision = kv.Value
		case "vcs.time":
			vcs.LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			vcs.DirtyBuild = kv.Value == "true"
		}
	}
}
