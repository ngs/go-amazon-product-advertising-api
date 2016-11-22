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

const expectedCartModifySignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&CartId=351-9409673-0414064&HMAC=%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D&Item.1.CartItemId=test1&Item.1.Quantity=1&Item.2.CartItemId=test2&Item.2.Quantity=0&Item.3.Action=SaveForLater&Item.3.CartItemId=test3&Item.4.Action=MoveToCart&Item.4.CartItemId=test4&Operation=CartModify&ResponseGroup=Cart%2CCartNewReleases%2CCartSimilarities%2CCartTopSellers&Service=AWSECommerceService&Signature=gSW1C2LfE%2BetUa2z14IJSE28XWQDNamNW59Eti2KMng%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartModifyRequest(client *Client) *CartModifyRequest {
	p := CartModifyParameters{
		ResponseGroups: []CartModifyResponseGroup{
			CartModifyResponseGroupCart,
			CartModifyResponseGroupCartNewReleases,
			CartModifyResponseGroupCartSimilarities,
			CartModifyResponseGroupCartTopSellers,
		},
		CartID: "351-9409673-0414064",
		HMAC:   "+ak+v8qGeDkkdQ/w0o+5uA2heQI=",
	}
	p.Items.ModifyQuantity("test1", 1)
	p.Items.ModifyQuantity("test2", 0)
	p.Items.SaveForLater("test3")
	p.Items.MoveToCart("test4")
	return client.CartModify(p)
}

func TestCartModifySignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartModifyRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		{expectedCartModifySignedURL, signedURL},
		{"AK", parsed.Query().Get("AWSAccessKeyId")},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"+ak+v8qGeDkkdQ/w0o+5uA2heQI=", parsed.Query().Get("HMAC")},
		{"351-9409673-0414064", parsed.Query().Get("CartId")},
		{"", parsed.Query().Get("Item.1.Action")},
		{"test1", parsed.Query().Get("Item.1.CartItemId")},
		{"1", parsed.Query().Get("Item.1.Quantity")},
		{"", parsed.Query().Get("Item.2.Action")},
		{"test2", parsed.Query().Get("Item.2.CartItemId")},
		{"0", parsed.Query().Get("Item.2.Quantity")},
		{"SaveForLater", parsed.Query().Get("Item.3.Action")},
		{"test3", parsed.Query().Get("Item.3.CartItemId")},
		{"", parsed.Query().Get("Item.3.Quantity")},
		{"MoveToCart", parsed.Query().Get("Item.4.Action")},
		{"test4", parsed.Query().Get("Item.4.CartItemId")},
		{"", parsed.Query().Get("Item.4.Quantity")},
		{"CartModify", parsed.Query().Get("Operation")},
		{"Cart,CartNewReleases,CartSimilarities,CartTopSellers", parsed.Query().Get("ResponseGroup")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"gSW1C2LfE+etUa2z14IJSE28XWQDNamNW59Eti2KMng=", parsed.Query().Get("Signature")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartModifyDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartModifyRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartModifyResponseErrorItem.xml")
	gock.New(strings.Replace(expectedCartModifySignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.ExactParameterRequirement: 次のパラメータのうち、1個がリクエストに含まれている必要があります：Quantity, Action", err.Error()}.Compare(t)
	}
}

func TestCartModifyDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartModifyRequest(client)
	gock.New(strings.Replace(expectedCartModifySignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartModifySignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartModifyDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartModifyRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartModify.xml")
	gock.New(strings.Replace(expectedCartModifySignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		{"352-4323378-0926412", res.Cart.ID},
		{1, len(res.Cart.SavedForLaterItems.SavedForLaterItem)},
		{"16416", res.Cart.SavedForLaterItems.SubTotal.Amount},
		{"JPY", res.Cart.SavedForLaterItems.SubTotal.CurrencyCode},
		{"￥ 16,416", res.Cart.SavedForLaterItems.SubTotal.FormattedPrice},
		{"4621300253", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ASIN},
		{"S30K5HAY097OZO", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ID},
		{"16416", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ItemTotal.Amount},
		{"JPY", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ItemTotal.CurrencyCode},
		{"￥ 16,416", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ItemTotal.FormattedPrice},
		{"4104", res.Cart.SavedForLaterItems.SavedForLaterItem[0].Price.Amount},
		{"JPY", res.Cart.SavedForLaterItems.SavedForLaterItem[0].Price.CurrencyCode},
		{"￥ 4,104", res.Cart.SavedForLaterItems.SavedForLaterItem[0].Price.FormattedPrice},
		{"Book", res.Cart.SavedForLaterItems.SavedForLaterItem[0].ProductGroup},
		{4, res.Cart.SavedForLaterItems.SavedForLaterItem[0].Quantity},
		{"Amazon.co.jp", res.Cart.SavedForLaterItems.SavedForLaterItem[0].SellerNickname},
		{"プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)", res.Cart.SavedForLaterItems.SavedForLaterItem[0].Title},
		// TODO: write more tests
	} {
		test.Compare(t)
	}
}
