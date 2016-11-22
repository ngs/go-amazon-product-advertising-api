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

const expectedCartAddSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&CartId=352-5038530-7983747&HMAC=l494DklLoQojeL3f8ajE%2F6fvalM%3D&Item.1.ASIN=4774182389&Item.1.Quantity=2&Item.2.OfferListingId=NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%252BbVV876t1%252Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%253D&Item.2.Quantity=4&Operation=CartAdd&ResponseGroup=CartNewReleases%2CCartSimilarities%2CCartTopSellers&Service=AWSECommerceService&Signature=dTqTWl5sKok1Y5iMxKy4iysCr%2BjjwPbNrzbSW3AsfSo%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

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
		{expectedCartAddSignedURL, signedURL},
		{"AK", parsed.Query().Get("AWSAccessKeyId")},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"4774182389", parsed.Query().Get("Item.1.ASIN")},
		{"", parsed.Query().Get("Item.1.OfferListingId")},
		{"2", parsed.Query().Get("Item.1.Quantity")},
		{"", parsed.Query().Get("Item.2.ASIN")},
		{"NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", parsed.Query().Get("Item.2.OfferListingId")},
		{"4", parsed.Query().Get("Item.2.Quantity")},
		{"CartAdd", parsed.Query().Get("Operation")},
		{"CartNewReleases,CartSimilarities,CartTopSellers", parsed.Query().Get("ResponseGroup")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"dTqTWl5sKok1Y5iMxKy4iysCr+jjwPbNrzbSW3AsfSo=", parsed.Query().Get("Signature")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"2013-08-01", parsed.Query().Get("Version")},
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
		{"351-5204090-0802017", res.Cart.ID},
		{"ba6oTpgnCNQTEfb67iNFlwqeLp8=", res.Cart.HMAC},
		{"ba6oTpgnCNQTEfb67iNFlwqeLp8%3D", res.Cart.URLEncodedHMAC},
		{"https://www.amazon.jp/gp/cart/aws-merge.html?cart-id=351-5204090-0802017%26associate-id=ngsio-22%26hmac=ba6oTpgnCNQTEfb67iNFlwqeLp8%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False", res.Cart.PurchaseURL},
		{"https://www.amazon.jp/gp/aw/rcart?cart-id=351-5204090-0802017%26associate-id=ngsio-22%26hmac=ba6oTpgnCNQTEfb67iNFlwqeLp8%3D%26SubscriptionId=AKIAITPH62XKCOOT7AKA%26MergeCart=False%26uid=NULLGWDOCOMO", res.Cart.MobileCartURL},
		{"21952", res.Cart.SubTotal.Amount},
		{"JPY", res.Cart.SubTotal.CurrencyCode},
		{"￥ 21,952", res.Cart.SubTotal.FormattedPrice},
		{"21952", res.Cart.CartItems.SubTotal.Amount},
		{"JPY", res.Cart.CartItems.SubTotal.CurrencyCode},
		{"￥ 21,952", res.Cart.CartItems.SubTotal.FormattedPrice},
		{3, len(res.Cart.CartItems.CartItem)},
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

		{"B01JRDPAGO", res.Cart.CartItems.CartItem[2].ASIN},
		{"C33X3SVC08ND0I", res.Cart.CartItems.CartItem[2].ID},
		{"1260", res.Cart.CartItems.CartItem[2].ItemTotal.Amount},
		{"JPY", res.Cart.CartItems.CartItem[2].ItemTotal.CurrencyCode},
		{"￥ 1,260", res.Cart.CartItems.CartItem[2].ItemTotal.FormattedPrice},
		{"630", res.Cart.CartItems.CartItem[2].Price.Amount},
		{"JPY", res.Cart.CartItems.CartItem[2].Price.CurrencyCode},
		{"￥ 630", res.Cart.CartItems.CartItem[2].Price.FormattedPrice},
		{"Book", res.Cart.CartItems.CartItem[2].ProductGroup},
		{2, res.Cart.CartItems.CartItem[2].Quantity},
		{"Amazon.co.jp", res.Cart.CartItems.CartItem[2].SellerNickname},
		{"WIRED VOL.25/特集 The Power of Blockchain ブロックチェーンは世界を変える", res.Cart.CartItems.CartItem[2].Title},
	} {
		test.Compare(t)
	}
}
