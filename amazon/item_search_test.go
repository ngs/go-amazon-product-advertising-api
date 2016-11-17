package amazon

import (
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

const expectedItemSearchSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&Actor=Actor&Artist=Artist&AssociateTag=ngsio-22&AudienceRating=AudienceRating&Author=Author&Availability=Available&Brand=Brand&BrowseNode=BrowseNode&Composer=Composer&Condition=Condition&Conductor=Conductor&Director=Director&IncludeReviewsSummary=True&ItemPage=100&Keywords=Keywords&Manufacturer=Manufacturer&MaximumPrice=1000&MerchantID=MerchantID&MinPercentageOff=20&MinimumPrice=10&Operation=ItemSearch&Orchestra=Orchestra&Power=Power&Publisher=Publisher&RelatedItemPage=50&RelationshipType=RelationshipType&ResponseGroup=Accessories%2CAlternateVersions%2CBrowseNodes%2CEditorialReview%2CImages%2CItemAttributesItemIds%2CLarge%2CListmaniaLists%2CMedium%2CMerchantItemAttributes%2COfferFull%2COfferListings%2COffers%2COfferSummary%2CPromotionalTag%2CPromotionDetails%2CPromotionSummary%2CRelatedItems%2CReviews%2CSalesRank%2CSearchBins%2CSearchInside%2CSimilarities%2CSmall%2CSubjects%2CTracks%2CVariationMatrix%2CVariationMinimum%2CVariationOffers%2CVariations%2CVariationSummary&SearchIndex=Automotive&Service=AWSECommerceService&Signature=YaOeE3uFyK0g4hNInK%2F8AZDPAchPjpZxm2y57rQVp3o%3D&Sort=Sort&Timestamp=2016-11-16T12%3A34%3A00Z&Title=Title&TruncateReviewsAt=60&VariationPage=30&Version=2013-08-01"

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
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.ItemSearch(ItemSearchParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestItemSearchBuildQuery(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	Test{
		expectedItemSearchSignedURL,
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
		Test{"YaOeE3uFyK0g4hNInK/8AZDPAchPjpZxm2y57rQVp3o=", parsed.Query().Get("Signature")},
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

func TestItemSearchDoErrorResponse(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestRequest(client)
	fixtureIO, _ := os.Open("_fixtures/ItemSearchResponseErrorItem.xml")
	gock.New(expectedItemSearchSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestItemSearchDoError(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestRequest(client)
	gock.New(expectedItemSearchSignedURL).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedItemSearchSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestItemSearchDo(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestRequest(client)
	fixtureIO, _ := os.Open("_fixtures/ItemSearchResponse.xml")
	gock.New(expectedItemSearchSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{190, res.TotalResults()},
		Test{19, res.TotalPages()},
		Test{"https://www.amazon.co.jp/gp/redirect.html?linkCode=xm2&SubscriptionId=AKIAITPH62XKCOOT7AKA&location=https%3A%2F%2Fwww.amazon.co.jp%2Fgp%2Fsearch%3Fkeywords%3DGo%2B%25E8%25A8%2580%25E8%25AA%259E%26url%3Dsearch-alias%253Dbooks-single-index&tag=atsushnagased-22&creative=5143&camp=2025", res.MoreSearchResultsURL()},
		Test{10, len(res.Items())},
		Test{"4621300253", res.Items()[0].ASIN},
		Test{"https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4621300253", res.Items()[0].DetailPageURL},
		Test{12609, res.Items()[0].SalesRank},
		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items()[0].SmallImage.URL},
		Test{59, res.Items()[0].SmallImage.Width.Value},
		Test{"pixels", res.Items()[0].SmallImage.Width.Units},
		Test{75, res.Items()[0].SmallImage.Height.Value},
		Test{"pixels", res.Items()[0].SmallImage.Height.Units},
		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg", res.Items()[0].MediumImage.URL},
		Test{127, res.Items()[0].MediumImage.Width.Value},
		Test{"pixels", res.Items()[0].MediumImage.Width.Units},
		Test{160, res.Items()[0].MediumImage.Height.Value},
		Test{"pixels", res.Items()[0].MediumImage.Height.Units},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L.jpg", res.Items()[0].LargeImage.URL},
		Test{396, res.Items()[0].LargeImage.Width.Value},
		Test{"pixels", res.Items()[0].LargeImage.Width.Units},

		Test{500, res.Items()[0].LargeImage.Height.Value},
		Test{"pixels", res.Items()[0].LargeImage.Height.Units},

		Test{1, len(res.Items()[0].ImageSets.ImageSet)},
		Test{"primary", res.Items()[0].ImageSets.ImageSet[0].Category},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL30_.jpg", res.Items()[0].ImageSets.ImageSet[0].SwatchImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].SwatchImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].SwatchImage.Width.Units},
		Test{30, res.Items()[0].ImageSets.ImageSet[0].SwatchImage.Height.Value},
		Test{24, res.Items()[0].ImageSets.ImageSet[0].SwatchImage.Width.Value},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items()[0].ImageSets.ImageSet[0].SmallImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].SmallImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].SmallImage.Width.Units},
		Test{75, res.Items()[0].ImageSets.ImageSet[0].SmallImage.Height.Value},
		Test{59, res.Items()[0].ImageSets.ImageSet[0].SmallImage.Width.Value},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items()[0].ImageSets.ImageSet[0].ThumbnailImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].ThumbnailImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].ThumbnailImage.Width.Units},
		Test{75, res.Items()[0].ImageSets.ImageSet[0].ThumbnailImage.Height.Value},
		Test{59, res.Items()[0].ImageSets.ImageSet[0].ThumbnailImage.Width.Value},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL110_.jpg", res.Items()[0].ImageSets.ImageSet[0].TinyImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].TinyImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].TinyImage.Width.Units},
		Test{110, res.Items()[0].ImageSets.ImageSet[0].TinyImage.Height.Value},
		Test{87, res.Items()[0].ImageSets.ImageSet[0].TinyImage.Width.Value},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg", res.Items()[0].ImageSets.ImageSet[0].MediumImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].MediumImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].MediumImage.Width.Units},
		Test{160, res.Items()[0].ImageSets.ImageSet[0].MediumImage.Height.Value},
		Test{127, res.Items()[0].ImageSets.ImageSet[0].MediumImage.Width.Value},

		Test{"http://ecx.images-amazon.com/images/I/410V3ulwP5L.jpg", res.Items()[0].ImageSets.ImageSet[0].LargeImage.URL},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].LargeImage.Height.Units},
		Test{"pixels", res.Items()[0].ImageSets.ImageSet[0].LargeImage.Width.Units},
		Test{500, res.Items()[0].ImageSets.ImageSet[0].LargeImage.Height.Value},
		Test{396, res.Items()[0].ImageSets.ImageSet[0].LargeImage.Width.Value},

		Test{4, len(res.Items()[0].ItemLinks.ItemLink)},
		Test{"Add To Wishlist", res.Items()[0].ItemLinks.ItemLink[0].Description},
		Test{"https://www.amazon.co.jp/gp/registry/wishlist/add-item.html%3Fasin.0%3D4621300253%26SubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items()[0].ItemLinks.ItemLink[0].URL},

		Test{"Tell A Friend", res.Items()[0].ItemLinks.ItemLink[1].Description},
		Test{"https://www.amazon.co.jp/gp/pdp/taf/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items()[0].ItemLinks.ItemLink[1].URL},

		Test{"All Customer Reviews", res.Items()[0].ItemLinks.ItemLink[2].Description},
		Test{"https://www.amazon.co.jp/review/product/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items()[0].ItemLinks.ItemLink[2].URL},

		Test{"All Offers", res.Items()[0].ItemLinks.ItemLink[3].Description},
		Test{"https://www.amazon.co.jp/gp/offer-listing/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items()[0].ItemLinks.ItemLink[3].URL},

		Test{2, len(res.Items()[0].ItemAttributes.Author)},
		Test{"Alan A.A. Donovan", res.Items()[0].ItemAttributes.Author[0]},
		Test{"Brian W. Kernighan", res.Items()[0].ItemAttributes.Author[1]},
		Test{"単行本（ソフトカバー）", res.Items()[0].ItemAttributes.Binding},
		Test{"翻訳", res.Items()[0].ItemAttributes.Creator.Role},
		Test{"柴田 芳樹", res.Items()[0].ItemAttributes.Creator.Name},
		Test{"9784621300251", res.Items()[0].ItemAttributes.EAN},
		Test{1, len(res.Items()[0].ItemAttributes.EANList.Element)},
		Test{"9784621300251", res.Items()[0].ItemAttributes.EANList.Element[0]},
		Test{false, res.Items()[0].ItemAttributes.IsAdultProduct},
		Test{"4621300253", res.Items()[0].ItemAttributes.ISBN},
		Test{"丸善出版", res.Items()[0].ItemAttributes.Label},
		Test{2, len(res.Items()[0].ItemAttributes.Languages.Language)},
		Test{"日本語", res.Items()[0].ItemAttributes.Languages.Language[0].Name},
		Test{"Published", res.Items()[0].ItemAttributes.Languages.Language[0].Type},
		Test{"日本語", res.Items()[0].ItemAttributes.Languages.Language[1].Name},
		Test{"Unknown", res.Items()[0].ItemAttributes.Languages.Language[1].Type},
		Test{"4104", res.Items()[0].ItemAttributes.ListPrice.Amount},
		Test{"JPY", res.Items()[0].ItemAttributes.ListPrice.CurrencyCode},
		Test{"￥ 4,104", res.Items()[0].ItemAttributes.ListPrice.FormattedPrice},
		Test{"丸善出版", res.Items()[0].ItemAttributes.Manufacturer},
		Test{462, res.Items()[0].ItemAttributes.NumberOfPages},
		Test{"hundredths-inches", res.Items()[0].ItemAttributes.PackageDimensions.Width.Units},
		Test{732, res.Items()[0].ItemAttributes.PackageDimensions.Width.Value},
		Test{"hundredths-inches", res.Items()[0].ItemAttributes.PackageDimensions.Height.Units},
		Test{110, res.Items()[0].ItemAttributes.PackageDimensions.Height.Value},
		Test{"hundredths-pounds", res.Items()[0].ItemAttributes.PackageDimensions.Weight.Units},
		Test{171, res.Items()[0].ItemAttributes.PackageDimensions.Weight.Value},
		Test{"hundredths-inches", res.Items()[0].ItemAttributes.PackageDimensions.Length.Units},
		Test{909, res.Items()[0].ItemAttributes.PackageDimensions.Length.Value},
		Test{"Book", res.Items()[0].ItemAttributes.ProductGroup},
		Test{"ABIS_BOOK", res.Items()[0].ItemAttributes.ProductTypeName},
		Test{"丸善出版", res.Items()[0].ItemAttributes.Publisher},
		Test{"丸善出版", res.Items()[0].ItemAttributes.Studio},
		Test{"プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)", res.Items()[0].ItemAttributes.Title},
		Test{time.Date(2016, 6, 20, 0, 0, 0, 0, time.UTC).UnixNano(), res.Items()[0].ItemAttributes.PublicationDate.UnixNano()},
	} {
		test.Compare(t)
	}
	// fmt.Printf("res %v\n", res)
}
