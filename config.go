package ghcp

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Domain     string `required:"true"`
	QueryParam string `required:"true" split_words:"true"`
	Regex      string `required:"true"`
}

var c Config

func Init() {
	err := envconfig.Process("ghcp", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("ghcp_domain: %s", c.Domain)
}

func getConfig() Config {
	return c
}
