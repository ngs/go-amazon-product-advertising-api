package amazon

import (
	"encoding/xml"
	"net/http"
)

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
	ASIN           string
	Items          CartRequestItems
}

// CartCreateRequest represents request for CartCreate operation
type CartCreateRequest struct {
	Client     *Client
	Parameters CartCreateParameters
}

// CartCreateResponse represents response for CartCreate operation
type CartCreateResponse struct {
	XMLName xml.Name `xml:"CartCreateResponse"`
	Cart    Cart
}

// Error returns Error found
func (res *CartCreateResponse) Error() error {
	if e := res.Cart.Request.Errors; e != nil {
		return e
	}
	return nil
}

func (req *CartCreateRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	q["Item"] = req.Parameters.Items.Query()
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
	respObj := CartCreateResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CartCreate returns new request for CartCreate
func (client *Client) CartCreate(parameters CartCreateParameters) *CartCreateRequest {
	return &CartCreateRequest{
		Client:     client,
		Parameters: parameters,
	}
}
