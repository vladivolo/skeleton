package main

import (
	"flag"
	"runtime"

	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/config"
	"github.com/vladivolo/skeleton/cmd/skeleton-srv/internal/service"
	"github.com/vladivolo/skeleton/shared/execute"
)

var (
	configFile = flag.String("config", "./conf/config.yaml", "/path/to/config.yaml")
)

func main() {
	flag.Parse()

	// Load config
	conf, err := config.New(*configFile)
	if err != nil {
		panic(err)
	}

	// Set GOMAXPROCS
	runtime.GOMAXPROCS(conf.System.GoMaxProcs)

	// Start service
	s := service.New(conf)
	err = execute.StartService(s)
	if err != nil {
		panic(err)
	}

	<-make(chan int)
}
