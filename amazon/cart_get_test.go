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

const expectedCartGetSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&CartId=351-9409673-0414064&CartItemId=CYOXSP4DGCG16&HMAC=%2Bak%2Bv8qGeDkkdQ%2Fw0o%2B5uA2heQI%3D&Operation=CartGet&ResponseGroup=Cart%2CCartNewReleases%2CCartSimilarities%2CCartTopSellers&Service=AWSECommerceService&Signature=KhpXLgLTedbMUSW3AnKGndR%2F6PHVTXHiGD5miR3VeHs%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createCartGetRequest(client *Client) *CartGetRequest {
	return client.CartGet(CartGetParameters{
		ResponseGroups: []CartGetResponseGroup{
			CartGetResponseGroupCart,
			CartGetResponseGroupCartNewReleases,
			CartGetResponseGroupCartSimilarities,
			CartGetResponseGroupCartTopSellers,
		},
		CartID:     "351-9409673-0414064",
		CartItemID: "CYOXSP4DGCG16",
		HMAC:       "+ak+v8qGeDkkdQ/w0o+5uA2heQI=",
	})
}

func TestCartGetSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartGetRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedCartGetSignedURL, signedURL},
		Test{"AK", parsed.Query().Get("AWSAccessKeyId")},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"+ak+v8qGeDkkdQ/w0o+5uA2heQI=", parsed.Query().Get("HMAC")},
		Test{"351-9409673-0414064", parsed.Query().Get("CartId")},
		Test{"CYOXSP4DGCG16", parsed.Query().Get("CartItemId")},
		Test{"CartGet", parsed.Query().Get("Operation")},
		Test{"Cart,CartNewReleases,CartSimilarities,CartTopSellers", parsed.Query().Get("ResponseGroup")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"KhpXLgLTedbMUSW3AnKGndR/6PHVTXHiGD5miR3VeHs=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestCartGetDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartGetRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartGetResponseErrorItem.xml")
	gock.New(strings.Replace(expectedCartGetSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、CartIdなどがあります。", err.Error()}.Compare(t)
	}
}

func TestCartGetDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartGetRequest(client)
	gock.New(strings.Replace(expectedCartGetSignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedCartGetSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestCartGetDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createCartGetRequest(client)
	fixtureIO, _ := os.Open("_fixtures/CartGet.xml")
	gock.New(strings.Replace(expectedCartGetSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{"354-9729779-1559716", res.Cart.ID},
		// TODO: write more tests
	} {
		test.Compare(t)
	}
}
