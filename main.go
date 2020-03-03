package main

import (
	"database/sql"
	"flixwebtest/config"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db   *sql.DB
	conf *config.Configuration
)

func main() {
	conf = config.Conf

	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.Info("Starting server")
	serverRun()
}
