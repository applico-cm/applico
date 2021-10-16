package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/maxzerbini/applico/config"
	"github.com/maxzerbini/applico/store"
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
