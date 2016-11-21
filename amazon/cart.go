package amazon

import "encoding/xml"

// Cart represents Cart
type Cart struct {
	XMLName               xml.Name `xml:"Cart"`
	Request               Request
	ID                    string `xml:"CartId"`
	HMAC                  string
	URLEncodedHMAC        string
	PurchaseURL           string
	MobileCartURL         string
	SubTotal              Price
	CartItems             CartItems
	SavedForLaterItems    SavedForLaterItems
	SimilarProducts       SimilarProducts
	NewReleases           NewReleases
	SimilarViewedProducts SimilarViewedProducts
}

// CartItems represents CartItems
type CartItems struct {
	SubTotal Price
	CartItem []CartItem
}

// CartItem represents CartItem
type CartItem struct {
	ID             string `xml:"CartItemId"`
	ASIN           string
	SellerNickname string
	Quantity       int
	Title          string
	ProductGroup   string
	Price          Price
	ItemTotal      Price
}

// SavedForLaterItems represents SavedForLaterItems
type SavedForLaterItems struct {
	SubTotal          Price
	SavedForLaterItem []SavedForLaterItem
}

// SavedForLaterItem represents SavedForLaterItem
type SavedForLaterItem struct {
	CartItem
}
