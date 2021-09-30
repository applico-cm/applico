package main

import (
	_ "github.com/lib/pq"
	"github.com/maxzerbini/applico/config"
	"github.com/maxzerbini/applico/store"
)

func main() {
	config.InitConfig()
	store.InitDatabaseConn()
	defer store.DbManager.Close()

}
