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
	ArticleNumber   string
	Ean             string
	ProductPrice    float64
	ComparisonOffer []Offer
}

// The 0th Price is the comparison-Price, the first Price is the best Price
func CheckProductForArticleNr(articleNumber string) (*ProductCheck, error) {

	pc := ProductCheck{
		ArticleNumber: articleNumber,
	}

	// fetch ean and comparison-Price
	pd, err := fetchProductDataByArticleNumber(articleNumber)
	if err != nil {
		return nil, err
	}
	pc.ProductName = pd.name
	pc.Ean = pd.ean

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
