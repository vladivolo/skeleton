package configs

import (
	"context"
	"fmt"

	env "github.com/sethvargo/go-envconfig"
)

type Db struct {
	Driver       string `yaml:"driver" env:"DB_DRIVER,overwrite,default=postgres"`
	Host         string `yaml:"host" env:"DB_HOST,overwrite,default=localhost"`
	Port         int    `yaml:"port" env:"DB_PORT,overwrite,default=5432"`
	Name         string `yaml:"name" env:"DB_NAME,overwrite"`
	User         string `yaml:"user" env:"DB_USER,overwrite"`
	Password     string `yaml:"password" env:"DB_PASSWORD,overwrite"`
	SSLMode      string `yaml:"sslmode" env:"DB_SSL,overwrite,default=enable"`
	MaxConns     int    `yaml:"max_conns" env:"DB_MAX_CONN,overwrite"`
	MaxIdleConns int    `yaml:"max_idle_conns" env:"DB_IDLE_CONN,overwrite"`
}

func NewDb() (*Db, error) {
	d := &Db{}
	if err := d.Process(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Db) Process() error {
	err := d.envPatch()
	if err != nil {
		return err
	}

	err = d.validate()
	if err != nil {
		return err
	}

	return nil
}

func (d *Db) envPatch() error {
	return env.Process(context.Background(), d)
}

func (d *Db) validate() error {
	if d.Driver == "" {
		return fmt.Errorf("Config:Database unknown database.driver")
	}

	if d.Name == "" {
		return fmt.Errorf("Config:Database unknown database.name")
	}

	if d.User == "" {
		return fmt.Errorf("Config:Database unknown database.user")
	}

	if d.Password == "" {
		return fmt.Errorf("Config:Database unknown database.password")
	}

	return nil
}

func (d *Db) ConnString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%d",
		d.User,
		d.Password,
		d.Name,
		d.SSLMode,
		d.Host,
		d.Port,
	)
}

func (d *Db) ConnURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		d.Driver, d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode)
}
