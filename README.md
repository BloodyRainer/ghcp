Example Usage of GHCP:

The ENV-Variable `GHCP_Domain` has to be set to www.example.com.

The ENV-Variable `GHCP_QUERY_PARAM` has to be set to a string value.

The ENV-Variable `GHCP_REGEX` has to be set to a string value.

```
// Init GHCP
ghcp.Init()

// Obtain a List of all Prices
prices, err := ghcp.GetPricesFromNumber("0045496452322")

if err != nil {
    log.Fatal(err)
}

for i, p := range  prices{
    fmt.Printf("%d: %v\n", i, p)
}
```