package config

import (
	"net/url"
	"time"

	"github.com/alfred-landrum/fromenv"
)

type Configuration struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Host     string `json:"server" env:"PG_HOST=localhost"`
	User     string `json:"user" env:"PG_USER=postgres"`
	Password string `json:"password" env:"PG_DB=postgres"`
	Database string `json:"database" env:"PG_PASSWORD"`
	Port     string `json:"port" env:"PG_PORT=5432"`
}

func (c *Configuration) DatabaseConfiguration() *DatabaseConfiguration {
	return &c.Database
}

// ConfigurationFromEnv creates and initializes a Configuration from the current environment.
func ConfigurationFromEnv() (*Configuration, error) {
	var c Configuration
	err := fromenv.Unmarshal(&c, configOptions()...)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func configOptions() []fromenv.Option {
	return []fromenv.Option{
		fromenv.SetFunc(durationSet),
		fromenv.SetFunc(urlSet),
		fromenv.SetFunc(timeSet),
	}
}

func durationSet(t *time.Duration, s string) error {
	x, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}

func urlSet(u *url.URL, s string) error {
	x, err := url.Parse(s)
	if err != nil {
		return err
	}
	*u = *x
	return nil
}

func timeSet(t *time.Time, s string) error {
	x, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}
