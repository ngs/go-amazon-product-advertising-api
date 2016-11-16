package amazon

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

// BrowseNodeLookupRequest represents request for BrowseNodeLookup operation
type BrowseNodeLookupRequest struct {
	ResponseGroups []BrowseNodeLookupResponseGroup
	Client         *Client
}

// BrowseNodeLookupResponse represents response for BrowseNodeLookup operation
type BrowseNodeLookupResponse struct {
	Error error
}

// Do send request for the API
func (req *BrowseNodeLookupRequest) Do() (*BrowseNodeLookupResponse, error) {
	return nil, nil
}

// BrowseNodeLookup returns new request for BrowseNodeLookup
func (client *Client) BrowseNodeLookup(responseGroups ...BrowseNodeLookupResponseGroup) *BrowseNodeLookupRequest {
	return &BrowseNodeLookupRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
