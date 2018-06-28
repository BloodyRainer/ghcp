package ghcp

import (
	"bytes"
	"regexp"
	"errors"
	"os/exec"
)

func fetchEanAndPriceByArticleNumber(articleNr string) (string, string, error) {

	// Build URL
	var url bytes.Buffer
	url.WriteString("https://")
	url.WriteString(getConfig().EanDomain)
	url.WriteString("/p/search/")
	url.WriteString("?articlenumber=")
	url.WriteString(articleNr)

	out, err := exec.Command("google-chrome", "--headless", "--disable-gpu", "--dump-dom", url.String()).Output()

	if err != nil {
		return "", "", err
	}

	ean, err := parseEan(string(out))
	if err != nil {
		return "", "", err
	}

	price, err := parsePrice(string(out))
	if err != nil {
		return "", "", err
	}

	return ean, price, nil

}

func parseEan(productData string) (string, error) {

	eanRegex := regexp.MustCompile(`\,\"ean"\:\"(\d+)\"\,`)

	eanMatch := eanRegex.FindStringSubmatch(productData)
	if len(eanMatch) < 2 || eanMatch[1] == "" {
		return "", errors.New("no EAN found")
	}

	ean := eanMatch[1]
	if  ean == "null" {
		return "", errors.New("EAN is 'null'")
	}

	return ean, nil

}

func parsePrice(productData string) (string, error) {
	priceRegex := regexp.MustCompile(`\,\"techPriceAmount"\:\"(\d+\.\d+)\"\,`)

	priceMatch := priceRegex.FindStringSubmatch(productData)
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	price := priceMatch[1]
	if  price == "null" {
		return "", errors.New("price is 'null'")
	}

	return price, nil
}
