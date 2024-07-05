package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"

	"github.com/vladivolo/skeleton/shared/configs"
)

type Config struct {
	configs.Rest   `yaml:"rest"`
	configs.Log    `yaml:"log"`
	configs.Db     `yaml:"db"`
	configs.System `yaml:"system"`
}

type ConfigItem interface {
	Process() error
}

func New(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	Elements := reflect.ValueOf(cfg).Elem()
	for i := 0; i < Elements.NumField(); i++ {
		Element := Elements.Field(i)
		if Element.Kind() != reflect.Struct {
			continue
		}

		if ok := Element.CanAddr(); !ok {
			continue
		}

		NestedConfig, ok := Element.Addr().Interface().(ConfigItem)
		if !ok {
			continue
		}

		err := NestedConfig.Process()
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
