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
	fmt.Println("Getting items to cart =================================")
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
	time.Sleep(time.Second * 2)
	fmt.Println("Modifying items =================================")
	p4 := amazon.CartModifyParameters{
		ResponseGroups: []amazon.CartModifyResponseGroup{
			amazon.CartModifyResponseGroupCart,
			amazon.CartModifyResponseGroupCartSimilarities,
			amazon.CartModifyResponseGroupCartNewReleases,
			amazon.CartModifyResponseGroupCartTopSellers,
		},
		CartID: res3.Cart.ID,
		HMAC:   res3.Cart.HMAC,
	}
	p4.Items.ModifyQuantity(res3.Cart.CartItems.CartItem[0].ID, 0)
	p4.Items.SaveForLater(res3.Cart.CartItems.CartItem[1].ID)
	p4.Items.MoveToCart(res3.Cart.CartItems.CartItem[2].ID)
	res4, err := client.CartModify(p4).Do()
	fmt.Println(res4.Cart.PurchaseURL)
}
