Example Usage of GHCP:

The ENV-Variable `GHCP_PRICES_Domain` has to be set to www.example.com.

The ENV-Variable `GHCP_PRICES_QUERY_PARAM` has to be set to a string value.

The ENV-Variable `GHCP_PRICES_REGEX` has to be set to a string value.

The ENV-Variable `GHCP_EAN_DOMAIN` has to be set to www.example.com.

```
// Init GHCP
ghcp.Init()
defer ghcp.ShutDown()

// Fetch Prices
prices, err :=ghcp.FetchPricesForArticleNr("74320174")
if err != nil {
    log.Fatal(err)
}

for i, p := range  prices{
    fmt.Printf("%d: %v\n", i, p)
	}
```