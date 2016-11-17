package amazon

import (
	"encoding/xml"
	"time"
)

// Item represents item
type Item struct {
	XMLName        xml.Name `xml:"Item"`
	ASIN           string
	DetailPageURL  string
	SalesRank      int
	ItemLinks      ItemLinks
	SmallImage     Image
	MediumImage    Image
	LargeImage     Image
	ImageSets      ImageSets
	ItemAttributes ItemAttributes
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

// PublicationDate represents PublicationDate
type PublicationDate struct {
	time.Time
}

// UnmarshalXML parse time
func (c *PublicationDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02" // yyyymmdd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = PublicationDate{parse}
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
	PublicationDate   *PublicationDate
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
