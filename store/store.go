package store

import (
	"fmt"
	"os"

	dbr "github.com/gocraft/dbr/v2"
	"github.com/maxzerbini/applico/config"
)

var DbManager *dbr.Session

func NewDbConn() *dbr.Session {
	conf := config.ConfigManager.GetDatabaseConfiguration()
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

func InitDatabaseConn() {
	DbManager = NewDbConn()
}
