package ghcp

import (
	"bytes"
	"regexp"
	"errors"
	"os/exec"
	"strings"
)

var eanReg *regexp.Regexp
var priceReg *regexp.Regexp
var nameReg *regexp.Regexp

type productData struct {
	ean string
	price string
	name string
}

func init() {
	eanReg = regexp.MustCompile(`\,\"ean"\:\"(\d+)\"\,`)
	priceReg = regexp.MustCompile(`\,\"techPriceAmount"\:\"(\d+\.\d+)\"\,`)
	nameReg = regexp.MustCompile(`:{"id":".*","name":"(.*)","links"`)
}

func fetchProductDataByArticleNumber(articleNr string) (productData, error) {

	pd := productData{}

	// Build URL
	var url bytes.Buffer
	url.WriteString("https://")
	url.WriteString(getConfig().EanDomain)
	url.WriteString("/p/search/")
	url.WriteString("?articlenumber=")
	url.WriteString(articleNr)

	out, err := exec.Command("google-chrome", "--headless", "--disable-gpu", "--dump-dom", url.String()).Output()

	if err != nil {
		return pd, err
	}

	ean, err := parseEan(string(out))
	if err != nil {
		return pd, err
	}
	pd.ean = ean

	price, err := parsePrice(string(out))
	if err != nil {
		return pd, err
	}
	pd.price = price

	name, err := parseName(string(out))
	if err != nil {
		return pd, err
	}
	pd.name = name

	return pd, nil

}

func parseEan(productData string) (string, error) {

	match := eanReg.FindStringSubmatch(productData)

	if len(match) < 2 || match[1] == "" {
		return "", errors.New("no EAN found")
	}

	ean := match[1]
	if  ean == "null" {
		return "", errors.New("EAN is 'null'")
	}

	return ean, nil

}

func parsePrice(productData string) (string, error) {

	priceMatch := priceReg.FindStringSubmatch(productData)
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	price := priceMatch[1]
	if  price == "null" {
		return "", errors.New("price is 'null'")
	}

	return price, nil
}

func parseName(body string) (string, error) {
	nameMatch := nameReg.FindStringSubmatch(body)
	if len(nameMatch) < 2 || nameMatch[1] == "" {
		return "", errors.New("no name found")
	}

	// TODO: dirty hacks
	name := strings.Replace(nameMatch[1], "&quot;", `'`, -1)
	name = strings.Replace(name, "&amp;", "&", -1)

	return name, nil
}
