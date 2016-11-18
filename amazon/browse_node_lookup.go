package amazon

import (
	"encoding/xml"
	"net/http"
)

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
	BrowseNodeID   string
}

// BrowseNodeLookupRequest represents request for BrowseNodeLookup operation
type BrowseNodeLookupRequest struct {
	Client     *Client
	Parameters BrowseNodeLookupParameters
}

// BrowseNodeLookupResponse represents response for BrowseNodeLookup operation
type BrowseNodeLookupResponse struct {
	XMLName xml.Name    `xml:"BrowseNodeLookupResponse"`
	Results BrowseNodes `xml:"BrowseNodes"`
}

// Error returns Error found
func (res *BrowseNodeLookupResponse) Error() error {
	if e := res.Results.Request.Errors; e != nil {
		return e
	}
	return nil
}

// BrowseNodes returns found BrowseNodes
func (res *BrowseNodeLookupResponse) BrowseNodes() []BrowseNode {
	return res.Results.BrowseNode
}

func (req *BrowseNodeLookupRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	p := req.Parameters
	q["BrowseNodeId"] = p.BrowseNodeID
	q["ResponseGroup"] = p.ResponseGroups
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
	respObj := BrowseNodeLookupResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// BrowseNodeLookup returns new request for BrowseNodeLookup
func (client *Client) BrowseNodeLookup(parameters BrowseNodeLookupParameters) *BrowseNodeLookupRequest {
	return &BrowseNodeLookupRequest{
		Client:     client,
		Parameters: parameters,
	}
}
