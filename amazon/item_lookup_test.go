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

const expectedItemLookupSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Condition=New&IdType=ISBN&IncludeReviewsSummary=False&ItemId=foo%2Cbar%2Cbaz&MerchantId=Amazon&Operation=ItemLookup&RelatedItemPage=10&RelationshipType=DigitalMusicArranger&ResponseGroup=Accessories%2CBrowseNodes%2CEditorialReview%2CImages%2CItemAttributes%2CItemIds%2CLarge%2CMedium%2COfferFull%2COffers%2CPromotionSummary%2COfferSummary%2CRelatedItems%2CReviews%2CSalesRank%2CSimilarities%2CSmall%2CTracks%2CVariationImages%2CVariations%2CVariationSummary&SearchIndex=Books&Service=AWSECommerceService&Signature=wclnww68OCjD6o%2F5kLECcP0YRGiiODPpasHaaTyzc4Y%3D&Timestamp=2016-11-16T12%3A34%3A00Z&TruncateReviewsAt=10&VariationPage=2&Version=2013-08-01"

func createTestItemLookupRequest(client *Client) *ItemLookupRequest {
	return client.ItemLookup(ItemLookupParameters{
		ResponseGroups: []ItemLookupResponseGroup{
			ItemLookupResponseGroupAccessories,
			ItemLookupResponseGroupBrowseNodes,
			ItemLookupResponseGroupEditorialReview,
			ItemLookupResponseGroupImages,
			ItemLookupResponseGroupItemAttributes,
			ItemLookupResponseGroupItemIds,
			ItemLookupResponseGroupLarge,
			ItemLookupResponseGroupMedium,
			ItemLookupResponseGroupOfferFull,
			ItemLookupResponseGroupOffers,
			ItemLookupResponseGroupPromotionSummary,
			ItemLookupResponseGroupOfferSummary,
			ItemLookupResponseGroupRelatedItems,
			ItemLookupResponseGroupReviews,
			ItemLookupResponseGroupSalesRank,
			ItemLookupResponseGroupSimilarities,
			ItemLookupResponseGroupSmall,
			ItemLookupResponseGroupTracks,
			ItemLookupResponseGroupVariationImages,
			ItemLookupResponseGroupVariations,
			ItemLookupResponseGroupVariationSummary,
		},
		Condition:             ConditionNew,
		IDType:                IDTypeISBN,
		ItemIDs:               []string{"foo", "bar", "baz"},
		IncludeReviewsSummary: &[]bool{false}[0],
		MerchantID:            "Amazon",
		RelatedItemPage:       10,
		RelationshipType:      RelationshipTypeDigitalMusicArranger,
		SearchIndex:           SearchIndexBooks,
		TruncateReviewsAt:     &[]int{10}[0],
		VariationPage:         2,
	})
}

func TestItemLookupSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemLookupRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedItemLookupSignedURL, signedURL},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"New", parsed.Query().Get("Condition")},
		Test{"ISBN", parsed.Query().Get("IdType")},
		Test{"False", parsed.Query().Get("IncludeReviewsSummary")},
		Test{"foo,bar,baz", parsed.Query().Get("ItemId")},
		Test{"Amazon", parsed.Query().Get("MerchantId")},
		Test{"ItemLookup", parsed.Query().Get("Operation")},
		Test{"10", parsed.Query().Get("RelatedItemPage")},
		Test{"DigitalMusicArranger", parsed.Query().Get("RelationshipType")},
		Test{"Accessories,BrowseNodes,EditorialReview,Images,ItemAttributes,ItemIds,Large,Medium,OfferFull,Offers,PromotionSummary,OfferSummary,RelatedItems,Reviews,SalesRank,Similarities,Small,Tracks,VariationImages,Variations,VariationSummary", parsed.Query().Get("ResponseGroup")},
		Test{"Books", parsed.Query().Get("SearchIndex")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"wclnww68OCjD6o/5kLECcP0YRGiiODPpasHaaTyzc4Y=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"10", parsed.Query().Get("TruncateReviewsAt")},
		Test{"2", parsed.Query().Get("VariationPage")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestItemLookupDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/ItemLookupResponseErrorItem.xml")
	gock.New(strings.Replace(expectedItemLookupSignedURL, "%2B", "%5C%2B", -1)).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestItemLookupDoError(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemLookupRequest(client)
	gock.New(strings.Replace(expectedItemLookupSignedURL, "%2B", "%5C%2B", -1)).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedItemLookupSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestItemLookupDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/ItemLookup.xml")
	gock.New(expectedItemLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	// fmt.Println(res.Items.Item[0])
	for _, test := range []Test{
		Test{10, len(res.Items.Item)},
	} {
		test.Compare(t)
	}
}
