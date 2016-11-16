package amazon

// CartCreateResponseGroup represents constants those are capable ResponseGroups parameter
type CartCreateResponseGroup string

const (
	// CartCreateResponseGroupCart is a constant for Cart response group
	CartCreateResponseGroupCart CartCreateResponseGroup = "Cart"
	// CartCreateResponseGroupCartSimilarities is a constant for CartSimilarities response group
	CartCreateResponseGroupCartSimilarities CartCreateResponseGroup = "CartSimilarities"
	// CartCreateResponseGroupCartNewReleases is a constant for CartNewReleases response group
	CartCreateResponseGroupCartNewReleases CartCreateResponseGroup = "CartNewReleases"
	// CartCreateResponseGroupCartTopSellers is a constant for CartTopSellers response group
	CartCreateResponseGroupCartTopSellers CartCreateResponseGroup = "CartTopSellers"
)

// CartCreateRequest represents request for CartCreate operation
type CartCreateRequest struct {
	ResponseGroups []CartCreateResponseGroup
	Client         *Client
}

// CartCreateResponse represents response for CartCreate operation
type CartCreateResponse struct {
	Error error
}

func (req *CartCreateRequest) buildQuery() map[string]string {
	q := map[string]string{}
	return q
}

// Do send request for the API
func (req *CartCreateRequest) Do() (*CartCreateResponse, error) {
	return nil, nil
}

// CartCreate returns new request for CartCreate
func (client *Client) CartCreate(responseGroups ...CartCreateResponseGroup) *CartCreateRequest {
	return &CartCreateRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
