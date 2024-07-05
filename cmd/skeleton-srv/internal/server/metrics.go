package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	grpcGetRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_get_tokens_request",
		Help: "Counter of grpc get tokens",
	})

	httpGetRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_get_tokens_request",
		Help: "Counter of http get tokens",
	})

	grpcRefreshRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_refresh_tokens_request",
		Help: "Counter of grpc refresh tokens",
	})

	httpRefreshRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_refresh_tokens_request",
		Help: "Counter of http refresh tokens",
	})

	grpcDenyRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_deny_users_request",
		Help: "Counter of grpc/deny users request",
	})

	httpDenyRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_deny_users_request",
		Help: "Counter of http/deny users request",
	})
)

func GrpcGetRequest() {
	grpcGetRequestCounter.Inc()
}

func HttpGetRequest() {
	httpGetRequestCounter.Inc()
}

func GrpcRefreshRequest() {
	grpcRefreshRequestCounter.Inc()
}

func HttpRefreshRequest() {
	httpRefreshRequestCounter.Inc()
}

func GrpcDenyRequest() {
	grpcDenyRequestCounter.Inc()
}

func HttpDenyRequest() {
	httpDenyRequestCounter.Inc()
}
