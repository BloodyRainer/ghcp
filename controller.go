package ghcp

import (
	"strconv"
	"log"
)

func Init() {
	log.Println("Start GHCP")
	initConfig()
}

type ProductCheck struct {
	ProductName       string
	ProductPrice      float64
	ComparisionPrices []float64
}

// The 0th price is the comparison-price, the first price is the best price
func CheckProductForArticleNr(articleNumber string) (*ProductCheck, error) {

	pc := ProductCheck{}

	// fetch ean and comparison-price
	pd, err := fetchProductDataByArticleNumber(articleNumber)
	if err != nil {
		return nil, err
	}
	pc.ProductName = pd.name

	p, err := strconv.ParseFloat(pd.price, 64)
	if err != nil {
		return nil, err
	}
	pc.ProductPrice = p

	// fetch best prices
	prices, err := fetchPricesForEan(pd.ean)
	if err != nil {
		return nil, err
	}
	pc.ComparisionPrices = prices

	return &pc, nil
}
