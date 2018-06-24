package ghcp

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"strconv"
	"errors"
)

var priceRegex *regexp.Regexp

func fetchPricesForEan(ean string) ([]float64, error) {

	log.Printf("Trying to fetch prices for EAN: %s", ean)

	client := &http.Client{}

	domain := getConfig().PricesDomain

	// Build URL
	var url bytes.Buffer
	url.WriteString("https://")
	url.WriteString(domain)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add Query Parameters to URL
	q := req.URL.Query()
	q.Add(getConfig().PricesQueryParam, ean)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// read prices from html
	prices := parsePrices(string(body))
	if len(prices) == 0 {
		return nil, errors.New("EAN not found on: " + getConfig().PricesDomain)
	}

	return prices, nil

}

func parsePrices(html string) []float64 {

	priceRegex = regexp.MustCompile(getConfig().PricesRegex)

	var prices []float64

	rawPriceMatches := priceRegex.FindAllStringSubmatch(html, -1)

	for _, p := range rawPriceMatches {
		var priceString string
		priceString = strings.Replace(p[1], ",", ".", -1)
		priceString = strings.Replace(priceString, "--", "00", -1)

		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			log.Println("warn: could not parse '" + priceString + "' to float64")
			continue
		}

		prices = append(prices, price)
	}

	return prices
}
