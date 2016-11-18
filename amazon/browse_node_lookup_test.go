package amazon

import (
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

const expectedBrowseNodeLookupSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&BrowseNodeId=492352&Operation=BrowseNodeLookup&ResponseGroup=BrowseNodeInfo%2CNewReleases%2CMostGifted%2CTopSellers%2CMostWishedFor&Service=AWSECommerceService&Signature=DsQWqEYtxLtHgkHCiLmyTF7geOrNNqkvznh3KSplb7w%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createBrowseNodeLookupRequest(client *Client) *BrowseNodeLookupRequest {
	return client.BrowseNodeLookup(BrowseNodeLookupParameters{
		ResponseGroups: []BrowseNodeLookupResponseGroup{
			BrowseNodeLookupResponseGroupBrowseNodeInfo,
			BrowseNodeLookupResponseGroupNewReleases,
			BrowseNodeLookupResponseGroupMostGifted,
			BrowseNodeLookupResponseGroupTopSellers,
			BrowseNodeLookupResponseGroupMostWishedFor,
		},
		BrowseNodeID: "492352",
	})
}

func TestBrowseNodeLookupSignedURL(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedBrowseNodeLookupSignedURL, signedURL},
		Test{"AK", parsed.Query().Get("AWSAccessKeyId")},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"492352", parsed.Query().Get("BrowseNodeId")},
		Test{"BrowseNodeLookup", parsed.Query().Get("Operation")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"DsQWqEYtxLtHgkHCiLmyTF7geOrNNqkvznh3KSplb7w=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestBrowseNodeLookupDoErrorResponse(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/BrowseNodeLookupResponseErrorItem.xml")
	gock.New(expectedBrowseNodeLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestBrowseNodeLookupDoError(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	gock.New(expectedBrowseNodeLookupSignedURL).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedBrowseNodeLookupSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestBrowseNodeLookupDo(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/BrowseNodeLookup.xml")
	gock.New(expectedBrowseNodeLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{1, len(res.BrowseNodes())},
		Test{"492352", res.BrowseNodes()[0].ID},
		Test{"プログラミング", res.BrowseNodes()[0].Name},
		Test{5, len(res.BrowseNodes()[0].Children.BrowseNode)},
		Test{1, len(res.BrowseNodes()[0].Ancestors.BrowseNode)},
	} {
		test.Compare(t)
	}
}
