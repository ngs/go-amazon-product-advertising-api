package amazon

// SimilarityLookupResponseGroup represents constants those are capable ResponseGroups parameter
type SimilarityLookupResponseGroup string

const (
	// SimilarityLookupResponseGroupBrowseNodes is a constant for BrowseNodes response group
	SimilarityLookupResponseGroupBrowseNodes SimilarityLookupResponseGroup = "BrowseNodes"
	// SimilarityLookupResponseGroupPromotionDetails is a constant for PromotionDetails response group
	SimilarityLookupResponseGroupPromotionDetails SimilarityLookupResponseGroup = "PromotionDetails"
	// SimilarityLookupResponseGroupEditorialReview is a constant for EditorialReview response group
	SimilarityLookupResponseGroupEditorialReview SimilarityLookupResponseGroup = "EditorialReview"
	// SimilarityLookupResponseGroupPromotionSummary is a constant for PromotionSummary response group
	SimilarityLookupResponseGroupPromotionSummary SimilarityLookupResponseGroup = "PromotionSummary"
	// SimilarityLookupResponseGroupImages is a constant for Images response group
	SimilarityLookupResponseGroupImages SimilarityLookupResponseGroup = "Images"
	// SimilarityLookupResponseGroupReviews is a constant for Reviews response group
	SimilarityLookupResponseGroupReviews SimilarityLookupResponseGroup = "Reviews"
	// SimilarityLookupResponseGroupItemAttributes is a constant for ItemAttributes response group
	SimilarityLookupResponseGroupItemAttributes SimilarityLookupResponseGroup = "ItemAttributes"
	// SimilarityLookupResponseGroupSalesRankItemIds is a constant for SalesRankItemIds response group
	SimilarityLookupResponseGroupSalesRankItemIds SimilarityLookupResponseGroup = "SalesRankItemIds"
	// SimilarityLookupResponseGroupSimilarities is a constant for Similarities response group
	SimilarityLookupResponseGroupSimilarities SimilarityLookupResponseGroup = "Similarities"
	// SimilarityLookupResponseGroupLarge is a constant for Large response group
	SimilarityLookupResponseGroupLarge SimilarityLookupResponseGroup = "Large"
	// SimilarityLookupResponseGroupSmall is a constant for Small response group
	SimilarityLookupResponseGroupSmall SimilarityLookupResponseGroup = "Small"
	// SimilarityLookupResponseGroupListmaniaLists is a constant for ListmaniaLists response group
	SimilarityLookupResponseGroupListmaniaLists SimilarityLookupResponseGroup = "ListmaniaLists"
	// SimilarityLookupResponseGroupSubjects is a constant for Subjects response group
	SimilarityLookupResponseGroupSubjects SimilarityLookupResponseGroup = "Subjects"
	// SimilarityLookupResponseGroupMedium is a constant for Medium response group
	SimilarityLookupResponseGroupMedium SimilarityLookupResponseGroup = "Medium"
	// SimilarityLookupResponseGroupTracks is a constant for Tracks response group
	SimilarityLookupResponseGroupTracks SimilarityLookupResponseGroup = "Tracks"
	// SimilarityLookupResponseGroupOfferFull is a constant for OfferFull response group
	SimilarityLookupResponseGroupOfferFull SimilarityLookupResponseGroup = "OfferFull"
	// SimilarityLookupResponseGroupVariationMinimum is a constant for VariationMinimum response group
	SimilarityLookupResponseGroupVariationMinimum SimilarityLookupResponseGroup = "VariationMinimum"
	// SimilarityLookupResponseGroupOfferListings is a constant for OfferListings response group
	SimilarityLookupResponseGroupOfferListings SimilarityLookupResponseGroup = "OfferListings"
	// SimilarityLookupResponseGroupVariations is a constant for Variations response group
	SimilarityLookupResponseGroupVariations SimilarityLookupResponseGroup = "Variations"
	// SimilarityLookupResponseGroupOffers is a constant for Offers response group
	SimilarityLookupResponseGroupOffers SimilarityLookupResponseGroup = "Offers"
	// SimilarityLookupResponseGroupVariationSummary is a constant for VariationSummary response group
	SimilarityLookupResponseGroupVariationSummary SimilarityLookupResponseGroup = "VariationSummary"
)

// SimilarityLookupRequest represents request for SimilarityLookup operation
type SimilarityLookupRequest struct {
	ResponseGroups []SimilarityLookupResponseGroup
	Client         *Client
}

// SimilarityLookupResponse represents response for SimilarityLookup operation
type SimilarityLookupResponse struct {
	Error error
}

// Do send request for the API
func (req *SimilarityLookupRequest) Do() (*SimilarityLookupResponse, error) {
	return nil, nil
}

// SimilarityLookup returns new request for SimilarityLookup
func (client *Client) SimilarityLookup(responseGroups ...SimilarityLookupResponseGroup) *SimilarityLookupRequest {
	return &SimilarityLookupRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
