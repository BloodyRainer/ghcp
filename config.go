package ghcp

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	PricesDomain     string `required:"true" split_words:"true"`
	PricesQueryParam string `required:"true" split_words:"true"`
	PricesRegex      string `required:"true" split_words:"true"`
	EanDomain        string `required:"true" split_words:"true"`
	Debug            string `default:"false"`
}

var c Config

func initConfig() {
	err := envconfig.Process("ghcp", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("GHCP params:\nghcp_prices_domain: %s\n"+
		"ghcp_prices_query_param: %s\n"+
		"ghcp_prices_regex: %s\n"+
		"ghcp_ean_domain: %s\n"+
		"ghcp_debug: %s\n"+
		"\n",
		c.PricesDomain,
		c.PricesQueryParam,
		c.PricesRegex,
		c.EanDomain,
		c.Debug)
}

func getConfig() Config {
	return c
}
