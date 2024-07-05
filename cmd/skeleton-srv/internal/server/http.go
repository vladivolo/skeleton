package server

import (
	"log/slog"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/config"
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/worker"

	"github.com/vladivolo/skeleton/shared/logger"
)

type HttpServer struct {
	app    *fiber.App
	worker *worker.Handler

	cfg *config.Config
	log *logger.Logger
}

func NewHttpServer(cfg *config.Config, worker *worker.Handler, log *logger.Logger) (*HttpServer, error) {
	hs := &HttpServer{
		app:    fiber.New(),
		worker: worker,
		cfg:    cfg,
		log:    log.WithField("module", "httpserver"),
	}

	hs.app.Use(cors.New(cors.Config{
		// attempt to mitigate CORS issues - pay attention to last /
		//AllowOrigins: "http://localhost:8080, http://localhost:8080/,", //
		//AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		Next:         nil,
	}))

	hs.initMetrics()

	hs.setRoutes()

	hs.log.Info(
		"NewHttpServer",
		slog.String("http_listen_addr", cfg.Rest.ListenAddrPort),
		slog.Bool("success", true),
	)

	return hs, nil
}

func (hs *HttpServer) Start() error {
	hs.log.Debug("start")

	// Listen from a different goroutine
	go func() {
		if err := hs.app.Listen(hs.cfg.Rest.ListenAddrPort); err != nil {
			hs.log.Error(
				"listen",
				slog.String("addr", hs.cfg.Rest.ListenAddrPort),
				slog.String("error", err.Error()),
			)
		}
	}()

	return nil
}

func (hs *HttpServer) Stop() {
}

func (hs *HttpServer) setRoutes() {
	// Ping/Pong
	hs.app.Get("/ping", hs.PingHandler)

	// Version
	hs.app.Get("/version", hs.VersionHandler)

	// readiness
	hs.app.Get("/readiness", hs.ReadinessHandler)

	// liveliness
	hs.app.Get("/liveliness", hs.LivelinessHandler)
}

func (hs *HttpServer) initMetrics() {
	prometheus := fiberprometheus.New("skeleton-srv")
	prometheus.RegisterAt(hs.app, "/metrics")
	hs.app.Use(prometheus.Middleware)
}

func (hs *HttpServer) PingHandler(ctx *fiber.Ctx) error {
	hs.log.Debug("Ping")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Pong",
	})
}

func (hs *HttpServer) VersionHandler(ctx *fiber.Ctx) error {
	hs.log.Debug("Version")

	if vcs == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"revision":    vcs.Revision,
		"last_commit": vcs.LastCommit,
		"dirty_build": vcs.DirtyBuild,
	})
}

func (hs *HttpServer) ReadinessHandler(ctx *fiber.Ctx) error {
	hs.log.Debug("Readiness")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (hs *HttpServer) LivelinessHandler(ctx *fiber.Ctx) error {
	hs.log.Debug("Liveliness")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}
