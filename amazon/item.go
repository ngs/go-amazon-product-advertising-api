package amazon

import (
	"encoding/xml"
	"time"
)

// Item represents item
type Item struct {
	XMLName         xml.Name `xml:"Item"`
	ASIN            string
	DetailPageURL   string
	SalesRank       int
	ItemLinks       ItemLinks
	SmallImage      Image
	MediumImage     Image
	LargeImage      Image
	ImageSets       ImageSets
	ItemAttributes  ItemAttributes
	OfferSummary    OfferSummary
	Offers          Offers
	CustomerReviews CustomerReviews
	SimilarProducts SimilarProducts
	BrowseNodes     BrowseNodes
}

// ItemLinks represents ItemLinks
type ItemLinks struct {
	ItemLink []ItemLink
}

// ItemLink represents ItemLink
type ItemLink struct {
	Description string
	URL         string
}

// Image represents Image
type Image struct {
	URL    string
	Height Size
	Width  Size
}

// Size represents Size
type Size struct {
	Value int    `xml:",chardata"`
	Units string `xml:",attr"`
}

// ImageSets represents ImageSets
type ImageSets struct {
	ImageSet []ImageSet
}

// ImageSet represents ImageSet
type ImageSet struct {
	Category       string `xml:",attr"`
	SwatchImage    Image
	SmallImage     Image
	ThumbnailImage Image
	TinyImage      Image
	MediumImage    Image
	LargeImage     Image
}

// Date represents short form date with yyyy-mm-dd date format
type Date struct {
	time.Time
}

// UnmarshalXML parse time
func (c *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02" // yyyy-mm-dd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = Date{parse}
	return nil
}

// ItemAttributes represents ItemAttributes
type ItemAttributes struct {
	Author            []string
	Binding           string
	Creator           Creator
	EAN               string
	EANList           EANList
	IsAdultProduct    bool
	ISBN              string
	Label             string
	Languages         Languages
	ListPrice         Price
	Manufacturer      string
	NumberOfPages     int
	PackageDimensions PackageDimensions
	ProductGroup      string
	ProductTypeName   string
	PublicationDate   *Date
	Publisher         string
	Studio            string
	Title             string
}

// Creator represents Creator
type Creator struct {
	Role string `xml:",attr"`
	Name string `xml:",chardata"`
}

// EANList represents EANList
type EANList struct {
	Element []string `xml:"EANListElement"`
}

// Languages represents Languages
type Languages struct {
	Language []Language
}

// Language represents Language
type Language struct {
	Name string
	Type string
}

// Price represents Price
type Price struct {
	Amount         string
	CurrencyCode   string
	FormattedPrice string
}

// PackageDimensions represents PackageDimensions
type PackageDimensions struct {
	Height Size
	Length Size
	Weight Size
	Width  Size
}

// OfferSummary represents OfferSummary
type OfferSummary struct {
	LowestNewPrice   Price
	LowestUsedPrice  Price
	TotalNew         int
	TotalUsed        int
	TotalCollectible int
	TotalRefurbished int
}

// Offers represents Offers
type Offers struct {
	TotalOffers     int
	TotalOfferPages int
	MoreOffersURL   string `xml:"MoreOffersUrl"`
	Offer           []Offer
}

// Offer represents Offer
type Offer struct {
	OfferAttributes OfferAttributes
	OfferListing    OfferListing
	LoyaltyPoints   LoyaltyPoints
}

// OfferAttributes represents OfferAttributes
type OfferAttributes struct {
	Condition string
}

// OfferListing represents OfferListing
type OfferListing struct {
	OfferListingID                  string `xml:"OfferListingId"`
	Price                           Price
	Availability                    string
	AvailabilityAttributes          AvailabilityAttributes
	IsEligibleForSuperSaverShipping bool
	IsEligibleForPrime              bool
}

// AvailabilityAttributes represents AvailabilityAttributes
type AvailabilityAttributes struct {
	AvailabilityType string
	MinimumHours     int
	MaximumHours     int
}

// LoyaltyPoints represents LoyaltyPoints
type LoyaltyPoints struct {
	Points                 int
	TypicalRedemptionValue Price
}

// CustomerReviews represents CustomerReviews
type CustomerReviews struct {
	IFrameURL  string
	HasReviews bool
}

// SimilarProducts represents SimilarProducts
type SimilarProducts struct {
	SimilarProduct []SimilarProduct
}

// SimilarProduct represents SimilarProduct
type SimilarProduct struct {
	ASIN  string
	Title string
}

// BrowseNodes represents BrowseNodes
type BrowseNodes struct {
	BrowseNode []BrowseNode
}

// BrowseNode represents BrowseNode
type BrowseNode struct {
	ID        string `xml:"BrowseNodeId"`
	Name      string
	Ancestors BrowseNodes
	Children  BrowseNodes
}
