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

// CartModifyAction constant for cart modify operation action
type CartModifyAction string

const (
	// CartModifyActionNone "None"
	CartModifyActionNone CartModifyAction = ""
	// CartModifyActionMoveToCart "MoveToCart"
	CartModifyActionMoveToCart CartModifyAction = "MoveToCart"
	// CartModifyActionSaveForLater "SaveForLater"
	CartModifyActionSaveForLater CartModifyAction = "SaveForLater"
)

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

// CartModifyRequestItems contains CartRequestItem
type CartModifyRequestItems struct {
	Items []CartModifyRequestItem
}

// Query returns list of map
func (cartItems *CartModifyRequestItems) Query() []map[string]string {
	items := make([]map[string]string, len(cartItems.Items))
	for i := range cartItems.Items {
		items[i] = cartItems.Items[i].Query()
	}
	return items
}

// CartModifyRequestItem represents items to send modify request
type CartModifyRequestItem struct {
	Quantity   *int
	CartItemID string
	Action     CartModifyAction
}

// ModifyQuantity modifies quantity for item
func (cartItems *CartModifyRequestItems) ModifyQuantity(cartItemID string, quantity int) {
	cartItems.Items = append(cartItems.Items, CartModifyRequestItem{
		CartItemID: cartItemID,
		Quantity:   &quantity,
	})
}

// MoveToCart moves item to cart
func (cartItems *CartModifyRequestItems) MoveToCart(cartItemID string) {
	cartItems.Items = append(cartItems.Items, CartModifyRequestItem{
		CartItemID: cartItemID,
		Action:     CartModifyActionMoveToCart,
	})
}

// SaveForLater saves item to later
func (cartItems *CartModifyRequestItems) SaveForLater(cartItemID string) {
	cartItems.Items = append(cartItems.Items, CartModifyRequestItem{
		CartItemID: cartItemID,
		Action:     CartModifyActionSaveForLater,
	})
}

// Query returns map key and value
func (item CartModifyRequestItem) Query() map[string]string {
	q := map[string]string{}
	if item.Quantity != nil {
		q["Quantity"] = strconv.Itoa(*item.Quantity)
	}
	if item.CartItemID != "" {
		q["CartItemId"] = item.CartItemID
	}
	if item.Action != CartModifyActionNone {
		q["Action"] = string(item.Action)
	}
	return q
}
