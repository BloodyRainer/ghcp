package ghcp

import (
	"bytes"
	"regexp"
	"errors"
	dp "github.com/chromedp/chromedp"
	"log"
	"time"
	"fmt"
)

func fetchEanByArticleNumber(articleNr string) (string, error){

	var err error
	dataChan := make(chan string)

	// Build URL
	var url bytes.Buffer
	url.WriteString("https://")
	url.WriteString(getConfig().EanDomain)
	url.WriteString("/p/search/")
	url.WriteString("?articlenumber=")
	url.WriteString(articleNr)

	// Fetch ProductData
	go func(url string, dc chan string) {
		var productData string
		err = cdpRun(fetchProductData(url,&productData))
		if err != nil {
			log.Panic(err)
		}
		dc <- productData
	}(url.String(), dataChan)

	select {
	case data := <- dataChan:

		fmt.Println("search result: " + data)

		return parseEan(data)
	case <-time.After(20 * time.Second):
		return "", errors.New("search for article number or EAN timed out")
	}

}

func fetchProductData(url string, data *string) dp.Tasks{
	if data == nil {
		log.Panic("data cannot be nil")
	}
	return dp.Tasks {
		dp.Navigate(url),
		dp.WaitVisible(`//div[@class='gridAndInfoContainer']`, dp.BySearch),
		dp.Sleep(100 * time.Millisecond),
		dp.InnerHTML(`//script[@id='productDataJson']`, data),
	}
}

func parseEan(productData string) (string, error) {

	eanRegex := regexp.MustCompile(`\,\"ean"\:\"(\d+)\"\,`)

	eanMatch := eanRegex.FindStringSubmatch(productData)
	if len(eanMatch) < 2 || eanMatch[1] == "" {
		return "", errors.New("no EAN found")
	}

	ean := eanMatch[1]
	if  ean == "null" {
		return "", errors.New("no EAN, EAN is 'null'")
	}

	return ean, nil

}
