## TODO
Elasticsearch is now commented out. Es connection issue. 
```go
EsClient, err = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{
				"http://elasticsearch:9200", // e.g., "http://localhost:9200"
			},
		})
```