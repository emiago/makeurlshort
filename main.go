package main

import (
	"flixwebtest/config"
	"flixwebtest/urlshorty"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/lithammer/shortuuid/v3"
)

var (
	// db   *sql.DB
	conf *config.Configuration
)

func main() {
	conf = config.Conf

	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	urlshorty.SetParser(&urlshorty.DefaultParser{
		Domain: "http://localhost:8080/r",
	})

	// urlshorty.SetParser(&CustomUrlShorty{
	// 	Domain: "http://localhost:8080/r",
	// })

	log.Info("Starting server")
	serverRun()
}

type CustomUrlShorty struct {
	Domain string
}

func (p *CustomUrlShorty) Parse(longurl string) (string, error) {
	return fmt.Sprintf("%s/%s", p.Domain, shortuuid.NewWithNamespace(longurl)), nil
}
