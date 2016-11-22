package amazon

import "encoding/xml"

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

// CartModifyParameters represents parameters for CartModify operation request
type CartModifyParameters struct {
	ResponseGroups []CartModifyResponseGroup
	CartID         string
	HMAC           string
	Items          CartModifyRequestItems
}

// CartModifyRequest represents request for CartModify operation
type CartModifyRequest struct {
	Client     *Client
	Parameters CartModifyParameters
}

// CartModifyResponse represents response for CartModify operation
type CartModifyResponse struct {
	XMLName xml.Name `xml:"CartModifyResponse"`
	Cart    Cart
}

// Error returns Error found
func (res *CartModifyResponse) Error() error {
	if e := res.Cart.Request.Errors; e != nil {
		return e
	}
	return nil
}

// Query returns query for sending request
func (req *CartModifyRequest) Query() map[string]interface{} {
	q := map[string]interface{}{}
	q["CartId"] = req.Parameters.CartID
	q["HMAC"] = req.Parameters.HMAC
	q["Item"] = req.Parameters.Items.Query()
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	return q
}

func (req *CartModifyRequest) httpMethod() string {
	return "GET"
}

func (req *CartModifyRequest) operation() string {
	return "CartModify"
}

// Do sends request for the API
func (req *CartModifyRequest) Do() (*CartModifyResponse, error) {
	respObj := CartModifyResponse{}
	if _, err := req.Client.DoRequest(req, &respObj); err != nil {
		return nil, err
	}
	if err := respObj.Error(); err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CartModify returns new request for CartModify
func (client *Client) CartModify(parameters CartModifyParameters) *CartModifyRequest {
	return &CartModifyRequest{
		Client:     client,
		Parameters: parameters,
	}
}
