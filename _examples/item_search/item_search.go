package main

import (
	"fmt"
	"log"

	"github.com/ngs/go-amazon-product-advertising-api/amazon"
)

func main() {
	client, err := amazon.NewFromEnvionment()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex: amazon.SearchIndexBooks,
		Keywords:    "Go 言語",
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d results found\n\n", res.TotalResults())
	for _, item := range res.Items() {
		fmt.Printf(`-------------------------------
[Title] %v
[URL]   %v
`, item.ItemAttributes.Title, item.DetailPageURL)
	}
}
