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

const expectedCartAddSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&CartId=352-5038530-7983747&HMAC=l494DklLoQojeL3f8ajE%2F6fvalM%3D&Item.1.ASIN=4774182389&Item.1.Quantity=2&Item.2.OfferListingId=NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%252BbVV876t1%252Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%253D&Item.2.Quantity=4&Operation=CartAdd&Service=AWSECommerceService&Signature=mwWv%2FF%2BRUUeyjf6UdYumqxS%2BkjtDvatQztH7gwsjq%2F4%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartAddRequest(client *Client) *CartAddRequest {
	p := CartAddParameters{
		ResponseGroups: []CartAddResponseGroup{
			CartAddResponseGroupCartNewReleases,
			CartAddResponseGroupCartSimilarities,
			CartAddResponseGroupCartTopSellers,
		},
		HMAC:   "l494DklLoQojeL3f8ajE/6fvalM=",
		CartID: "352-5038530-7983747",
	}
	p.Items.AddASIN("4774182389", 2)
	p.Items.AddOfferListingID("NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", 4)
	return client.CartAdd(p)
}

func TestCartAddSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartAddRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedCartAddSignedURL, signedURL},
		Test{"AK", parsed.Query().Get("AWSAccessKeyId")},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"4774182389", parsed.Query().Get("Item.1.ASIN")},
		Test{"", parsed.Query().Get("Item.1.OfferListingId")},
		Test{"2", parsed.Query().Get("Item.1.Quantity")},
		Test{"", parsed.Query().Get("Item.2.ASIN")},
		Test{"NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", parsed.Query().Get("Item.2.OfferListingId")},
		Test{"4", parsed.Query().Get("Item.2.Quantity")},
		Test{"CartAdd", parsed.Query().Get("Operation")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"mwWv/F+RUUeyjf6UdYumqxS+kjtDvatQztH7gwsjq/4=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartAddDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartAddRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartAddResponseErrorItem.xml")
	gock.New(strings.Replace(expectedCartAddSignedURL, "%2B", "%5C%2B", 2)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、CartIdなどがあります。", err.Error()}.Compare(t)
	}
}

func TestCartAddDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartAddRequest(client)
	gock.New(strings.Replace(expectedCartAddSignedURL, "%2B", "%5C%2B", 2)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartAddSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartAddDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartAddRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartAdd.xml")
	gock.New(strings.Replace(expectedCartAddSignedURL, "%2B", "%5C%2B", 2)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{"352-8545344-3979559", res.Cart.ID},
		Test{"1vXVtADNKYIk3rCThiJfETmY+Uc=", res.Cart.HMAC},
		Test{"1vXVtADNKYIk3rCThiJfETmY%2BUc%3D", res.Cart.URLEncodedHMAC},
		Test{"https://www.amazon.jp/gp/cart/aws-merge.html?cart-id=352-8545344-3979559%26associate-id=ngsio-22%26hmac=1vXVtADNKYIk3rCThiJfETmY%2BUc%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False", res.Cart.PurchaseURL},
		Test{"https://www.amazon.jp/gp/aw/rcart?cart-id=352-8545344-3979559%26associate-id=ngsio-22%26hmac=1vXVtADNKYIk3rCThiJfETmY%2BUc%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False%26uid=NULLGWDOCOMO", res.Cart.MobileCartURL},
		Test{"21952", res.Cart.SubTotal.Amount},
		Test{"JPY", res.Cart.SubTotal.CurrencyCode},
		Test{"￥ 21,952", res.Cart.SubTotal.FormattedPrice},
		Test{"21952", res.Cart.CartItems.SubTotal.Amount},
		Test{"JPY", res.Cart.CartItems.SubTotal.CurrencyCode},
		Test{"￥ 21,952", res.Cart.CartItems.SubTotal.FormattedPrice},
		Test{3, len(res.Cart.CartItems.CartItem)},
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

		Test{"B01JRDPAGO", res.Cart.CartItems.CartItem[2].ASIN},
		Test{"C33X3SVC08ND0I", res.Cart.CartItems.CartItem[2].ID},
		Test{"1260", res.Cart.CartItems.CartItem[2].ItemTotal.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[2].ItemTotal.CurrencyCode},
		Test{"￥ 1,260", res.Cart.CartItems.CartItem[2].ItemTotal.FormattedPrice},
		Test{"630", res.Cart.CartItems.CartItem[2].Price.Amount},
		Test{"JPY", res.Cart.CartItems.CartItem[2].Price.CurrencyCode},
		Test{"￥ 630", res.Cart.CartItems.CartItem[2].Price.FormattedPrice},
		Test{"Book", res.Cart.CartItems.CartItem[2].ProductGroup},
		Test{2, res.Cart.CartItems.CartItem[2].Quantity},
		Test{"Amazon.co.jp", res.Cart.CartItems.CartItem[2].SellerNickname},
		Test{"WIRED VOL.25/特集 The Power of Blockchain ブロックチェーンは世界を変える", res.Cart.CartItems.CartItem[2].Title},
	} {
		test.Compare(t)
	}
}
