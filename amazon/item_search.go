package amazon

import (
	"encoding/xml"
	"net/http"
)

// ItemSearchResponseGroup represents constants those are capable ResponseGroups parameter
type ItemSearchResponseGroup string

const (
	// ItemSearchResponseGroupTags is a constant for Tags response group
	ItemSearchResponseGroupTags ItemSearchResponseGroup = "Tags"
	// ItemSearchResponseGroupHelp is a constant for Help response group
	ItemSearchResponseGroupHelp ItemSearchResponseGroup = "Help"
	// ItemSearchResponseGroupListMinimum is a constant for ListMinimum response group
	ItemSearchResponseGroupListMinimum ItemSearchResponseGroup = "ListMinimum"
	// ItemSearchResponseGroupTransactionDetails is a constant for TransactionDetails response group
	ItemSearchResponseGroupTransactionDetails ItemSearchResponseGroup = "TransactionDetails"
	// ItemSearchResponseGroupVariationImages is a constant for VariationImages response group
	ItemSearchResponseGroupVariationImages ItemSearchResponseGroup = "VariationImages"
	// ItemSearchResponseGroupPartBrandBinsSummary is a constant for PartBrandBinsSummary response group
	ItemSearchResponseGroupPartBrandBinsSummary ItemSearchResponseGroup = "PartBrandBinsSummary"
	// ItemSearchResponseGroupCustomerFull is a constant for CustomerFull response group
	ItemSearchResponseGroupCustomerFull ItemSearchResponseGroup = "CustomerFull"
	// ItemSearchResponseGroupCartNewReleases is a constant for CartNewReleases response group
	ItemSearchResponseGroupCartNewReleases ItemSearchResponseGroup = "CartNewReleases"
	// ItemSearchResponseGroupItemIds is a constant for ItemIds response group
	ItemSearchResponseGroupItemIds ItemSearchResponseGroup = "ItemIds"
	// ItemSearchResponseGroupTagsSummary is a constant for TagsSummary response group
	ItemSearchResponseGroupTagsSummary ItemSearchResponseGroup = "TagsSummary"

	// ItemSearchResponseGroupFitments is a constant for Fitments response group
	ItemSearchResponseGroupFitments ItemSearchResponseGroup = "Fitments"
	// ItemSearchResponseGroupPartBrowseNodeBinsSummary is a constant for PartBrowseNodeBinsSummary response group
	ItemSearchResponseGroupPartBrowseNodeBinsSummary ItemSearchResponseGroup = "PartBrowseNodeBinsSummary"
	// ItemSearchResponseGroupTopSellers is a constant for TopSellers response group
	ItemSearchResponseGroupTopSellers ItemSearchResponseGroup = "TopSellers"
	// ItemSearchResponseGroupRequest is a constant for Request response group
	ItemSearchResponseGroupRequest ItemSearchResponseGroup = "Request"
	// ItemSearchResponseGroupHasPartCompatibility is a constant for HasPartCompatibility response group
	ItemSearchResponseGroupHasPartCompatibility ItemSearchResponseGroup = "HasPartCompatibility"
	// ItemSearchResponseGroupListFull is a constant for ListFull response group
	ItemSearchResponseGroupListFull ItemSearchResponseGroup = "ListFull"
	// ItemSearchResponseGroupSeller is a constant for Seller response group
	ItemSearchResponseGroupSeller ItemSearchResponseGroup = "Seller"
	// ItemSearchResponseGroupVehicleMakes is a constant for VehicleMakes response group
	ItemSearchResponseGroupVehicleMakes ItemSearchResponseGroup = "VehicleMakes"
	// ItemSearchResponseGroupTaggedItems is a constant for TaggedItems response group
	ItemSearchResponseGroupTaggedItems ItemSearchResponseGroup = "TaggedItems"
	// ItemSearchResponseGroupVehicleParts is a constant for VehicleParts response group
	ItemSearchResponseGroupVehicleParts ItemSearchResponseGroup = "VehicleParts"
	// ItemSearchResponseGroupBrowseNodeInfo is a constant for BrowseNodeInfo response group
	ItemSearchResponseGroupBrowseNodeInfo ItemSearchResponseGroup = "BrowseNodeInfo"
	// ItemSearchResponseGroupItemAttributes is a constant for ItemAttributes response group
	ItemSearchResponseGroupItemAttributes ItemSearchResponseGroup = "ItemAttributes"
	// ItemSearchResponseGroupVehicleOptions is a constant for VehicleOptions response group
	ItemSearchResponseGroupVehicleOptions ItemSearchResponseGroup = "VehicleOptions"
	// ItemSearchResponseGroupListItems is a constant for ListItems response group
	ItemSearchResponseGroupListItems ItemSearchResponseGroup = "ListItems"
	// ItemSearchResponseGroupTaggedGuides is a constant for TaggedGuides response group
	ItemSearchResponseGroupTaggedGuides ItemSearchResponseGroup = "TaggedGuides"
	// ItemSearchResponseGroupNewReleases is a constant for NewReleases response group
	ItemSearchResponseGroupNewReleases ItemSearchResponseGroup = "NewReleases"
	// ItemSearchResponseGroupVehiclePartFit is a constant for VehiclePartFit response group
	ItemSearchResponseGroupVehiclePartFit ItemSearchResponseGroup = "VehiclePartFit"
	// ItemSearchResponseGroupCartSimilarities is a constant for CartSimilarities response group
	ItemSearchResponseGroupCartSimilarities ItemSearchResponseGroup = "CartSimilarities"
	// ItemSearchResponseGroupShippingCharges is a constant for ShippingCharges response group
	ItemSearchResponseGroupShippingCharges ItemSearchResponseGroup = "ShippingCharges"
	// ItemSearchResponseGroupShippingOptions is a constant for ShippingOptions response group
	ItemSearchResponseGroupShippingOptions ItemSearchResponseGroup = "ShippingOptions"
	// ItemSearchResponseGroupCustomerInfo is a constant for CustomerInfo response group
	ItemSearchResponseGroupCustomerInfo ItemSearchResponseGroup = "CustomerInfo"
	// ItemSearchResponseGroupPartnerTransactionDetails is a constant for PartnerTransactionDetails response group
	ItemSearchResponseGroupPartnerTransactionDetails ItemSearchResponseGroup = "PartnerTransactionDetails"
	// ItemSearchResponseGroupVehicleYears is a constant for VehicleYears response group
	ItemSearchResponseGroupVehicleYears ItemSearchResponseGroup = "VehicleYears"
	// ItemSearchResponseGroupVehicleTrims is a constant for VehicleTrims response group
	ItemSearchResponseGroupVehicleTrims ItemSearchResponseGroup = "VehicleTrims"
	// ItemSearchResponseGroupCustomerReviews is a constant for CustomerReviews response group
	ItemSearchResponseGroupCustomerReviews ItemSearchResponseGroup = "CustomerReviews"
	// ItemSearchResponseGroupSellerListing is a constant for SellerListing response group
	ItemSearchResponseGroupSellerListing ItemSearchResponseGroup = "SellerListing"
	// ItemSearchResponseGroupCart is a constant for Cart response group
	ItemSearchResponseGroupCart ItemSearchResponseGroup = "Cart"
	// ItemSearchResponseGroupTaggedListmaniaLists is a constant for TaggedListmaniaLists response group
	ItemSearchResponseGroupTaggedListmaniaLists ItemSearchResponseGroup = "TaggedListmaniaLists"
	// ItemSearchResponseGroupVehicleModels is a constant for VehicleModels response group
	ItemSearchResponseGroupVehicleModels ItemSearchResponseGroup = "VehicleModels"
	// ItemSearchResponseGroupListInfo is a constant for ListInfo response group
	ItemSearchResponseGroupListInfo ItemSearchResponseGroup = "ListInfo"
	// ItemSearchResponseGroupCustomerLists is a constant for CustomerLists response group
	ItemSearchResponseGroupCustomerLists ItemSearchResponseGroup = "CustomerLists"
	// ItemSearchResponseGroupCartTopSellers is a constant for CartTopSellers response group
	ItemSearchResponseGroupCartTopSellers ItemSearchResponseGroup = "CartTopSellers"
	// ItemSearchResponseGroupCollections is a constant for Collections response group
	ItemSearchResponseGroupCollections ItemSearchResponseGroup = "Collections"

	// ItemSearchResponseGroupAccessories is a constant for Accessories response group
	ItemSearchResponseGroupAccessories ItemSearchResponseGroup = "Accessories"
	// ItemSearchResponseGroupAlternateVersions is a constant for AlternateVersions response group
	ItemSearchResponseGroupAlternateVersions ItemSearchResponseGroup = "AlternateVersions"
	// ItemSearchResponseGroupBrowseNodes is a constant for BrowseNodes response group
	ItemSearchResponseGroupBrowseNodes ItemSearchResponseGroup = "BrowseNodes"
	// ItemSearchResponseGroupEditorialReview is a constant for EditorialReview response group
	ItemSearchResponseGroupEditorialReview ItemSearchResponseGroup = "EditorialReview"
	// ItemSearchResponseGroupImages is a constant for Images response group
	ItemSearchResponseGroupImages ItemSearchResponseGroup = "Images"
	// ItemSearchResponseGroupItemAttributesItemIds is a constant for ItemAttributesItemIds response group
	ItemSearchResponseGroupItemAttributesItemIds ItemSearchResponseGroup = "ItemAttributesItemIds"
	// ItemSearchResponseGroupLarge is a constant for Large response group
	ItemSearchResponseGroupLarge ItemSearchResponseGroup = "Large"
	// ItemSearchResponseGroupListmaniaLists is a constant for ListmaniaLists response group
	ItemSearchResponseGroupListmaniaLists ItemSearchResponseGroup = "ListmaniaLists"
	// ItemSearchResponseGroupMedium is a constant for Medium response group
	ItemSearchResponseGroupMedium ItemSearchResponseGroup = "Medium"
	// ItemSearchResponseGroupMerchantItemAttributes is a constant for MerchantItemAttributes response group
	ItemSearchResponseGroupMerchantItemAttributes ItemSearchResponseGroup = "MerchantItemAttributes"
	// ItemSearchResponseGroupOfferFull is a constant for OfferFull response group
	ItemSearchResponseGroupOfferFull ItemSearchResponseGroup = "OfferFull"
	// ItemSearchResponseGroupOfferListings is a constant for OfferListings response group
	ItemSearchResponseGroupOfferListings ItemSearchResponseGroup = "OfferListings"
	// ItemSearchResponseGroupOffers is a constant for Offers response group
	ItemSearchResponseGroupOffers ItemSearchResponseGroup = "Offers"
	// ItemSearchResponseGroupOfferSummary is a constant for OfferSummary response group
	ItemSearchResponseGroupOfferSummary ItemSearchResponseGroup = "OfferSummary"
	// ItemSearchResponseGroupPromotionalTag is a constant for PromotionalTag response group
	ItemSearchResponseGroupPromotionalTag ItemSearchResponseGroup = "PromotionalTag"
	// ItemSearchResponseGroupPromotionDetails is a constant for PromotionDetails response group
	ItemSearchResponseGroupPromotionDetails ItemSearchResponseGroup = "PromotionDetails"
	// ItemSearchResponseGroupPromotionSummary is a constant for PromotionSummary response group
	ItemSearchResponseGroupPromotionSummary ItemSearchResponseGroup = "PromotionSummary"
	// ItemSearchResponseGroupRelatedItems is a constant for RelatedItems response group
	ItemSearchResponseGroupRelatedItems ItemSearchResponseGroup = "RelatedItems"
	// ItemSearchResponseGroupReviews is a constant for Reviews response group
	ItemSearchResponseGroupReviews ItemSearchResponseGroup = "Reviews"
	// ItemSearchResponseGroupSalesRank is a constant for SalesRank response group
	ItemSearchResponseGroupSalesRank ItemSearchResponseGroup = "SalesRank"
	// ItemSearchResponseGroupSearchBins is a constant for SearchBins response group
	ItemSearchResponseGroupSearchBins ItemSearchResponseGroup = "SearchBins"
	// ItemSearchResponseGroupSearchInside is a constant for SearchInside response group
	ItemSearchResponseGroupSearchInside ItemSearchResponseGroup = "SearchInside"
	// ItemSearchResponseGroupSimilarities is a constant for Similarities response group
	ItemSearchResponseGroupSimilarities ItemSearchResponseGroup = "Similarities"
	// ItemSearchResponseGroupSmall is a constant for Small response group
	ItemSearchResponseGroupSmall ItemSearchResponseGroup = "Small"
	// ItemSearchResponseGroupSubjects is a constant for Subjects response group
	ItemSearchResponseGroupSubjects ItemSearchResponseGroup = "Subjects"
	// ItemSearchResponseGroupTracks is a constant for Tracks response group
	ItemSearchResponseGroupTracks ItemSearchResponseGroup = "Tracks"
	// ItemSearchResponseGroupVariationMatrix is a constant for VariationMatrix response group
	ItemSearchResponseGroupVariationMatrix ItemSearchResponseGroup = "VariationMatrix"
	// ItemSearchResponseGroupVariationMinimum is a constant for VariationMinimum response group
	ItemSearchResponseGroupVariationMinimum ItemSearchResponseGroup = "VariationMinimum"
	// ItemSearchResponseGroupVariationOffers is a constant for VariationOffers response group
	ItemSearchResponseGroupVariationOffers ItemSearchResponseGroup = "VariationOffers"
	// ItemSearchResponseGroupVariations is a constant for Variations response group
	ItemSearchResponseGroupVariations ItemSearchResponseGroup = "Variations"
	// ItemSearchResponseGroupVariationSummary is a constant for VariationSummary response group
	ItemSearchResponseGroupVariationSummary ItemSearchResponseGroup = "VariationSummary"
)

