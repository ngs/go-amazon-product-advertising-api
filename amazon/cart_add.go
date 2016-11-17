package amazon

import "net/http"

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

// CartAddParameters represents parameters for CartAdd operation request
type CartAddParameters struct {
	ResponseGroups []CartAddResponseGroup
}

// CartAddRequest represents request for CartAdd operation
type CartAddRequest struct {
	Client     *Client
	Parameters CartAddParameters
}

// CartAddResponse represents response for CartAdd operation
type CartAddResponse struct {
	Error error
}

func (req *CartAddRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *CartAddRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartAddRequest) operation() string {
	return "CartAdd"
}

// Do sends request for the API
func (req *CartAddRequest) Do() (*CartAddResponse, error) {
	respObj := CartAddResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CartAdd returns new request for CartAdd
func (client *Client) CartAdd(parameters CartAddParameters) *CartAddRequest {
	return &CartAddRequest{
		Client:     client,
		Parameters: parameters,
	}
}
