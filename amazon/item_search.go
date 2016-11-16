package amazon

// ItemSearchResponseGroup represents constants those are capable ResponseGroups parameter
type ItemSearchResponseGroup string

const (
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

// ItemSearchRequest represents request for ItemSearch operation
type ItemSearchRequest struct {
	ResponseGroups []ItemSearchResponseGroup
	Client         *Client
}

// ItemSearchResponse represents response for ItemSearch operation
type ItemSearchResponse struct {
	Error error
}

func (req *ItemSearchRequest) buildQuery() map[string]string {
	q := map[string]string{}
	return q
}

// Do send request for the API
func (req *ItemSearchRequest) Do() (*ItemSearchResponse, error) {
	return nil, nil
}

// ItemSearch returns new request for ItemSearch
func (client *Client) ItemSearch(responseGroups ...ItemSearchResponseGroup) *ItemSearchRequest {
	return &ItemSearchRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