// ItemSearchParameters represents parameters for ItemSearch operation request
type ItemSearchParameters struct {
	// Actor name associated with the item. You can enter all or part of the name.
	Actor string
	// Artist name associated with the item. You can enter all or part of the name.
	Artist string
	// Movie ratings based on MPAA ratings or age, depending on locale.
	// You can specify one or more values in a comma-separated list
	AudienceRating string
	// Author name associated with the item. You can enter all or part of the name.
	Author string
	// Returns available items only
	OnlyAvailable bool
	// Brand name associated with the item. You can enter all or part of the name.
	Brand string
	// Browse nodes are numbers that identify product categories. For example, the browse node for Literature & Fiction is 17, while the browse node for Outdoors & Nature is 290060.
	BrowseNode string
	// Composer name associated with the item. You can enter all or part of the name.
	Composer string
	// Condition filters offers by condition type. By default, Condition equals New. When the Availability parameter is set to Available, the Condition parameter cannot be set to New.
	Condition
	// Conductor Conductor name associated with the item. You can enter all or part of the name.
	Conductor string
	// Director Director name associated with the item. You can enter all or part of the name.
	Director string
	// IncludeReviewsSummary Returns the reviews summary URL.
	IncludeReviewsSummary *bool
	// ItemPage returns a specific page of items from the available search results. Up to ten items are returned per page.
	// If you do not include ItemPage in your request, the first page is returned. The total number of pages found is returned in the TotalPages response element.
	// If Condition is set to All, ItemSearch returns additional offers for those items, one offer per condition type.
	// Valid values: 1 to 10 (1 to 5 when search index is All)
	ItemPage int
	// Keywords A word or phrase that describes an item, including author, artist, description, manufacturer, title, and so on.
	// For example, when SearchIndex is set to MusicTracks, the Keywords parameter can search for song title.
	Keywords string
	// Manufacturer Manufacturer name associated with the item. You can enter all or part of the name.
	Manufacturer string
	// MaxPrice Specifies the maximum item price in the response. Prices appear in the lowest currency denomination. For example, 3241 is $32.41. MaximumPrice can be used with every index, except All and Blended.
	MaximumPrice int
	// MerchantId Filters search results and offer listings to items sold by Amazon. By default, the Product Advertising API returns items sold by merchants and Amazon.
	MerchantID string
	// MinimumPrice Specifies the minimum item price in the response. Prices appear in the lowest currency denomination. For example, 3241 is $32.41. MinimumPrice can be used with every index, except All and Blended.
	MinimumPrice int
	// MinPercentageOff Specifies the minimum percentage off the item price.
	MinPercentageOff int
	// Orchestra Orchestra name associated with the item. You can enter all or part of the name.
	Orchestra string
	// Power Performs a book search with a complex query string. The parameter can be used only when SearchIndex is set to Books.
	// See http://docs.aws.amazon.com/AWSECommerceService/latest/DG/PowerSearchSyntax.html
	Power string
	// Publisher Publisher name associated with the item. You can enter all or part of the name.
	Publisher string
	// RelatedItemPage Returns a specific page of related items from the available search results. Up to ten items are returned per page. This parameter can be used with the RelatedItems response group.
	RelatedItemPage int
	// RelationshipType This parameter is required when the RelatedItems response group is used. The type of related item returned is specified by the RelationshipType parameter. Sample values include Episode, Season, and Tracks.
	RelationshipType
	// SearchIndex The product category to search.
	SearchIndex
	// Sort The way in which items in the response are ordered.
	Sort string
	// Title Title associated with the item. You can enter all or part of the title. Title searches are a subset of Keyword searches. Use a Keywords search if a Title search does not return the items you want.
	Title string
	// TruncateReviewsAt By default, reviews are truncated to 1000 characters. Choose a value to specify a length. To return the entire review, use 0 .
	TruncateReviewsAt *int
	// VariationPage Returns a specific page of variations. For example, set VariationPage to 2 to return offers 11 to 20 . The total number of pages appears in the TotalPages element.
	VariationPage *int
	// Specifies the types of values to return. Separate
	ResponseGroups []ItemSearchResponseGroup
}

