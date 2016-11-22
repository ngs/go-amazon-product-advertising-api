package amazon

import (
	"encoding/xml"
	"net/http"
	"strings"
)

// ItemLookupResponseGroup represents constants those are capable ResponseGroups parameter
type ItemLookupResponseGroup string

const (
	// ItemLookupResponseGroupAccessories is a constant for Accessories response group
	ItemLookupResponseGroupAccessories ItemLookupResponseGroup = "Accessories"
	// ItemLookupResponseGroupBrowseNodes is a constant for BrowseNodes response group
	ItemLookupResponseGroupBrowseNodes ItemLookupResponseGroup = "BrowseNodes"
	// ItemLookupResponseGroupEditorialReview is a constant for EditorialReview response group
	ItemLookupResponseGroupEditorialReview ItemLookupResponseGroup = "EditorialReview"
	// ItemLookupResponseGroupImages is a constant for Images response group
	ItemLookupResponseGroupImages ItemLookupResponseGroup = "Images"
	// ItemLookupResponseGroupItemAttributes is a constant for ItemAttributes response group
	ItemLookupResponseGroupItemAttributes ItemLookupResponseGroup = "ItemAttributes"
	// ItemLookupResponseGroupItemIds is a constant for ItemIds response group
	ItemLookupResponseGroupItemIds ItemLookupResponseGroup = "ItemIds"
	// ItemLookupResponseGroupLarge is a constant for Large response group
	ItemLookupResponseGroupLarge ItemLookupResponseGroup = "Large"
	// ItemLookupResponseGroupMedium is a constant for Medium response group
	ItemLookupResponseGroupMedium ItemLookupResponseGroup = "Medium"
	// ItemLookupResponseGroupOfferFull is a constant for OfferFull response group
	ItemLookupResponseGroupOfferFull ItemLookupResponseGroup = "OfferFull"
	// ItemLookupResponseGroupOffers is a constant for Offers response group
	ItemLookupResponseGroupOffers ItemLookupResponseGroup = "Offers"
	// ItemLookupResponseGroupPromotionSummary is a constant for PromotionSummary response group
	ItemLookupResponseGroupPromotionSummary ItemLookupResponseGroup = "PromotionSummary"
	// ItemLookupResponseGroupOfferSummary is a constant for OfferSummary response group
	ItemLookupResponseGroupOfferSummary ItemLookupResponseGroup = "OfferSummary"
	// ItemLookupResponseGroupRelatedItems is a constant for RelatedItems response group
	ItemLookupResponseGroupRelatedItems ItemLookupResponseGroup = "RelatedItems"
	// ItemLookupResponseGroupReviews is a constant for Reviews response group
	ItemLookupResponseGroupReviews ItemLookupResponseGroup = "Reviews"
	// ItemLookupResponseGroupSalesRank is a constant for SalesRank response group
	ItemLookupResponseGroupSalesRank ItemLookupResponseGroup = "SalesRank"
	// ItemLookupResponseGroupSimilarities is a constant for Similarities response group
	ItemLookupResponseGroupSimilarities ItemLookupResponseGroup = "Similarities"
	// ItemLookupResponseGroupSmall is a constant for Small response group
	ItemLookupResponseGroupSmall ItemLookupResponseGroup = "Small"
	// ItemLookupResponseGroupTracks is a constant for Tracks response group
	ItemLookupResponseGroupTracks ItemLookupResponseGroup = "Tracks"
	// ItemLookupResponseGroupVariationImages is a constant for VariationImages response group
	ItemLookupResponseGroupVariationImages ItemLookupResponseGroup = "VariationImages"
	// ItemLookupResponseGroupVariations is a constant for Variations response group
	ItemLookupResponseGroupVariations ItemLookupResponseGroup = "Variations"
	// ItemLookupResponseGroupVariationSummary is a constant for VariationSummary response group
	ItemLookupResponseGroupVariationSummary ItemLookupResponseGroup = "VariationSummary"
)

// ItemLookupParameters represents parameters for ItemLookup operation request
type ItemLookupParameters struct {
	ResponseGroups        []ItemLookupResponseGroup
	Condition             Condition
	IDType                IDType
	ItemIDs               []string
	IncludeReviewsSummary *bool
	MerchantID            string
	RelatedItemPage       int
	RelationshipType      RelationshipType
	SearchIndex           SearchIndex
	TruncateReviewsAt     *int
	VariationPage         int
}

// ItemLookupRequest represents request for ItemLookup operation
type ItemLookupRequest struct {
	Client     *Client
	Parameters ItemLookupParameters
}

// ItemLookupResponse represents response for ItemLookup operation
type ItemLookupResponse struct {
	XMLName xml.Name `xml:"ItemLookupResponse"`
	Items   Items    `xml:"Items"`
}

// Error returns Error found
func (res *ItemLookupResponse) Error() error {
	if e := res.Items.Request.Errors; e != nil {
		return e
	}
	return nil
}

func (req *ItemLookupRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	if string(req.Parameters.Condition) != "" {
		q["Condition"] = string(req.Parameters.Condition)
	}
	q["IdType"] = string(req.Parameters.IDType)
	q["ItemId"] = strings.Join(req.Parameters.ItemIDs, ",")
	if req.Parameters.IncludeReviewsSummary != nil {
		q["IncludeReviewsSummary"] = *req.Parameters.IncludeReviewsSummary
	}
	if req.Parameters.MerchantID != "" {
		q["MerchantId"] = req.Parameters.MerchantID
	}
	if req.Parameters.RelatedItemPage > 0 {
		q["RelatedItemPage"] = req.Parameters.RelatedItemPage
	}
	if req.Parameters.RelationshipType != "" {
		q["RelationshipType"] = req.Parameters.RelationshipType
	}
	if req.Parameters.SearchIndex != "" {
		q["SearchIndex"] = req.Parameters.SearchIndex
	}
	if req.Parameters.TruncateReviewsAt != nil {
		q["TruncateReviewsAt"] = *req.Parameters.TruncateReviewsAt
	}
	if req.Parameters.VariationPage > 0 {
		q["VariationPage"] = req.Parameters.VariationPage
	}
	q["ResponseGroup"] = req.Parameters.ResponseGroups
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
	respObj := ItemLookupResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// ItemLookup returns new request for ItemLookup
func (client *Client) ItemLookup(parameters ItemLookupParameters) *ItemLookupRequest {
	return &ItemLookupRequest{
		Client:     client,
		Parameters: parameters,
	}
}
