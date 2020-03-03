package main

import (
	"database/sql"
	"flixwebtest/config"
	"fmt"

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

	log.Info("Connecting to database")
	connectDatabase()
	log.Info("Creating tables")
	createTables()
	log.Info("Starting server")
	serverRun()
}

func connectDatabase() {
	var err error
	socket := fmt.Sprintf("tcp(%s:%d)", conf.Mysql.Adress, conf.Mysql.Port)
	// if conf.Mysql.Adress == "localhost" {
	// 	socket = "unix(/var/run/mysqld/mysqld.sock)"
	// }

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", conf.Mysql.User, conf.Mysql.Password, socket, conf.Mysql.Database))
	if err != nil {
		log.WithError(err).Fatal("Failed to connect")
	}

	if err = db.Ping(); err != nil {
		log.WithError(err).Fatal("Failed to ping")
	}
}

func createTables() {
	q := `
CREATE TABLE IF NOT EXISTS buses (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(50) DEFAULT '' NOT NULL,
line INT DEFAULT 0 NOT NULL
) ENGINE=InnoDB`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}
