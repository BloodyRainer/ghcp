package ghcp

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"log"
	"regexp"
	"errors"
	"strings"
	"strconv"
)

type PriceInfo struct {
	Price    float64
	Merchant string
}

var priceRegex *regexp.Regexp

func fetchPricesForEan(ean string) ([]PriceInfo, error) {

	log.Printf("Trying to fetch pis for EAN: %s\n", ean)

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

	// read pis from html
	pis := parsePriceInfos(string(body))
	if len(pis) == 0 {
		return nil, errors.New("EAN not found on: " + getConfig().PricesDomain)
	}

	return pis, nil

}

func parsePriceInfos(html string) []PriceInfo {

	priceRegex = regexp.MustCompile(getConfig().PricesRegex)

	var pis []PriceInfo

	rawMatches := priceRegex.FindAllStringSubmatch(html, -1)

	//fmt.Println(html)

	for _, p := range rawMatches {

		var priceString string
		priceString = strings.Replace(p[1], ",", ".", -1)
		priceString = strings.Replace(priceString, "--", "00", -1)

		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			log.Println("warn: could not parse '" + priceString + "' to float64")
			continue
		}

		merchant := p[2]
		pi := PriceInfo{
			Price: price,
			Merchant: merchant,
		}

		pis = append(pis, pi)
	}

	return pis
}
