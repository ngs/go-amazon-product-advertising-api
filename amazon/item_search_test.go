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

func createTestItemSearchRequest(client *Client) *ItemSearchRequest {
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

func TestItemSearchSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemSearchRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		{expectedItemSearchSignedURL, signedURL},
		{"Actor", parsed.Query().Get("Actor")},
		{"Artist", parsed.Query().Get("Artist")},
		{"AudienceRating", parsed.Query().Get("AudienceRating")},
		{"Author", parsed.Query().Get("Author")},
		{"Available", parsed.Query().Get("Availability")},
		{"Brand", parsed.Query().Get("Brand")},
		{"BrowseNode", parsed.Query().Get("BrowseNode")},
		{"Composer", parsed.Query().Get("Composer")},
		{"Condition", parsed.Query().Get("Condition")},
		{"Conductor", parsed.Query().Get("Conductor")},
		{"Director", parsed.Query().Get("Director")},
		{"True", parsed.Query().Get("IncludeReviewsSummary")},
		{"100", parsed.Query().Get("ItemPage")},
		{"Keywords", parsed.Query().Get("Keywords")},
		{"Manufacturer", parsed.Query().Get("Manufacturer")},
		{"1000", parsed.Query().Get("MaximumPrice")},
		{"MerchantID", parsed.Query().Get("MerchantID")},
		{"20", parsed.Query().Get("MinPercentageOff")},
		{"10", parsed.Query().Get("MinimumPrice")},
		{"ItemSearch", parsed.Query().Get("Operation")},
		{"Orchestra", parsed.Query().Get("Orchestra")},
		{"Power", parsed.Query().Get("Power")},
		{"Publisher", parsed.Query().Get("Publisher")},
		{"50", parsed.Query().Get("RelatedItemPage")},
		{"RelationshipType", parsed.Query().Get("RelationshipType")},
		{"Accessories,AlternateVersions,BrowseNodes,EditorialReview,Images,ItemAttributesItemIds,Large,ListmaniaLists,Medium,MerchantItemAttributes,OfferFull,OfferListings,Offers,OfferSummary,PromotionalTag,PromotionDetails,PromotionSummary,RelatedItems,Reviews,SalesRank,SearchBins,SearchInside,Similarities,Small,Subjects,Tracks,VariationMatrix,VariationMinimum,VariationOffers,Variations,VariationSummary", parsed.Query().Get("ResponseGroup")},
		{"Automotive", parsed.Query().Get("SearchIndex")},
		{"AWSECommerceService", parsed.Query().Get("Service")},
		{"YaOeE3uFyK0g4hNInK/8AZDPAchPjpZxm2y57rQVp3o=", parsed.Query().Get("Signature")},
		{"Sort", parsed.Query().Get("Sort")},
		{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		{"Title", parsed.Query().Get("Title")},
		{"60", parsed.Query().Get("TruncateReviewsAt")},
		{"30", parsed.Query().Get("VariationPage")},
		{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestItemSearchDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemSearchRequest(client)
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
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemSearchRequest(client)
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
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createTestItemSearchRequest(client)
	fixtureIO, _ := os.Open("_fixtures/ItemSearch.xml")
	gock.New(expectedItemSearchSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	// fmt.Println(res.Items.Item[0])
	for _, test := range []Test{
		{190, res.Items.TotalResults},
		{19, res.Items.TotalPages},
		{"https://www.amazon.co.jp/gp/redirect.html?linkCode=xm2&SubscriptionId=AKIAITPH62XKCOOT7AKA&location=https%3A%2F%2Fwww.amazon.co.jp%2Fgp%2Fsearch%3Fkeywords%3DGo%2B%25E8%25A8%2580%25E8%25AA%259E%26url%3Dsearch-alias%253Dbooks-single-index&tag=atsushnagased-22&creative=5143&camp=2025", res.Items.MoreSearchResultsURL},
		{10, len(res.Items.Item)},
		{"4621300253", res.Items.Item[0].ASIN},
		{"https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4621300253", res.Items.Item[0].DetailPageURL},
		{14521, res.Items.Item[0].SalesRank},
		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items.Item[0].SmallImage.URL},
		{59, res.Items.Item[0].SmallImage.Width.Value},
		{"pixels", res.Items.Item[0].SmallImage.Width.Units},
		{75, res.Items.Item[0].SmallImage.Height.Value},
		{"pixels", res.Items.Item[0].SmallImage.Height.Units},
		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg", res.Items.Item[0].MediumImage.URL},
		{127, res.Items.Item[0].MediumImage.Width.Value},
		{"pixels", res.Items.Item[0].MediumImage.Width.Units},
		{160, res.Items.Item[0].MediumImage.Height.Value},
		{"pixels", res.Items.Item[0].MediumImage.Height.Units},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L.jpg", res.Items.Item[0].LargeImage.URL},
		{396, res.Items.Item[0].LargeImage.Width.Value},
		{"pixels", res.Items.Item[0].LargeImage.Width.Units},

		{500, res.Items.Item[0].LargeImage.Height.Value},
		{"pixels", res.Items.Item[0].LargeImage.Height.Units},

		{1, len(res.Items.Item[0].ImageSets.ImageSet)},
		{"primary", res.Items.Item[0].ImageSets.ImageSet[0].Category},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL30_.jpg", res.Items.Item[0].ImageSets.ImageSet[0].SwatchImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].SwatchImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].SwatchImage.Width.Units},
		{30, res.Items.Item[0].ImageSets.ImageSet[0].SwatchImage.Height.Value},
		{24, res.Items.Item[0].ImageSets.ImageSet[0].SwatchImage.Width.Value},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items.Item[0].ImageSets.ImageSet[0].SmallImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].SmallImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].SmallImage.Width.Units},
		{75, res.Items.Item[0].ImageSets.ImageSet[0].SmallImage.Height.Value},
		{59, res.Items.Item[0].ImageSets.ImageSet[0].SmallImage.Width.Value},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL75_.jpg", res.Items.Item[0].ImageSets.ImageSet[0].ThumbnailImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].ThumbnailImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].ThumbnailImage.Width.Units},
		{75, res.Items.Item[0].ImageSets.ImageSet[0].ThumbnailImage.Height.Value},
		{59, res.Items.Item[0].ImageSets.ImageSet[0].ThumbnailImage.Width.Value},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL110_.jpg", res.Items.Item[0].ImageSets.ImageSet[0].TinyImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].TinyImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].TinyImage.Width.Units},
		{110, res.Items.Item[0].ImageSets.ImageSet[0].TinyImage.Height.Value},
		{87, res.Items.Item[0].ImageSets.ImageSet[0].TinyImage.Width.Value},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg", res.Items.Item[0].ImageSets.ImageSet[0].MediumImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].MediumImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].MediumImage.Width.Units},
		{160, res.Items.Item[0].ImageSets.ImageSet[0].MediumImage.Height.Value},
		{127, res.Items.Item[0].ImageSets.ImageSet[0].MediumImage.Width.Value},

		{"http://ecx.images-amazon.com/images/I/410V3ulwP5L.jpg", res.Items.Item[0].ImageSets.ImageSet[0].LargeImage.URL},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].LargeImage.Height.Units},
		{"pixels", res.Items.Item[0].ImageSets.ImageSet[0].LargeImage.Width.Units},
		{500, res.Items.Item[0].ImageSets.ImageSet[0].LargeImage.Height.Value},
		{396, res.Items.Item[0].ImageSets.ImageSet[0].LargeImage.Width.Value},

		{4, len(res.Items.Item[0].ItemLinks.ItemLink)},
		{"Add To Wishlist", res.Items.Item[0].ItemLinks.ItemLink[0].Description},
		{"https://www.amazon.co.jp/gp/registry/wishlist/add-item.html%3Fasin.0%3D4621300253%26SubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items.Item[0].ItemLinks.ItemLink[0].URL},

		{"Tell A Friend", res.Items.Item[0].ItemLinks.ItemLink[1].Description},
		{"https://www.amazon.co.jp/gp/pdp/taf/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items.Item[0].ItemLinks.ItemLink[1].URL},

		{"All Customer Reviews", res.Items.Item[0].ItemLinks.ItemLink[2].Description},
		{"https://www.amazon.co.jp/review/product/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items.Item[0].ItemLinks.ItemLink[2].URL},

		{"All Offers", res.Items.Item[0].ItemLinks.ItemLink[3].Description},
		{"https://www.amazon.co.jp/gp/offer-listing/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items.Item[0].ItemLinks.ItemLink[3].URL},

		{2, len(res.Items.Item[0].ItemAttributes.Author)},
		{"Alan A.A. Donovan", res.Items.Item[0].ItemAttributes.Author[0]},
		{"Brian W. Kernighan", res.Items.Item[0].ItemAttributes.Author[1]},
		{"単行本（ソフトカバー）", res.Items.Item[0].ItemAttributes.Binding},
		{"翻訳", res.Items.Item[0].ItemAttributes.Creator.Role},
		{"柴田 芳樹", res.Items.Item[0].ItemAttributes.Creator.Name},
		{"9784621300251", res.Items.Item[0].ItemAttributes.EAN},
		{1, len(res.Items.Item[0].ItemAttributes.EANList.Element)},
		{"9784621300251", res.Items.Item[0].ItemAttributes.EANList.Element[0]},
		{false, res.Items.Item[0].ItemAttributes.IsAdultProduct},
		{"4621300253", res.Items.Item[0].ItemAttributes.ISBN},
		{"丸善出版", res.Items.Item[0].ItemAttributes.Label},
		{2, len(res.Items.Item[0].ItemAttributes.Languages.Language)},
		{"日本語", res.Items.Item[0].ItemAttributes.Languages.Language[0].Name},
		{"Published", res.Items.Item[0].ItemAttributes.Languages.Language[0].Type},
		{"日本語", res.Items.Item[0].ItemAttributes.Languages.Language[1].Name},
		{"Unknown", res.Items.Item[0].ItemAttributes.Languages.Language[1].Type},
		{"4104", res.Items.Item[0].ItemAttributes.ListPrice.Amount},
		{"JPY", res.Items.Item[0].ItemAttributes.ListPrice.CurrencyCode},
		{"￥ 4,104", res.Items.Item[0].ItemAttributes.ListPrice.FormattedPrice},
		{"丸善出版", res.Items.Item[0].ItemAttributes.Manufacturer},
		{462, res.Items.Item[0].ItemAttributes.NumberOfPages},
		{"hundredths-inches", res.Items.Item[0].ItemAttributes.PackageDimensions.Width.Units},
		{732, res.Items.Item[0].ItemAttributes.PackageDimensions.Width.Value},
		{"hundredths-inches", res.Items.Item[0].ItemAttributes.PackageDimensions.Height.Units},
		{110, res.Items.Item[0].ItemAttributes.PackageDimensions.Height.Value},
		{"hundredths-pounds", res.Items.Item[0].ItemAttributes.PackageDimensions.Weight.Units},
		{171, res.Items.Item[0].ItemAttributes.PackageDimensions.Weight.Value},
		{"hundredths-inches", res.Items.Item[0].ItemAttributes.PackageDimensions.Length.Units},
		{909, res.Items.Item[0].ItemAttributes.PackageDimensions.Length.Value},
		{"Book", res.Items.Item[0].ItemAttributes.ProductGroup},
		{"ABIS_BOOK", res.Items.Item[0].ItemAttributes.ProductTypeName},
		{"丸善出版", res.Items.Item[0].ItemAttributes.Publisher},
		{"丸善出版", res.Items.Item[0].ItemAttributes.Studio},
		{"プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)", res.Items.Item[0].ItemAttributes.Title},
		{"4104", res.Items.Item[0].OfferSummary.LowestNewPrice.Amount},
		{"JPY", res.Items.Item[0].OfferSummary.LowestNewPrice.CurrencyCode},
		{"￥ 4,104", res.Items.Item[0].OfferSummary.LowestNewPrice.FormattedPrice},
		{"3799", res.Items.Item[0].OfferSummary.LowestUsedPrice.Amount},
		{"JPY", res.Items.Item[0].OfferSummary.LowestUsedPrice.CurrencyCode},
		{"￥ 3,799", res.Items.Item[0].OfferSummary.LowestUsedPrice.FormattedPrice},
		{1, res.Items.Item[0].OfferSummary.TotalNew},
		{4, res.Items.Item[0].OfferSummary.TotalUsed},
		{0, res.Items.Item[0].OfferSummary.TotalCollectible},
		{0, res.Items.Item[0].OfferSummary.TotalRefurbished},
		{1, res.Items.Item[0].Offers.TotalOfferPages},
		{1, res.Items.Item[0].Offers.TotalOffers},
		{"https://www.amazon.co.jp/gp/offer-listing/4621300253%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Datsushnagased-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D5143%26creativeASIN%3D4621300253", res.Items.Item[0].Offers.MoreOffersURL},
		{"Amazon.co.jp", res.Items.Item[0].Offers.Offer[0].Merchant.Name},
		{"New", res.Items.Item[0].Offers.Offer[0].OfferAttributes.Condition},
		{"NTPIbOCYgxigjLlkf1iTQhB6UfAcRHvlKju5nT%2BbVV876t1%2Bpt0pciArjHlsl9LS8iUJP9D5bajBzNN3VDdglcEAAS8lMPyCUArUG6CxF0A%3D", res.Items.Item[0].Offers.Offer[0].OfferListing.ID},
		{"4104", res.Items.Item[0].Offers.Offer[0].OfferListing.Price.Amount},
		{"JPY", res.Items.Item[0].Offers.Offer[0].OfferListing.Price.CurrencyCode},
		{"￥ 4,104", res.Items.Item[0].Offers.Offer[0].OfferListing.Price.FormattedPrice},
		{"在庫あり。", res.Items.Item[0].Offers.Offer[0].OfferListing.Availability},
		{"now", res.Items.Item[0].Offers.Offer[0].OfferListing.AvailabilityAttributes.AvailabilityType},
		{0, res.Items.Item[0].Offers.Offer[0].OfferListing.AvailabilityAttributes.MaximumHours},
		{0, res.Items.Item[0].Offers.Offer[0].OfferListing.AvailabilityAttributes.MinimumHours},
		{true, res.Items.Item[0].Offers.Offer[0].OfferListing.IsEligibleForPrime},
		{true, res.Items.Item[0].Offers.Offer[0].OfferListing.IsEligibleForSuperSaverShipping},
		{124, res.Items.Item[0].Offers.Offer[0].LoyaltyPoints.Points},
		{"124", res.Items.Item[0].Offers.Offer[0].LoyaltyPoints.TypicalRedemptionValue.Amount},
		{"JPY", res.Items.Item[0].Offers.Offer[0].LoyaltyPoints.TypicalRedemptionValue.CurrencyCode},
		{"￥ 124", res.Items.Item[0].Offers.Offer[0].LoyaltyPoints.TypicalRedemptionValue.FormattedPrice},
		{true, res.Items.Item[0].CustomerReviews.HasReviews},
		{"https://www.amazon.jp/reviews/iframe?akid=AKIAITPH62XKCOOT7AKA&alinkCode=xm2&asin=4621300253&atag=atsushnagased-22&exp=2016-11-18T19%3A16%3A15Z&v=2&sig=sbHAYHDNwLsf%2BAJMCbViRRSbKk%2Fq6%2FKplOPxiUVL2zw%3D", res.Items.Item[0].CustomerReviews.IFrameURL},
		{10, len(res.Items.Item[0].SimilarProducts.SimilarProduct)},
		{"477418392X", res.Items.Item[0].SimilarProducts.SimilarProduct[0].ASIN},
		{"みんなのGo言語【現場で使える実践テクニック】", res.Items.Item[0].SimilarProducts.SimilarProduct[0].Title},
		{"4873117526", res.Items.Item[0].SimilarProducts.SimilarProduct[1].ASIN},
		{"Go言語によるWebアプリケーション開発", res.Items.Item[0].SimilarProducts.SimilarProduct[1].Title},
		{"4798142417", res.Items.Item[0].SimilarProducts.SimilarProduct[2].ASIN},
		{"スターティングGo言語 (CodeZine BOOKS)", res.Items.Item[0].SimilarProducts.SimilarProduct[2].Title},
		{"4274219151", res.Items.Item[0].SimilarProducts.SimilarProduct[3].ASIN},
		{"プログラミングElixir", res.Items.Item[0].SimilarProducts.SimilarProduct[3].Title},
		{"4873117763", res.Items.Item[0].SimilarProducts.SimilarProduct[4].ASIN},
		{"Docker", res.Items.Item[0].SimilarProducts.SimilarProduct[4].Title},
		{"4798147400", res.Items.Item[0].SimilarProducts.SimilarProduct[5].ASIN},
		{"詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド (NEXT ONE)", res.Items.Item[0].SimilarProducts.SimilarProduct[5].Title},
		{"4774184322", res.Items.Item[0].SimilarProducts.SimilarProduct[6].ASIN},
		{"WEB+DB PRESS Vol.95", res.Items.Item[0].SimilarProducts.SimilarProduct[6].Title},
		{"4873117607", res.Items.Item[0].SimilarProducts.SimilarProduct[7].ASIN},
		{"マイクロサービスアーキテクチャ", res.Items.Item[0].SimilarProducts.SimilarProduct[7].Title},
		{"4798144576", res.Items.Item[0].SimilarProducts.SimilarProduct[8].ASIN},
		{"プログラマのためのSQLグラフ原論 リレーショナルデータベースで木と階層構造を扱うために", res.Items.Item[0].SimilarProducts.SimilarProduct[8].Title},
		{"4844381490", res.Items.Item[0].SimilarProducts.SimilarProduct[9].ASIN},
		{"Scalaスケーラブルプログラミング第3版", res.Items.Item[0].SimilarProducts.SimilarProduct[9].Title},
		{3, len(res.Items.Item[0].BrowseNodes.BrowseNode)},
		{"3229704051", res.Items.Item[0].BrowseNodes.BrowseNode[0].ID},
		{"ソフトウェア開発・言語", res.Items.Item[0].BrowseNodes.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Children.BrowseNode)},
		{"492352", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"プログラミング", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"466298", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"コンピュータ・IT", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"465610", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"ジャンル別", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"465392", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"本", res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"10805961", res.Items.Item[0].BrowseNodes.BrowseNode[1].ID},
		{"丸善", res.Items.Item[0].BrowseNodes.BrowseNode[1].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Children.BrowseNode)},
		{"465614", res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].ID},
		{"By Publishers", res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"465392", res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"本", res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[1].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"616893011", res.Items.Item[0].BrowseNodes.BrowseNode[2].ID},
		{"Custom Stores", res.Items.Item[0].BrowseNodes.BrowseNode[2].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode)},
		{6, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode)},
		{"3824511", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].ID},
		{"COOP", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"515742", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"Stores", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{1, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"465392", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		{"本", res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		{"13783651", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[0].ID},
		{"IDGストア", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[0].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[0].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[0].Children.BrowseNode)},

		{"13014911", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[1].ID},
		{"ソフトバンクの本", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[1].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[1].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[1].Children.BrowseNode)},

		{"1198480", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[2].ID},
		{"光文社ストア", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[2].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[2].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[2].Children.BrowseNode)},

		{"10924781", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[3].ID},
		{"宝島社ストア", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[3].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[3].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[3].Children.BrowseNode)},

		{"3370851", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[4].ID},
		{"日経BP社ストア", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[4].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[4].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[4].Children.BrowseNode)},

		{"3078071", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[5].ID},
		{"講談社ストア", res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[5].Name},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[5].Ancestors.BrowseNode)},
		{0, len(res.Items.Item[0].BrowseNodes.BrowseNode[2].Children.BrowseNode[5].Children.BrowseNode)},

		{time.Date(2016, 6, 20, 0, 0, 0, 0, time.UTC).UnixNano(), res.Items.Item[0].ItemAttributes.PublicationDate.UnixNano()},
		{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(), res.Items.Item[1].ItemAttributes.PublicationDate.UnixNano()},
		{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(), res.Items.Item[2].ItemAttributes.PublicationDate.UnixNano()},
	} {
		test.Compare(t)
	}
	// fmt.Printf("res %v\n", res)
}
