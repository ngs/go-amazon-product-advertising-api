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

const expectedCartClearSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&CartId=351-9409673-0414064&HMAC=%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D&Operation=CartClear&ResponseGroup=Cart&Service=AWSECommerceService&Signature=IhwJobQIwyjG13bmsehXb6znZvOj8Iz4OGtVaffotMo%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartClearRequest(client *Client) *CartClearRequest {
	return client.CartClear(CartClearParameters{
		ResponseGroups: []CartClearResponseGroup{
			CartClearResponseGroupCart,
		},
		CartID: "351-9409673-0414064",
		HMAC:   "+ak+v8qGeDkkdQ/w0o+5uA2heQI=",
	})
}

func TestCartClearSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartClearRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		{expectedCartClearSignedURL, signedURL},
		{"AK", parsed.Query().Get("AWSAccessKeyId")},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"+ak+v8qGeDkkdQ/w0o+5uA2heQI=", parsed.Query().Get("HMAC")},
		{"351-9409673-0414064", parsed.Query().Get("CartId")},
		{"CartClear", parsed.Query().Get("Operation")},
		{"Cart", parsed.Query().Get("ResponseGroup")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"IhwJobQIwyjG13bmsehXb6znZvOj8Iz4OGtVaffotMo=", parsed.Query().Get("Signature")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartClearDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartClearRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartClearResponseErrorItem.xml")
	gock.New(strings.Replace(expectedCartClearSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、CartIdなどがあります。", err.Error()}.Compare(t)
	}
}

func TestCartClearDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartClearRequest(client)
	gock.New(strings.Replace(expectedCartClearSignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartClearSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartClearDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartClearRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartClear.xml")
	gock.New(strings.Replace(expectedCartClearSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		{"353-7649034-4766017", res.Cart.ID},
	} {
		test.Compare(t)
	}
}
