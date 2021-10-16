package store

import (
	"fmt"
	"os"

	"github.com/applico-cm/applico/config"
	dbr "github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

var DbManager *dbr.Session

func NewDbConn(conf *config.DatabaseConfiguration) *dbr.Session {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	conn, err := dbr.Open("postgres", psqlInfo, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := conn.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn.SetMaxOpenConns(10)

	return conn.NewSession(nil)
}

func InitDatabaseConn(conf *config.DatabaseConfiguration) {
	DbManager = NewDbConn(conf)
}
