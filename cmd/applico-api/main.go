package main

import (
	"log"

	"github.com/applico-cm/applico/config"
	"github.com/applico-cm/applico/store"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.ConfigurationFromEnv()
	if err != nil {
		log.Printf("Configuration initialization error: %v", err)
		return
	}
	store.InitDatabaseConn(conf.DatabaseConfiguration())
	defer store.DbManager.Close()

}
