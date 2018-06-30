package ghcp

import (
	"strconv"
	"log"
)

func Init() {
	log.Println("Start GHCP")
	initConfig()
}

// The 0th price is the comparison-price, the first price is the best price
func FetchPricesForArticleNr(articleNumber string) ([]float64, error){

	var prices []float64

	// fetch ean and comparison-price
	ean, price, err := fetchEanAndPriceByArticleNumber(articleNumber)
	if err != nil {
		return nil, err
	}

	// set comparison-price to 0th element of price slice
	p, err := strconv.ParseFloat(price, 64)
	prices = append(prices, p)

	// fetch best prices
	fp, err := fetchPricesForEan(ean)
	if err != nil {
		return nil, err
	}

	// append best prices to slice
	prices = append(prices, fp...)

	return prices, nil
}
