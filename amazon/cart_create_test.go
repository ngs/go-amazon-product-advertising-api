package amazon

import (
	"errors"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

const expectedCartCreateSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Item.1.ASIN=4774182389&Item.1.Quantity=2&Item.2.OfferListingId=NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%252BbVV876t1%252Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%253D&Item.2.Quantity=4&Operation=CartCreate&ResponseGroup=Cart%2CCartNewReleases%2CCartSimilarities%2CCartTopSellers&Service=AWSECommerceService&Signature=SwtdYeDHgAXZdi9%2BrFk2xmpx4bUtvOZtPpOmU7FWfV4%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartCreateRequest(client *Client) *CartCreateRequest {
	p := CartCreateParameters{
		ResponseGroups: []CartCreateResponseGroup{
			CartCreateResponseGroupCart,
			CartCreateResponseGroupCartNewReleases,
			CartCreateResponseGroupCartSimilarities,
			CartCreateResponseGroupCartTopSellers,
		},
	}
	p.Items.AddASIN("4774182389", 2)
	p.Items.AddOfferListingID("NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", 4)
	return client.CartCreate(p)
}

func TestCartCreateSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		{expectedCartCreateSignedURL, signedURL},
		{"AK", parsed.Query().Get("AWSAccessKeyId")},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"4774182389", parsed.Query().Get("Item.1.ASIN")},
		{"", parsed.Query().Get("Item.1.OfferListingId")},
		{"2", parsed.Query().Get("Item.1.Quantity")},
		{"", parsed.Query().Get("Item.2.ASIN")},
		{"NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", parsed.Query().Get("Item.2.OfferListingId")},
		{"4", parsed.Query().Get("Item.2.Quantity")},
		{"CartCreate", parsed.Query().Get("Operation")},
		{"Cart,CartNewReleases,CartSimilarities,CartTopSellers", parsed.Query().Get("ResponseGroup")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"SwtdYeDHgAXZdi9+rFk2xmpx4bUtvOZtPpOmU7FWfV4=", parsed.Query().Get("Signature")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartCreateDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartCreateResponseErrorItem.xml")
	gock.New(strings.Replace(expectedCartCreateSignedURL, "%2B", "%5C%2B", -1)).
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
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	gock.New(strings.Replace(expectedCartCreateSignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartCreateSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartCreateDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartCreateRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartCreate.xml")
	gock.New(strings.Replace(expectedCartCreateSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		{"351-9409673-0414064", res.Cart.ID},
		{"+ak+v8qGeDkkdQ/w0o+5uA2heQI=", res.Cart.HMAC},
		{"%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D", res.Cart.URLEncodedHMAC},
		{"https://www.amazon.jp/gp/cart/aws-merge.html?cart-id=351-9409673-0414064%26associate-id=ngsio-22%26hmac=%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False", res.Cart.PurchaseURL},
		{"https://www.amazon.jp/gp/aw/rcart?cart-id=351-9409673-0414064%26associate-id=ngsio-22%26hmac=%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False%26uid=NULLGWDOCOMO", res.Cart.MobileCartURL},
		{"20692", res.Cart.SubTotal.Amount},
		{"JPY", res.Cart.SubTotal.CurrencyCode},
		{"￥ 20,692", res.Cart.SubTotal.FormattedPrice},
		{"20692", res.Cart.CartItems.SubTotal.Amount},
		{"JPY", res.Cart.CartItems.SubTotal.CurrencyCode},
		{"￥ 20,692", res.Cart.CartItems.SubTotal.FormattedPrice},
		{2, len(res.Cart.CartItems.CartItem)},

		{"4774182389", res.Cart.CartItems.CartItem[0].ASIN},
		{"CYOXSP4DGCG16", res.Cart.CartItems.CartItem[0].ID},
		{"4276", res.Cart.CartItems.CartItem[0].ItemTotal.Amount},
		{"JPY", res.Cart.CartItems.CartItem[0].ItemTotal.CurrencyCode},
		{"￥ 4,276", res.Cart.CartItems.CartItem[0].ItemTotal.FormattedPrice},
		{"2138", res.Cart.CartItems.CartItem[0].Price.Amount},
		{"JPY", res.Cart.CartItems.CartItem[0].Price.CurrencyCode},
		{"￥ 2,138", res.Cart.CartItems.CartItem[0].Price.FormattedPrice},
		{"Book", res.Cart.CartItems.CartItem[0].ProductGroup},
		{2, res.Cart.CartItems.CartItem[0].Quantity},
		{"Amazon.co.jp", res.Cart.CartItems.CartItem[0].SellerNickname},
		{"Slack入門 [ChatOpsによるチーム開発の効率化]", res.Cart.CartItems.CartItem[0].Title},

		{"4621300253", res.Cart.CartItems.CartItem[1].ASIN},
		{"C30K5HAY097OZO", res.Cart.CartItems.CartItem[1].ID},
		{"16416", res.Cart.CartItems.CartItem[1].ItemTotal.Amount},
		{"JPY", res.Cart.CartItems.CartItem[1].ItemTotal.CurrencyCode},
		{"￥ 16,416", res.Cart.CartItems.CartItem[1].ItemTotal.FormattedPrice},
		{"4104", res.Cart.CartItems.CartItem[1].Price.Amount},
		{"JPY", res.Cart.CartItems.CartItem[1].Price.CurrencyCode},
		{"￥ 4,104", res.Cart.CartItems.CartItem[1].Price.FormattedPrice},
		{"Book", res.Cart.CartItems.CartItem[1].ProductGroup},
		{4, res.Cart.CartItems.CartItem[1].Quantity},
		{"Amazon.co.jp", res.Cart.CartItems.CartItem[1].SellerNickname},
		{"プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)", res.Cart.CartItems.CartItem[1].Title},
	} {
		test.Compare(t)
	}
}
