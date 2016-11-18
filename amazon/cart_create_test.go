package amazon

import (
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

func TestCartCreateItemQuery(t *testing.T) {
	Test{
		map[string]string{
			"ASIN":     "test1",
			"Quantity": "2",
		},
		CartCreateItem{
			ASIN:     "test1",
			Quantity: 2,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"OfferListingId": "test1",
			"Quantity":       "2",
		},
		CartCreateItem{
			OfferListingID: "test1",
			Quantity:       2,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"ASIN": "test1",
		},
		CartCreateItem{
			ASIN: "test1",
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"OfferListingId": "test1",
		},
		CartCreateItem{
			OfferListingID: "test1",
		}.Query(),
	}.DeepEqual(t)
}

func TestCartCreateParametersAddItem(t *testing.T) {
	p := &CartCreateParameters{}
	Test{0, len(p.Items)}.Compare(t)
	p.AddItemWithASIN("test1", 2)
	Test{1, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].ASIN}.Compare(t)
	Test{2, p.Items[0].Quantity}.Compare(t)
	p.AddItemWithOfferListingID("test2", 4)
	Test{2, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].ASIN}.Compare(t)
	Test{"", p.Items[0].OfferListingID}.Compare(t)
	Test{2, p.Items[0].Quantity}.Compare(t)
	Test{"", p.Items[1].ASIN}.Compare(t)
	Test{"test2", p.Items[1].OfferListingID}.Compare(t)
	Test{4, p.Items[1].Quantity}.Compare(t)
}

const expectedCartCreateSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Item.1.ASIN=4774182389&Item.1.Quantity=2&Item.2.OfferListingId=NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%252BbVV876t1%252Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%253D&Item.2.Quantity=4&Operation=CartCreate&Service=AWSECommerceService&Signature=Gx1wPKKFuodv6rxdbkeShOLj5SPWJ7JaNIYLapPlZew%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartCreateRequest(client *Client) *CartCreateRequest {
	p := CartCreateParameters{
		ResponseGroups: []CartCreateResponseGroup{
			CartCreateResponseGroupCart,
			CartCreateResponseGroupCartNewReleases,
			CartCreateResponseGroupCartSimilarities,
			CartCreateResponseGroupCartTopSellers,
		},
	}
	p.AddItemWithASIN("4774182389", 2)
	p.AddItemWithOfferListingID("NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", 4)
	return client.CartCreate(p)
}

func TestCartCreateSignedURL(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedCartCreateSignedURL, signedURL},
		Test{"AK", parsed.Query().Get("AWSAccessKeyId")},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"4774182389", parsed.Query().Get("Item.1.ASIN")},
		Test{"", parsed.Query().Get("Item.1.OfferListingId")},
		Test{"2", parsed.Query().Get("Item.1.Quantity")},
		Test{"", parsed.Query().Get("Item.2.ASIN")},
		Test{"NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", parsed.Query().Get("Item.2.OfferListingId")},
		Test{"4", parsed.Query().Get("Item.2.Quantity")},
		Test{"CartCreate", parsed.Query().Get("Operation")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"Gx1wPKKFuodv6rxdbkeShOLj5SPWJ7JaNIYLapPlZew=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartCreateDoErrorResponse(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartCreateResponseErrorItem.xml")
	gock.New(expectedCartCreateSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestCartCreateDoError(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	gock.New(expectedCartCreateSignedURL).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartCreateSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartCreateDo(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartCreate.xml")
	gock.New(expectedCartCreateSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{"352-5038530-7983747", res.Cart.ID},
		Test{"l494DklLoQojeL3f8ajE/6fvalM=", res.Cart.HMAC},
		Test{"l494DklLoQojeL3f8ajE%2F6fvalM%3D", res.Cart.URLEncodedHMAC},
		Test{"https://www.amazon.jp/gp/cart/aws-merge.html?cart-id=352-5038530-7983747%26associate-id=ngsio-22%26hmac=l494DklLoQojeL3f8ajE%2F6fvalM%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False", res.Cart.PurchaseURL},
		Test{"https://www.amazon.jp/gp/aw/rcart?cart-id=352-5038530-7983747%26associate-id=ngsio-22%26hmac=l494DklLoQojeL3f8ajE%2F6fvalM%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False%26uid=NULLGWDOCOMO", res.Cart.MobileCartURL},
		Test{"20692", res.Cart.SubTotal.Amount},
		Test{"JPY", res.Cart.SubTotal.CurrencyCode},
		Test{"￥ 20,692", res.Cart.SubTotal.FormattedPrice},
		Test{"20692", res.Cart.CartItems.SubTotal.Amount},
		Test{"JPY", res.Cart.CartItems.SubTotal.CurrencyCode},
		Test{"￥ 20,692", res.Cart.CartItems.SubTotal.FormattedPrice},
		Test{2, len(res.Cart.CartItems.CartItem)},
		Test{"4774182389", res.Cart.CartItems.CartItem[0].ASIN},
		Test{"CYOXSP4DGCG16", res.Cart.CartItems.CartItem[0].ID},
		Test{"4276", res.Cart.CartItems.CartItem[0].ItemTotal.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[0].ItemTotal.CurrencyCode},
		Test{"￥ 4,276", res.Cart.CartItems.CartItem[0].ItemTotal.FormattedPrice},
		Test{"2138", res.Cart.CartItems.CartItem[0].Price.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[0].Price.CurrencyCode},
		Test{"￥ 2,138", res.Cart.CartItems.CartItem[0].Price.FormattedPrice},
		Test{"Book", res.Cart.CartItems.CartItem[0].ProductGroup},
		Test{2, res.Cart.CartItems.CartItem[0].Quantity},
		Test{"Amazon.co.jp", res.Cart.CartItems.CartItem[0].SellerNickname},
		Test{"Slack入門 [ChatOpsによるチーム開発の効率化]", res.Cart.CartItems.CartItem[0].Title},
		Test{"4621300253", res.Cart.CartItems.CartItem[1].ASIN},
		Test{"C30K5HAY097OZO", res.Cart.CartItems.CartItem[1].ID},
		Test{"16416", res.Cart.CartItems.CartItem[1].ItemTotal.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[1].ItemTotal.CurrencyCode},
		Test{"￥ 16,416", res.Cart.CartItems.CartItem[1].ItemTotal.FormattedPrice},
		Test{"4104", res.Cart.CartItems.CartItem[1].Price.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[1].Price.CurrencyCode},
		Test{"￥ 4,104", res.Cart.CartItems.CartItem[1].Price.FormattedPrice},
		Test{"Book", res.Cart.CartItems.CartItem[1].ProductGroup},
		Test{4, res.Cart.CartItems.CartItem[1].Quantity},
		Test{"Amazon.co.jp", res.Cart.CartItems.CartItem[1].SellerNickname},
		Test{"プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)", res.Cart.CartItems.CartItem[1].Title},
	} {
		test.Compare(t)
	}
}
