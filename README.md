# GHCP

## ENV-Variables:

The ENV-Variable `GHCP_PRICES_Domain` has to be set to www.example.com.

The ENV-Variable `GHCP_PRICES_QUERY_PARAM` has to be set to a string value.

The ENV-Variable `GHCP_PRICES_REGEX` has to be set to a string value.

The ENV-Variable `GHCP_EAN_DOMAIN` has to be set to www.example.com.

The ENV-Variable `GHCP_DEBUG` can be set to TRUE to enable Debug-Logging.

## Example Usage:
```
ghcp.Init() // Init GHCP

wg := sync.WaitGroup{}

c := ghcp.MakeCounter(60*time.Second, 2) // Every 60s count the occurences and return the top 2 articles
caChan := c.CountArticleChan() // fill this chan with article numbers
topArticlesChan := c.TopNArticlesChan() // receive the top articles

start := time.Now()
wg.Add(1)
go func() {
    for i := 0; i < 1200000; i++ {
        time.Sleep(20 * time.Millisecond)
        an := RandomArticleNr()
        caChan <- an
    }
    log.Println("finished articles")
    wg.Done()
}()

go func() {
    for articles := range topArticlesChan {

        log.Printf("----------------------------------------------------------------\n")
        log.Printf("1: aNr: %v, Count: %v \n", articles[0].Number, articles[0].Count)
        log.Printf("2: aNr: %v, Count: %v \n", articles[1].Number, articles[1].Count)

        prices, err :=ghcp.FetchPricesForArticleNr(articles[0].Number)
        if err != nil {
            log.Printf("failed to fetch prices for EAN: %v", err.Error())
        }

        for i, p := range  prices{
            fmt.Printf("%d: %v\n", i, p)
        }
    }
}()

wg.Wait()

log.Printf("Finished after %v", time.Since(start))
```