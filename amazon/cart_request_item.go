package amazon

import "strconv"

// CartRequestItems contains CartRequestItem
type CartRequestItems struct {
	Items []CartRequestItem
}

// Query returns list of map
func (cartItems *CartRequestItems) Query() []map[string]string {
	items := make([]map[string]string, len(cartItems.Items))
	for i := range cartItems.Items {
		items[i] = cartItems.Items[i].Query()
	}
	return items
}

// CartRequestItem represents items to send create or add cart request
type CartRequestItem struct {
	ASIN           string
	Quantity       int
	OfferListingID string
}

// Query returns map key and value
func (item CartRequestItem) Query() map[string]string {
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

// AddASIN adds item with ASIN and Quantity
func (cartItems *CartRequestItems) AddASIN(ASIN string, quantity int) {
	cartItems.Items = append(cartItems.Items, CartRequestItem{
		ASIN:     ASIN,
		Quantity: quantity,
	})
}

// AddOfferListingID adds item with OfferListingID and Quantity
func (cartItems *CartRequestItems) AddOfferListingID(offerListingID string, quantity int) {
	cartItems.Items = append(cartItems.Items, CartRequestItem{
		OfferListingID: offerListingID,
		Quantity:       quantity,
	})
}
