package config

import (
	"os"
)

var ConfigManager *Configuration

type Configuration struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Host     string `json:"server"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

func loadConfiguration() (c *Configuration) {
	c = &Configuration{}
	c.Database.Host = os.Getenv(PGHost)
	c.Database.User = os.Getenv(PGUser)
	c.Database.Password = os.Getenv(PGPwd)
	c.Database.Database = os.Getenv(PGDb)
	c.Database.Port = os.Getenv(PGPort)
	if c.Database.Port == "" {
		c.Database.Port = "5432"
	}

	return c
}

func (c *Configuration) GetDatabaseConfiguration() DatabaseConfiguration {
	return c.Database
}

func InitConfig() {
	ConfigManager = loadConfiguration()
}