// ItemSearchRequest represents request for ItemSearch operation
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemSearch.html
type ItemSearchRequest struct {
	Client     *Client
	Parameters ItemSearchParameters
}

// ItemSearchResponse represents response for ItemSearch operation
type ItemSearchResponse struct {
	XMLName xml.Name `xml:"ItemSearchResponse"`
	Items   Items    `xml:"Items"`
}

// Error returns Error found
func (res *ItemSearchResponse) Error() error {
	if e := res.Items.Request.Errors; e != nil {
		return e
	}
	return nil
}

// Query returns query for sending request
func (req *ItemSearchRequest) Query() map[string]interface{} {
	q := map[string]interface{}{}
	p := req.Parameters
	for k, strp := range map[string]string{
		"Actor":            p.Actor,
		"Artist":           p.Artist,
		"AudienceRating":   p.AudienceRating,
		"Author":           p.Author,
		"Brand":            p.Brand,
		"BrowseNode":       p.BrowseNode,
		"Composer":         p.Composer,
		"Condition":        string(p.Condition),
		"Conductor":        p.Conductor,
		"Director":         p.Director,
		"Keywords":         p.Keywords,
		"Manufacturer":     p.Manufacturer,
		"MerchantID":       p.MerchantID,
		"Orchestra":        p.Orchestra,
		"Power":            p.Power,
		"Publisher":        p.Publisher,
		"SearchIndex":      string(p.SearchIndex),
		"Sort":             p.Sort,
		"Title":            p.Title,
		"RelationshipType": string(p.RelationshipType),
	} {
		if strp != "" {
			q[k] = strp
		}
	}
	for k, intp := range map[string]int{
		"ItemPage":         p.ItemPage,
		"MaximumPrice":     p.MaximumPrice,
		"MinimumPrice":     p.MinimumPrice,
		"MinPercentageOff": p.MinPercentageOff,
		"RelatedItemPage":  p.RelatedItemPage,
	} {
		if intp > 0 {
			q[k] = intp
		}
	}
	if p.OnlyAvailable {
		q["Availability"] = "Available"
	}
	if p.IncludeReviewsSummary != nil {
		q["IncludeReviewsSummary"] = *p.IncludeReviewsSummary
	}
	if p.TruncateReviewsAt != nil {
		q["TruncateReviewsAt"] = *p.TruncateReviewsAt
	}
	if p.VariationPage != nil {
		q["VariationPage"] = *p.VariationPage
	}
	q["ResponseGroup"] = p.ResponseGroups
	return q
}

func (req *ItemSearchRequest) httpMethod() string {
	return http.MethodGet
}

func (req *ItemSearchRequest) operation() string {
	return "ItemSearch"
}

// Do sends request for the API
func (req *ItemSearchRequest) Do() (*ItemSearchResponse, error) {
	respObj := ItemSearchResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// ItemSearch returns new request for ItemSearch
func (client *Client) ItemSearch(parameters ItemSearchParameters) *ItemSearchRequest {
	return &ItemSearchRequest{
		Client:     client,
		Parameters: parameters,
	}
}
