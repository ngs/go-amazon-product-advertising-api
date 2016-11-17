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
		ResponseGroups: []amazon.ItemSearchResponseGroup{
			amazon.ItemSearchResponseGroupLarge,
			amazon.ItemSearchResponseGroupOfferFull,
		},
		SearchIndex: amazon.SearchIndexBooks,
		Keywords:    "Go 言語",
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
