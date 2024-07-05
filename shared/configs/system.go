package configs

import (
	"context"

	env "github.com/sethvargo/go-envconfig"
)

type System struct {
	GoMaxProcs int `yaml:"gomaxprocs" env:"GOMAXPROCS,overwrite,default=0"`
}

func NewSystem() (*System, error) {
	s := &System{}

	err := s.Process()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *System) Process() error {
	if err := s.envPatch(); err != nil {
		return err
	}

	if err := s.validate(); err != nil {
		return err
	}

	return nil
}

func (s *System) envPatch() error {
	return env.Process(context.Background(), s)
}

func (s *System) validate() error {
	return nil
}
