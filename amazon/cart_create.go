package amazon

import (
	"encoding/xml"
	"net/http"
	"strconv"
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

// CartCreateItem items in cart
type CartCreateItem struct {
	ASIN           string
	Quantity       int
	OfferListingID string
}

// Query returns map key and value
func (item CartCreateItem) Query() map[string]string {
	q := map[string]string{}
	if item.ASIN != "" {
		q["ASIN"] = item.ASIN
	}
	if item.OfferListingID != "" {
		q["OfferListingId"] = item.OfferListingID
	}
	if item.Quantity > 0 {
		q["Quantity"] = strconv.Itoa(item.Quantity)
	}
	return q
}

// CartCreateParameters represents parameters for CartCreate operation request
type CartCreateParameters struct {
	ResponseGroups []CartCreateResponseGroup
	ASIN           string
	Items          []CartCreateItem
}

// AddItemWithASIN adds item with ASIN and Quantity
func (p *CartCreateParameters) AddItemWithASIN(ASIN string, quantity int) {
	p.Items = append(p.Items, CartCreateItem{
		ASIN:     ASIN,
		Quantity: quantity,
	})
}

// AddItemWithOfferListingID adds item with OfferListingID and Quantity
func (p *CartCreateParameters) AddItemWithOfferListingID(offerListingID string, quantity int) {
	p.Items = append(p.Items, CartCreateItem{
		OfferListingID: offerListingID,
		Quantity:       quantity,
	})
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
	cartItems := req.Parameters.Items
	items := make([]map[string]string, len(cartItems))
	for i := range cartItems {
		items[i] = cartItems[i].Query()
	}
	q["Item"] = items
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
