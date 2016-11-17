package amazon

import "net/http"

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

// CartCreateParameters represents parameters for CartCreate operation request
type CartCreateParameters struct {
	ResponseGroups []CartCreateResponseGroup
}

// CartCreateRequest represents request for CartCreate operation
type CartCreateRequest struct {
	Client     *Client
	Parameters CartCreateParameters
}

// CartCreateResponse represents response for CartCreate operation
type CartCreateResponse struct {
	Error error
}

func (req *CartCreateRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *CartCreateRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartCreateRequest) operation() string {
	return "CartCreate"
}

// Do sends request for the API
func (req *CartCreateRequest) Do() (*CartCreateResponse, error) {
	_, err := req.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// CartCreate returns new request for CartCreate
func (client *Client) CartCreate(parameters CartCreateParameters) *CartCreateRequest {
	return &CartCreateRequest{
		Client:     client,
		Parameters: parameters,
	}
}
