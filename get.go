package ghcp

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"strconv"
)

var priceRegex *regexp.Regexp

func GetPricesFromNumber(ean string) ([]float64, error) {

	client := &http.Client{}

	domain := getConfig().Domain

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
	q.Add(getConfig().QueryParam, ean)
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
	prices := getPrices(string(body))

	return prices, nil

}

func getPrices(html string) []float64 {

	priceRegex = regexp.MustCompile(getConfig().Regex)

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
