package configs

import (
	"context"
	"fmt"

	env "github.com/sethvargo/go-envconfig"
)

type Log struct {
	Level     string `yaml:"level" env:"LOG_LEVEL,overwrite,default=error"`
	Format    string `yaml:"format" env:"LOG_FORMAT,overwrite,default=json"`
	AddSource bool   `yaml:"add_source" env:"LOG_ADD_SOURCE,overwrite,default=false"`
}

func NewLog() (*Log, error) {
	l := &Log{}

	err := l.Process()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Log) Process() error {
	if err := l.envPatch(); err != nil {
		return err
	}

	if err := l.validate(); err != nil {
		return err
	}

	return nil
}

func (l *Log) envPatch() error {
	return env.Process(context.Background(), l)
}

func (l *Log) validate() error {
	if l.Format != "text" && l.Format != "json" {
		return fmt.Errorf("Unsupported log format")
	}

	return nil
}
