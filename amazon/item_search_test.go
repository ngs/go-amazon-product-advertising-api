package amazon

import (
	"net/url"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

func createTestRequest(client *Client) *ItemSearchRequest {
	return client.ItemSearch(ItemSearchParameters{
		Actor:                 "Actor",
		Artist:                "Artist",
		AudienceRating:        "AudienceRating",
		Author:                "Author",
		OnlyAvailable:         true,
		Brand:                 "Brand",
		BrowseNode:            "BrowseNode",
		Composer:              "Composer",
		Condition:             "Condition",
		Conductor:             "Conductor",
		Director:              "Director",
		IncludeReviewsSummary: &[]bool{true}[0],
		ItemPage:              100,
		Keywords:              "Keywords",
		Manufacturer:          "Manufacturer",
		MaximumPrice:          1000,
		MerchantID:            "MerchantID",
		MinimumPrice:          10,
		MinPercentageOff:      20,
		Orchestra:             "Orchestra",
		Power:                 "Power",
		Publisher:             "Publisher",
		RelatedItemPage:       50,
		RelationshipType:      "RelationshipType",
		SearchIndex:           SearchIndexAutomotive,
		Sort:                  "Sort",
		Title:                 "Title",
		TruncateReviewsAt:     &[]int{60}[0],
		VariationPage:         &[]int{30}[0],
		ResponseGroups: []ItemSearchResponseGroup{
			ItemSearchResponseGroupAccessories,
			ItemSearchResponseGroupAlternateVersions,
			ItemSearchResponseGroupBrowseNodes,
			ItemSearchResponseGroupEditorialReview,
			ItemSearchResponseGroupImages,
			ItemSearchResponseGroupItemAttributesItemIds,
			ItemSearchResponseGroupLarge,
			ItemSearchResponseGroupListmaniaLists,
			ItemSearchResponseGroupMedium,
			ItemSearchResponseGroupMerchantItemAttributes,
			ItemSearchResponseGroupOfferFull,
			ItemSearchResponseGroupOfferListings,
			ItemSearchResponseGroupOffers,
			ItemSearchResponseGroupOfferSummary,
			ItemSearchResponseGroupPromotionalTag,
			ItemSearchResponseGroupPromotionDetails,
			ItemSearchResponseGroupPromotionSummary,
			ItemSearchResponseGroupRelatedItems,
			ItemSearchResponseGroupReviews,
			ItemSearchResponseGroupSalesRank,
			ItemSearchResponseGroupSearchBins,
			ItemSearchResponseGroupSearchInside,
			ItemSearchResponseGroupSimilarities,
			ItemSearchResponseGroupSmall,
			ItemSearchResponseGroupSubjects,
			ItemSearchResponseGroupTracks,
			ItemSearchResponseGroupVariationMatrix,
			ItemSearchResponseGroupVariationMinimum,
			ItemSearchResponseGroupVariationOffers,
			ItemSearchResponseGroupVariations,
			ItemSearchResponseGroupVariationSummary,
		},
	})
}

func TestItemSearch(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.ItemSearch(ItemSearchParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestItemSearchBuildQuery(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "JP")
	op := createTestRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	Test{
		"https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&Actor=Actor&Artist=Artist&AudienceRating=AudienceRating&Author=Author&Availability=Available&Brand=Brand&BrowseNode=BrowseNode&Composer=Composer&Condition=Condition&Conductor=Conductor&Director=Director&IncludeReviewsSummary=True&ItemPage=100&Keywords=Keywords&Manufacturer=Manufacturer&MaximumPrice=1000&MerchantID=MerchantID&MinPercentageOff=20&MinimumPrice=10&Operation=ItemSearch&Orchestra=Orchestra&Power=Power&Publisher=Publisher&RelatedItemPage=50&RelationshipType=RelationshipType&ResponseGroup=Accessories%2CAlternateVersions%2CBrowseNodes%2CEditorialReview%2CImages%2CItemAttributesItemIds%2CLarge%2CListmaniaLists%2CMedium%2CMerchantItemAttributes%2COfferFull%2COfferListings%2COffers%2COfferSummary%2CPromotionalTag%2CPromotionDetails%2CPromotionSummary%2CRelatedItems%2CReviews%2CSalesRank%2CSearchBins%2CSearchInside%2CSimilarities%2CSmall%2CSubjects%2CTracks%2CVariationMatrix%2CVariationMinimum%2CVariationOffers%2CVariations%2CVariationSummary&SearchIndex=Automotive&Service=AWSECommerceService&Signature=up4lYm_xGBEt8JRSVU8LaSVK13JC2q185vvA3PmyrV4%3D&Sort=Sort&Timestamp=2016-11-16T12%3A34%3A00Z&Title=Title&TruncateReviewsAt=60&VariationPage=30&Version=2013-08-01",
		signedURL,
	}.Compare(t)
	for _, test := range []Test{
		Test{"Actor", parsed.Query().Get("Actor")},
		Test{"Artist", parsed.Query().Get("Artist")},
		Test{"AudienceRating", parsed.Query().Get("AudienceRating")},
		Test{"Author", parsed.Query().Get("Author")},
		Test{"Available", parsed.Query().Get("Availability")},
		Test{"Brand", parsed.Query().Get("Brand")},
		Test{"BrowseNode", parsed.Query().Get("BrowseNode")},
		Test{"Composer", parsed.Query().Get("Composer")},
		Test{"Condition", parsed.Query().Get("Condition")},
		Test{"Conductor", parsed.Query().Get("Conductor")},
		Test{"Director", parsed.Query().Get("Director")},
		Test{"True", parsed.Query().Get("IncludeReviewsSummary")},
		Test{"100", parsed.Query().Get("ItemPage")},
		Test{"Keywords", parsed.Query().Get("Keywords")},
		Test{"Manufacturer", parsed.Query().Get("Manufacturer")},
		Test{"1000", parsed.Query().Get("MaximumPrice")},
		Test{"MerchantID", parsed.Query().Get("MerchantID")},
		Test{"20", parsed.Query().Get("MinPercentageOff")},
		Test{"10", parsed.Query().Get("MinimumPrice")},
		Test{"ItemSearch", parsed.Query().Get("Operation")},
		Test{"Orchestra", parsed.Query().Get("Orchestra")},
		Test{"Power", parsed.Query().Get("Power")},
		Test{"Publisher", parsed.Query().Get("Publisher")},
		Test{"50", parsed.Query().Get("RelatedItemPage")},
		Test{"RelationshipType", parsed.Query().Get("RelationshipType")},
		Test{"Accessories,AlternateVersions,BrowseNodes,EditorialReview,Images,ItemAttributesItemIds,Large,ListmaniaLists,Medium,MerchantItemAttributes,OfferFull,OfferListings,Offers,OfferSummary,PromotionalTag,PromotionDetails,PromotionSummary,RelatedItems,Reviews,SalesRank,SearchBins,SearchInside,Similarities,Small,Subjects,Tracks,VariationMatrix,VariationMinimum,VariationOffers,Variations,VariationSummary", parsed.Query().Get("ResponseGroup")},
		Test{"Automotive", parsed.Query().Get("SearchIndex")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"up4lYm_xGBEt8JRSVU8LaSVK13JC2q185vvA3PmyrV4=", parsed.Query().Get("Signature")},
		Test{"Sort", parsed.Query().Get("Sort")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"Title", parsed.Query().Get("Title")},
		Test{"60", parsed.Query().Get("TruncateReviewsAt")},
		Test{"30", parsed.Query().Get("VariationPage")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestItemSearchDo(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "JP")
	op := createTestRequest(client)
	gock.New("https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&Actor=Actor&Artist=Artist&AudienceRating=AudienceRating&Author=Author&Availability=Available&Brand=Brand&BrowseNode=BrowseNode&Composer=Composer&Condition=Condition&Conductor=Conductor&Director=Director&IncludeReviewsSummary=True&ItemPage=100&Keywords=Keywords&Manufacturer=Manufacturer&MaximumPrice=1000&MerchantID=MerchantID&MinPercentageOff=20&MinimumPrice=10&Operation=ItemSearch&Orchestra=Orchestra&Power=Power&Publisher=Publisher&RelatedItemPage=50&RelationshipType=RelationshipType&ResponseGroup=Accessories%2CAlternateVersions%2CBrowseNodes%2CEditorialReview%2CImages%2CItemAttributesItemIds%2CLarge%2CListmaniaLists%2CMedium%2CMerchantItemAttributes%2COfferFull%2COfferListings%2COffers%2COfferSummary%2CPromotionalTag%2CPromotionDetails%2CPromotionSummary%2CRelatedItems%2CReviews%2CSalesRank%2CSearchBins%2CSearchInside%2CSimilarities%2CSmall%2CSubjects%2CTracks%2CVariationMatrix%2CVariationMinimum%2CVariationOffers%2CVariations%2CVariationSummary&SearchIndex=Automotive&Service=AWSECommerceService&Signature=up4lYm_xGBEt8JRSVU8LaSVK13JC2q185vvA3PmyrV4%3D&Sort=Sort&Timestamp=2016-11-16T12%3A34%3A00Z&Title=Title&TruncateReviewsAt=60&VariationPage=30&Version=2013-08-01").
		Reply(200).
		BodyString("ok") // TODO
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	t.Skipf("Not yet implemented: %v", res)
}
