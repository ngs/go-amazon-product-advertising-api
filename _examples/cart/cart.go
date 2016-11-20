package main

import (
	"fmt"
	"log"
	"time"

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
	p.Items.AddASIN("B01JRDPAGO", 2)
	res, err := client.CartCreate(p).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Cart.PurchaseURL)
	time.Sleep(time.Second * 2)
	fmt.Println("Adding items to cart =================================")
	p2 := amazon.CartAddParameters{
		ResponseGroups: []amazon.CartAddResponseGroup{
			amazon.CartAddResponseGroupCart,
			amazon.CartAddResponseGroupCartNewReleases,
			amazon.CartAddResponseGroupCartSimilarities,
			amazon.CartAddResponseGroupCartTopSellers,
		},
		HMAC:   res.Cart.HMAC,
		CartID: res.Cart.ID,
	}
	p2.Items.AddASIN("4774182389", 2)
	p2.Items.AddOfferListingID("NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", 4)
	res2, err := client.CartAdd(p2).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2.Cart.PurchaseURL)
	time.Sleep(time.Second * 2)
	res3, err := client.CartGet(amazon.CartGetParameters{
		ResponseGroups: []amazon.CartGetResponseGroup{
			amazon.CartGetResponseGroupCart,
			amazon.CartGetResponseGroupCartTopSellers,
			amazon.CartGetResponseGroupCartSimilarities,
			amazon.CartGetResponseGroupCartNewReleases,
		},
		CartID:     res2.Cart.ID,
		HMAC:       res2.Cart.HMAC,
		CartItemID: res2.Cart.CartItems.CartItem[0].ID,
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res3.Cart.PurchaseURL)
}
