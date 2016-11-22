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

const expectedSimilarityLookupSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Condition=New&ItemId=foo%2Cbar%2Cbaz&MerchantId=Amazon&Operation=SimilarityLookup&ResponseGroup=Accessories%2CBrowseNodes%2CEditorialReview%2CImages%2CLarge%2CItemAttributes%2CItemIds%2CMedium%2COffers%2COfferSummary%2CPromotionSummary%2CReviews%2CSalesRank%2CSimilarities%2CSmall%2CTracks%2CVariations%2CVariationSummary&Service=AWSECommerceService&Signature=yQtyToXWmhs80nzcFa5RBmb1MDv4Wi0wFeNKruxNw2A%3D&SimilarityType=Intersection&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createTestSimilarityLookupRequest(client *Client) *SimilarityLookupRequest {
	return client.SimilarityLookup(SimilarityLookupParameters{
		ResponseGroups: []SimilarityLookupResponseGroup{
			SimilarityLookupResponseGroupAccessories,
			SimilarityLookupResponseGroupBrowseNodes,
			SimilarityLookupResponseGroupEditorialReview,
			SimilarityLookupResponseGroupImages,
			SimilarityLookupResponseGroupLarge,
			SimilarityLookupResponseGroupItemAttributes,
			SimilarityLookupResponseGroupItemIds,
			SimilarityLookupResponseGroupMedium,
			SimilarityLookupResponseGroupOffers,
			SimilarityLookupResponseGroupOfferSummary,
			SimilarityLookupResponseGroupPromotionSummary,
			SimilarityLookupResponseGroupReviews,
			SimilarityLookupResponseGroupSalesRank,
			SimilarityLookupResponseGroupSimilarities,
			SimilarityLookupResponseGroupSmall,
			SimilarityLookupResponseGroupTracks,
			SimilarityLookupResponseGroupVariations,
			SimilarityLookupResponseGroupVariationSummary,
		},
		Condition:      ConditionNew,
		ItemIDs:        []string{"foo", "bar", "baz"},
		MerchantID:     "Amazon",
		SimilarityType: SimilarityTypeIntersection,
	})
}

func TestSimilarityLookupSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestSimilarityLookupRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		{expectedSimilarityLookupSignedURL, signedURL},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"AK", parsed.Query().Get("AWSAccessKeyId")},
		{"ngsio-22", parsed.Query().Get("AssociateTag")},
		{"New", parsed.Query().Get("Condition")},
		{"foo,bar,baz", parsed.Query().Get("ItemId")},
		{"Amazon", parsed.Query().Get("MerchantId")},
		{"SimilarityLookup", parsed.Query().Get("Operation")},
		{"Accessories,BrowseNodes,EditorialReview,Images,Large,ItemAttributes,ItemIds,Medium,Offers,OfferSummary,PromotionSummary,Reviews,SalesRank,Similarities,Small,Tracks,Variations,VariationSummary", parsed.Query().Get("ResponseGroup")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"yQtyToXWmhs80nzcFa5RBmb1MDv4Wi0wFeNKruxNw2A=", parsed.Query().Get("Signature")},
		{"Intersection", parsed.Query().Get("SimilarityType")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestSimilarityLookupDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestSimilarityLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/SimilarityLookupResponseErrorItem.xml")
	gock.New(strings.Replace(expectedSimilarityLookupSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestSimilarityLookupDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestSimilarityLookupRequest(client)
	gock.New(strings.Replace(expectedSimilarityLookupSignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedSimilarityLookupSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestSimilarityLookupDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestSimilarityLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/SimilarityLookup.xml")
	gock.New(expectedSimilarityLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	// fmt.Println(res.Items.Item[0])
	for _, test := range []Test{
		{10, len(res.Items.Item)},
	} {
		test.Compare(t)
	}
}
