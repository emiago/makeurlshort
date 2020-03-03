package config

import (
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type DbConf struct {
	Adress   string
	Port     int
	User     string
	Password string
	Database string
}

type Configuration struct {
	Mysql DbConf
}

var (
	Conf *Configuration = DefaultConf()
)

func DefaultConf() *Configuration {
	return &Configuration{
		Mysql: DbConf{
			Adress:   "localhost",
			Port:     3306,
			User:     "flix",
			Password: "test123",
			Database: "flix",
		},
	}
}

func InitUsingFile(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}

	if _, err := toml.DecodeFile(filename, Conf); err != nil {
		return err
	}

	return nil
}

func Print(out io.Writer) {
	enc := toml.NewEncoder(out)
	enc.Indent = ""
	if err := enc.Encode(Conf); err != nil {
		fmt.Fprint(out, err)
	}
}
