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
	ProductName     string
	ProductPrice    float64
	ComparisonOffer []Offer
}

// The 0th Price is the comparison-Price, the first Price is the best Price
func CheckProductForArticleNr(articleNumber string) (*ProductCheck, error) {

	pc := ProductCheck{}

	// fetch ean and comparison-Price
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
	pc.ComparisonOffer = prices

	return &pc, nil
}
