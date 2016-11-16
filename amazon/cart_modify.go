package amazon

import "net/http"

// CartModifyResponseGroup represents constants those are capable ResponseGroups parameter
type CartModifyResponseGroup string

const (
	// CartModifyResponseGroupCart is a constant for Cart response group
	CartModifyResponseGroupCart CartModifyResponseGroup = "Cart"
	// CartModifyResponseGroupCartSimilarities is a constant for CartSimilarities response group
	CartModifyResponseGroupCartSimilarities CartModifyResponseGroup = "CartSimilarities"
	// CartModifyResponseGroupCartNewReleases is a constant for CartNewReleases response group
	CartModifyResponseGroupCartNewReleases CartModifyResponseGroup = "CartNewReleases"
	// CartModifyResponseGroupCartTopSellers is a constant for CartTopSellers response group
	CartModifyResponseGroupCartTopSellers CartModifyResponseGroup = "CartTopSellers"
)

// CartModifyRequest represents request for CartModify operation
type CartModifyRequest struct {
	ResponseGroups []CartModifyResponseGroup
	Client         *Client
}

// CartModifyResponse represents response for CartModify operation
type CartModifyResponse struct {
	Error error
}

func (req *CartModifyRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *CartModifyRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartModifyRequest) operation() string {
	return "CartModify"
}

// Do sends request for the API
func (req *CartModifyRequest) Do() (*CartModifyResponse, error) {
	_, err := req.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// CartModify returns new request for CartModify
func (client *Client) CartModify(responseGroups ...CartModifyResponseGroup) *CartModifyRequest {
	return &CartModifyRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
