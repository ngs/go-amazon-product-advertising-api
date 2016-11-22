package amazon

import (
	"encoding/xml"
	"net/http"
)

// CartClearResponseGroup represents constants those are capable ResponseGroups parameter
type CartClearResponseGroup string

const (
	// CartClearResponseGroupCart is a constant for Cart response group
	CartClearResponseGroupCart CartClearResponseGroup = "Cart"
)

// CartClearParameters represents parameters for CartClear operation request
type CartClearParameters struct {
	ResponseGroups []CartClearResponseGroup
	CartID         string
	HMAC           string
}

// CartClearRequest represents request for CartClear operation
type CartClearRequest struct {
	Client     *Client
	Parameters CartClearParameters
}

// CartClearResponse represents response for CartClear operation
type CartClearResponse struct {
	XMLName xml.Name `xml:"CartClearResponse"`
	Cart    Cart
}

// Error returns Error found
func (res *CartClearResponse) Error() error {
	if e := res.Cart.Request.Errors; e != nil {
		return e
	}
	return nil
}

func (req *CartClearRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	q["CartId"] = req.Parameters.CartID
	q["HMAC"] = req.Parameters.HMAC
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	return q
}

func (req *CartClearRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartClearRequest) operation() string {
	return "CartClear"
}

// Do sends request for the API
func (req *CartClearRequest) Do() (*CartClearResponse, error) {
	respObj := CartClearResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CartClear returns new request for CartClear
func (client *Client) CartClear(parameters CartClearParameters) *CartClearRequest {
	return &CartClearRequest{
		Client:     client,
		Parameters: parameters,
	}
}
