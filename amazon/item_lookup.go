package amazon

import "net/http"

// ItemLookupResponseGroup represents constants those are capable ResponseGroups parameter
type ItemLookupResponseGroup string

const (
	// ItemLookupResponseGroupAccessories is a constant for Accessories response group
	ItemLookupResponseGroupAccessories ItemLookupResponseGroup = "Accessories"
	// ItemLookupResponseGroupAlternateVersions is a constant for AlternateVersions response group
	ItemLookupResponseGroupAlternateVersions ItemLookupResponseGroup = "AlternateVersions"
	// ItemLookupResponseGroupBrowseNodes is a constant for BrowseNodes response group
	ItemLookupResponseGroupBrowseNodes ItemLookupResponseGroup = "BrowseNodes"
	// ItemLookupResponseGroupEditorialReview is a constant for EditorialReview response group
	ItemLookupResponseGroupEditorialReview ItemLookupResponseGroup = "EditorialReview"
	// ItemLookupResponseGroupImages is a constant for Images response group
	ItemLookupResponseGroupImages ItemLookupResponseGroup = "Images"
	// ItemLookupResponseGroupItemAttributesItemIds is a constant for ItemAttributesItemIds response group
	ItemLookupResponseGroupItemAttributesItemIds ItemLookupResponseGroup = "ItemAttributesItemIds"
	// ItemLookupResponseGroupLarge is a constant for Large response group
	ItemLookupResponseGroupLarge ItemLookupResponseGroup = "Large"
	// ItemLookupResponseGroupListmaniaLists is a constant for ListmaniaLists response group
	ItemLookupResponseGroupListmaniaLists ItemLookupResponseGroup = "ListmaniaLists"
	// ItemLookupResponseGroupMedium is a constant for Medium response group
	ItemLookupResponseGroupMedium ItemLookupResponseGroup = "Medium"
	// ItemLookupResponseGroupMerchantItemAttributes is a constant for MerchantItemAttributes response group
	ItemLookupResponseGroupMerchantItemAttributes ItemLookupResponseGroup = "MerchantItemAttributes"
	// ItemLookupResponseGroupOfferFull is a constant for OfferFull response group
	ItemLookupResponseGroupOfferFull ItemLookupResponseGroup = "OfferFull"
	// ItemLookupResponseGroupOfferListings is a constant for OfferListings response group
	ItemLookupResponseGroupOfferListings ItemLookupResponseGroup = "OfferListings"
	// ItemLookupResponseGroupOffers is a constant for Offers response group
	ItemLookupResponseGroupOffers ItemLookupResponseGroup = "Offers"
	// ItemLookupResponseGroupOfferSummary is a constant for OfferSummary response group
	ItemLookupResponseGroupOfferSummary ItemLookupResponseGroup = "OfferSummary"
	// ItemLookupResponseGroupPromotionalTag is a constant for PromotionalTag response group
	ItemLookupResponseGroupPromotionalTag ItemLookupResponseGroup = "PromotionalTag"
	// ItemLookupResponseGroupPromotionDetails is a constant for PromotionDetails response group
	ItemLookupResponseGroupPromotionDetails ItemLookupResponseGroup = "PromotionDetails"
	// ItemLookupResponseGroupPromotionSummary is a constant for PromotionSummary response group
	ItemLookupResponseGroupPromotionSummary ItemLookupResponseGroup = "PromotionSummary"
	// ItemLookupResponseGroupRelatedItems is a constant for RelatedItems response group
	ItemLookupResponseGroupRelatedItems ItemLookupResponseGroup = "RelatedItems"
	// ItemLookupResponseGroupReviews is a constant for Reviews response group
	ItemLookupResponseGroupReviews ItemLookupResponseGroup = "Reviews"
	// ItemLookupResponseGroupSalesRank is a constant for SalesRank response group
	ItemLookupResponseGroupSalesRank ItemLookupResponseGroup = "SalesRank"
	// ItemLookupResponseGroupSearchInside is a constant for SearchInside response group
	ItemLookupResponseGroupSearchInside ItemLookupResponseGroup = "SearchInside"
	// ItemLookupResponseGroupShippingCharges is a constant for ShippingCharges response group
	ItemLookupResponseGroupShippingCharges ItemLookupResponseGroup = "ShippingCharges"
	// ItemLookupResponseGroupSimilarities is a constant for Similarities response group
	ItemLookupResponseGroupSimilarities ItemLookupResponseGroup = "Similarities"
	// ItemLookupResponseGroupSmall is a constant for Small response group
	ItemLookupResponseGroupSmall ItemLookupResponseGroup = "Small"
	// ItemLookupResponseGroupSubjects is a constant for Subjects response group
	ItemLookupResponseGroupSubjects ItemLookupResponseGroup = "Subjects"
	// ItemLookupResponseGroupTracks is a constant for Tracks response group
	ItemLookupResponseGroupTracks ItemLookupResponseGroup = "Tracks"
	// ItemLookupResponseGroupVariationImages is a constant for VariationImages response group
	ItemLookupResponseGroupVariationImages ItemLookupResponseGroup = "VariationImages"
	// ItemLookupResponseGroupVariationMatrix is a constant for VariationMatrix response group
	ItemLookupResponseGroupVariationMatrix ItemLookupResponseGroup = "VariationMatrix"
	// ItemLookupResponseGroupVariationMinimum is a constant for VariationMinimum response group
	ItemLookupResponseGroupVariationMinimum ItemLookupResponseGroup = "VariationMinimum"
	// ItemLookupResponseGroupVariationOffers is a constant for VariationOffers response group
	ItemLookupResponseGroupVariationOffers ItemLookupResponseGroup = "VariationOffers"
	// ItemLookupResponseGroupVariations is a constant for Variations response group
	ItemLookupResponseGroupVariations ItemLookupResponseGroup = "Variations"
	// ItemLookupResponseGroupVariationSummary is a constant for VariationSummary response group
	ItemLookupResponseGroupVariationSummary ItemLookupResponseGroup = "VariationSummary"
)

// ItemLookupRequest represents request for ItemLookup operation
type ItemLookupRequest struct {
	ResponseGroups []ItemLookupResponseGroup
	Client         *Client
}

// ItemLookupResponse represents response for ItemLookup operation
type ItemLookupResponse struct {
	Error error
}

func (req *ItemLookupRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *ItemLookupRequest) httpMethod() string {
	return http.MethodGet
}

func (req *ItemLookupRequest) operation() string {
	return "ItemLookup"
}

// Do sends request for the API
func (req *ItemLookupRequest) Do() (*ItemLookupResponse, error) {
	_, err := req.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ItemLookup returns new request for ItemLookup
func (client *Client) ItemLookup(responseGroups ...ItemLookupResponseGroup) *ItemLookupRequest {
	return &ItemLookupRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
