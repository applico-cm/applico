package main

import (
	"log"

	"github.com/applico-cm/applico/config"
	"github.com/applico-cm/applico/store"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.ConfigurationFromEnv()
	if err != nil {
		log.Printf("Configuration initialization error: %v", err)
		return
	}
	store.InitDatabaseConn(conf.DatabaseConfiguration())
	defer store.DbManager.Close()

	server := NewServer()

	router := gin.Default()

	router.GET("/users/:customerid", server.ListUsers)

	router.Run("0.0.0.0:8080")

}
