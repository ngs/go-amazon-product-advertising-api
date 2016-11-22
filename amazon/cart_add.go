package amazon

import "encoding/xml"

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
	CartID         string
	HMAC           string
	Items          CartRequestItems
}

// CartAddRequest represents request for CartAdd operation
type CartAddRequest struct {
	Client     *Client
	Parameters CartAddParameters
}

// CartAddResponse represents response for CartAdd operation
type CartAddResponse struct {
	XMLName xml.Name `xml:"CartAddResponse"`
	Cart    Cart
}

// Error returns Error found
func (res *CartAddResponse) Error() error {
	if e := res.Cart.Request.Errors; e != nil {
		return e
	}
	return nil
}

// Query returns query for sending request
func (req *CartAddRequest) Query() map[string]interface{} {
	q := map[string]interface{}{}
	q["CartId"] = req.Parameters.CartID
	q["HMAC"] = req.Parameters.HMAC
	q["Item"] = req.Parameters.Items.Query()
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	return q
}

func (req *CartAddRequest) httpMethod() string {
	return "GET"
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
	if err := respObj.Error(); err != nil {
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
