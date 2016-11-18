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
	fmt.Println("Creating cart =================================")
	p := amazon.CartCreateParameters{
		ResponseGroups: []amazon.CartCreateResponseGroup{
			amazon.CartCreateResponseGroupCart,
			amazon.CartCreateResponseGroupCartNewReleases,
			amazon.CartCreateResponseGroupCartSimilarities,
			amazon.CartCreateResponseGroupCartTopSellers,
		},
	}
	p.AddItemWithASIN("4774182389", 2)
	p.AddItemWithOfferListingID("NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", 4)
	res, err := client.CartCreate(p).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
