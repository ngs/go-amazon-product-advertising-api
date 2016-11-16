package amazon

// CartAddResponseGroup represents constants those are capable ResponseGroups parameter
type CartAddResponseGroup string

const (
	// CartAddResponseGroupCart is a constant for Cart response group
	CartAddResponseGroupCart CartAddResponseGroup = "Cart"
	// CartAddResponseGroupCartSimilarities is a constant for CartSimilarities response group
	CartAddResponseGroupCartSimilarities CartAddResponseGroup = "CartSimilarities"
	// CartAddResponseGroupCartNewReleases is a constant for CartNewReleases response group
	CartAddResponseGroupCartNewReleases CartAddResponseGroup = "CartNewReleases"
	// CartAddResponseGroupCartTopSellers is a constant for CartTopSellers response group
	CartAddResponseGroupCartTopSellers CartAddResponseGroup = "CartTopSellers"
)

// CartAddRequest represents request for CartAdd operation
type CartAddRequest struct {
	ResponseGroups []CartAddResponseGroup
	Client         *Client
}

// CartAddResponse represents response for CartAdd operation
type CartAddResponse struct {
	Error error
}

// Do send request for the API
func (req *CartAddRequest) Do() (*CartAddResponse, error) {
	return nil, nil
}

// CartAdd returns new request for CartAdd
func (client *Client) CartAdd(responseGroups ...CartAddResponseGroup) *CartAddRequest {
	return &CartAddRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
