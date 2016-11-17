package amazon

import "net/http"

// CartGetResponseGroup represents constants those are capable ResponseGroups parameter
type CartGetResponseGroup string

const (
	// CartGetResponseGroupCart is a constant for Cart response group
	CartGetResponseGroupCart CartGetResponseGroup = "Cart"
	// CartGetResponseGroupCartTopSellers is a constant for CartTopSellers response group
	CartGetResponseGroupCartTopSellers CartGetResponseGroup = "CartTopSellers"
	// CartGetResponseGroupCartSimilarities is a constant for CartSimilarities response group
	CartGetResponseGroupCartSimilarities CartGetResponseGroup = "CartSimilarities"
	// CartGetResponseGroupCartNewReleases is a constant for CartNewReleases response group
	CartGetResponseGroupCartNewReleases CartGetResponseGroup = "CartNewReleases"
)

// CartGetParameters represents parameters for CartGet operation request
type CartGetParameters struct {
	ResponseGroups []CartGetResponseGroup
}

// CartGetRequest represents request for CartGet operation
type CartGetRequest struct {
	Client     *Client
	Parameters CartGetParameters
}

// CartGetResponse represents response for CartGet operation
type CartGetResponse struct {
	Error error
}

func (req *CartGetRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *CartGetRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartGetRequest) operation() string {
	return "CartGet"
}

// Do sends request for the API
func (req *CartGetRequest) Do() (*CartGetResponse, error) {
	respObj := CartGetResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CartGet returns new request for CartGet
func (client *Client) CartGet(parameters CartGetParameters) *CartGetRequest {
	return &CartGetRequest{
		Client:     client,
		Parameters: parameters,
	}
}
