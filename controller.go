package ghcp

func Init() {
	initConfig()
	initChromeDp()
}

func ShutDown() {
	shutDownChromeDP()
}

func FetchPricesForArticleNr(articleNumber string) ([]float64, error){

	ean, err := fetchEanByArticleNumber(articleNumber)
	if err != nil {
		return nil, err
	}

	prices, err := fetchPricesForEan(ean)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
