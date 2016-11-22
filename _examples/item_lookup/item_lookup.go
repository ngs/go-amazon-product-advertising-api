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
	res, err := client.ItemLookup(amazon.ItemLookupParameters{
		ResponseGroups: []amazon.ItemLookupResponseGroup{
			amazon.ItemLookupResponseGroupLarge,
		},
		IDType: amazon.IDTypeASIN,
		ItemIDs: []string{
			"477418392X",
			"B01FH3KRTI",
			"4621300253",
			"4873117526",
			"4865940413",
			"4863541783",
			"4798031801",
			"4774184322",
			"4863541171",
			"B01LDFK76M",
		},
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range res.Items.Item {
		fmt.Printf(`-------------------------------
[Title] %v
[URL]   %v
`, item.ItemAttributes.Title, item.DetailPageURL)
	}
}
