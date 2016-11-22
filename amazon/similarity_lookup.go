package amazon

import (
	"encoding/xml"
	"net/http"
	"strings"
)

// SimilarityLookupResponseGroup represents constants those are capable ResponseGroups parameter
type SimilarityLookupResponseGroup string

const (
	// SimilarityLookupResponseGroupAccessories constant "Accessories"
	SimilarityLookupResponseGroupAccessories SimilarityLookupResponseGroup = "Accessories"
	// SimilarityLookupResponseGroupBrowseNodes constant "BrowseNodes"
	SimilarityLookupResponseGroupBrowseNodes SimilarityLookupResponseGroup = "BrowseNodes"
	// SimilarityLookupResponseGroupEditorialReview constant "EditorialReview"
	SimilarityLookupResponseGroupEditorialReview SimilarityLookupResponseGroup = "EditorialReview"
	// SimilarityLookupResponseGroupImages constant "Images"
	SimilarityLookupResponseGroupImages SimilarityLookupResponseGroup = "Images"
	// SimilarityLookupResponseGroupLarge constant "Large"
	SimilarityLookupResponseGroupLarge SimilarityLookupResponseGroup = "Large"
	// SimilarityLookupResponseGroupItemAttributes constant "ItemAttributes"
	SimilarityLookupResponseGroupItemAttributes SimilarityLookupResponseGroup = "ItemAttributes"
	// SimilarityLookupResponseGroupItemIds constant "ItemIds"
	SimilarityLookupResponseGroupItemIds SimilarityLookupResponseGroup = "ItemIds"
	// SimilarityLookupResponseGroupMedium constant "Medium"
	SimilarityLookupResponseGroupMedium SimilarityLookupResponseGroup = "Medium"
	// SimilarityLookupResponseGroupOffers constant "Offers"
	SimilarityLookupResponseGroupOffers SimilarityLookupResponseGroup = "Offers"
	// SimilarityLookupResponseGroupOfferSummary constant "OfferSummary"
	SimilarityLookupResponseGroupOfferSummary SimilarityLookupResponseGroup = "OfferSummary"
	// SimilarityLookupResponseGroupPromotionSummary constant "PromotionSummary"
	SimilarityLookupResponseGroupPromotionSummary SimilarityLookupResponseGroup = "PromotionSummary"
	// SimilarityLookupResponseGroupReviews constant "Reviews"
	SimilarityLookupResponseGroupReviews SimilarityLookupResponseGroup = "Reviews"
	// SimilarityLookupResponseGroupSalesRank constant "SalesRank"
	SimilarityLookupResponseGroupSalesRank SimilarityLookupResponseGroup = "SalesRank"
	// SimilarityLookupResponseGroupSimilarities constant "Similarities"
	SimilarityLookupResponseGroupSimilarities SimilarityLookupResponseGroup = "Similarities"
	// SimilarityLookupResponseGroupSmall constant "Small"
	SimilarityLookupResponseGroupSmall SimilarityLookupResponseGroup = "Small"
	// SimilarityLookupResponseGroupTracks constant "Tracks"
	SimilarityLookupResponseGroupTracks SimilarityLookupResponseGroup = "Tracks"
	// SimilarityLookupResponseGroupVariations constant "Variations"
	SimilarityLookupResponseGroupVariations SimilarityLookupResponseGroup = "Variations"
	// SimilarityLookupResponseGroupVariationSummary constant "VariationSummary"
	SimilarityLookupResponseGroupVariationSummary SimilarityLookupResponseGroup = "VariationSummary"
)

// SimilarityType typed constant for SimilarityType parameter
type SimilarityType string

const (
	// SimilarityTypeIntersection constant "Intersection"
	SimilarityTypeIntersection SimilarityType = "Intersection"
	// SimilarityTypeRandom constant "Random"
	SimilarityTypeRandom SimilarityType = "Random"
)

// SimilarityLookupParameters represents parameters for SimilarityLookup operation request
type SimilarityLookupParameters struct {
	ResponseGroups []SimilarityLookupResponseGroup
	Condition      Condition
	ItemIDs        []string
	MerchantID     string
	SimilarityType SimilarityType
}

// SimilarityLookupRequest represents request for SimilarityLookup operation
type SimilarityLookupRequest struct {
	Client     *Client
	Parameters SimilarityLookupParameters
}

// SimilarityLookupResponse represents response for SimilarityLookup operation
type SimilarityLookupResponse struct {
	XMLName xml.Name `xml:"SimilarityLookupResponse"`
	Items   Items    `xml:"Items"`
}

// Error returns Error found
func (res *SimilarityLookupResponse) Error() error {
	if e := res.Items.Request.Errors; e != nil {
		return e
	}
	return nil
}

func (req *SimilarityLookupRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	if string(req.Parameters.Condition) != "" {
		q["Condition"] = req.Parameters.Condition
	}
	if req.Parameters.MerchantID != "" {
		q["MerchantId"] = req.Parameters.MerchantID
	}
	if string(req.Parameters.SimilarityType) != "" {
		q["SimilarityType"] = string(req.Parameters.SimilarityType)
	}
	q["ItemId"] = strings.Join(req.Parameters.ItemIDs, ",")
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	return q
}

func (req *SimilarityLookupRequest) httpMethod() string {
	return http.MethodGet
}

func (req *SimilarityLookupRequest) operation() string {
	return "SimilarityLookup"
}

// Do sends request for the API
func (req *SimilarityLookupRequest) Do() (*SimilarityLookupResponse, error) {
	respObj := SimilarityLookupResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// SimilarityLookup returns new request for SimilarityLookup
func (client *Client) SimilarityLookup(parameters SimilarityLookupParameters) *SimilarityLookupRequest {
	return &SimilarityLookupRequest{
		Client:     client,
		Parameters: parameters,
	}
}
