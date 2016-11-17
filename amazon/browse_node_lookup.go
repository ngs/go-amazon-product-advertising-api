package amazon

import "net/http"

// BrowseNodeLookupResponseGroup represents constants those are capable ResponseGroups parameter
type BrowseNodeLookupResponseGroup string

const (
	// BrowseNodeLookupResponseGroupBrowseNodeInfo is a constant for BrowseNodeInfo response group
	BrowseNodeLookupResponseGroupBrowseNodeInfo BrowseNodeLookupResponseGroup = "BrowseNodeInfo"
	// BrowseNodeLookupResponseGroupNewReleases is a constant for NewReleases response group
	BrowseNodeLookupResponseGroupNewReleases BrowseNodeLookupResponseGroup = "NewReleases"
	// BrowseNodeLookupResponseGroupMostGifted is a constant for MostGifted response group
	BrowseNodeLookupResponseGroupMostGifted BrowseNodeLookupResponseGroup = "MostGifted"
	// BrowseNodeLookupResponseGroupTopSellers is a constant for TopSellers response group
	BrowseNodeLookupResponseGroupTopSellers BrowseNodeLookupResponseGroup = "TopSellers"
	// BrowseNodeLookupResponseGroupMostWishedFor is a constant for MostWishedFor response group
	BrowseNodeLookupResponseGroupMostWishedFor BrowseNodeLookupResponseGroup = "MostWishedFor"
)

// BrowseNodeLookupParameters represents parameters for BrowseNodeLookup operation request
type BrowseNodeLookupParameters struct {
	ResponseGroups []BrowseNodeLookupResponseGroup
}

// BrowseNodeLookupRequest represents request for BrowseNodeLookup operation
type BrowseNodeLookupRequest struct {
	Client     *Client
	Parameters BrowseNodeLookupParameters
}

// BrowseNodeLookupResponse represents response for BrowseNodeLookup operation
type BrowseNodeLookupResponse struct {
	Error error
}

func (req *BrowseNodeLookupRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *BrowseNodeLookupRequest) httpMethod() string {
	return http.MethodGet
}

func (req *BrowseNodeLookupRequest) operation() string {
	return "BrowseNodeLookup"
}

// Do sends request for the API
func (req *BrowseNodeLookupRequest) Do() (*BrowseNodeLookupResponse, error) {
	_, err := req.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// BrowseNodeLookup returns new request for BrowseNodeLookup
func (client *Client) BrowseNodeLookup(parameters BrowseNodeLookupParameters) *BrowseNodeLookupRequest {
	return &BrowseNodeLookupRequest{
		Client:     client,
		Parameters: parameters,
	}
}
