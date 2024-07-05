package configs

import (
	"context"

	env "github.com/sethvargo/go-envconfig"
)

type Rest struct {
	ListenAddrPort string `yaml:"http_listen_addr" env:"HTTP_LISTEN_ADDR_PORT,overwrite"`
}

func NewRest() (*Rest, error) {
	r := &Rest{}

	err := r.Process()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Rest) Process() error {
	if err := r.envPatch(); err != nil {
		return err
	}

	if err := r.validate(); err != nil {
		return err
	}

	return nil
}

func (r *Rest) envPatch() error {
	return env.Process(context.Background(), r)
}

func (r *Rest) validate() error {
	return nil
}
