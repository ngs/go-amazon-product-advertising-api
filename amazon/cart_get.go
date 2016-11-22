package amazon

import "encoding/xml"

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
	CartID         string
	CartItemID     string
	HMAC           string
}

// CartGetRequest represents request for CartGet operation
type CartGetRequest struct {
	Client     *Client
	Parameters CartGetParameters
}

// CartGetResponse represents response for CartGet operation
type CartGetResponse struct {
	XMLName xml.Name `xml:"CartGetResponse"`
	Cart    Cart
}

// Error returns Error found
func (res *CartGetResponse) Error() error {
	if e := res.Cart.Request.Errors; e != nil {
		return e
	}
	return nil
}

// Query returns query for sending request
func (req *CartGetRequest) Query() map[string]interface{} {
	q := map[string]interface{}{}
	q["ResponseGroup"] = req.Parameters.ResponseGroups
	q["CartId"] = req.Parameters.CartID
	q["CartItemId"] = req.Parameters.CartItemID
	q["HMAC"] = req.Parameters.HMAC
	return q
}

func (req *CartGetRequest) httpMethod() string {
	return "GET"
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
	if err := respObj.Error(); err != nil {
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
