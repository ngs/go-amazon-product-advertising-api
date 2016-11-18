package amazon

// BrowseNodes represents BrowseNodes
type BrowseNodes struct {
	BrowseNode []BrowseNode
	Request    Request
}

// BrowseNode represents BrowseNode
type BrowseNode struct {
	ID         string `xml:"BrowseNodeId"`
	Name       string
	Ancestors  BrowseNodes
	Children   BrowseNodes
	TopSellers TopSellers
	TopItemSet []TopItemSet
}

// TopSellers represents TopSellers
type TopSellers struct {
	TopSeller []TopSeller
}

// TopSeller represents TopSelle
type TopSeller struct {
	ASIN  string
	Title string
}

// NewReleases represents NewReleases
type NewReleases struct {
	NewRelease []NewRelease
}

// NewRelease represents NewRelease
type NewRelease struct {
	ASIN  string
	Title string
}

// TopItemSet represents TopItemSet
type TopItemSet struct {
	Type    string
	TopItem []TopItem
}

// TopItem represents TopItem
type TopItem struct {
	ASIN          string
	Title         string
	DetailPageURL string
	ProductGroup  string
	Author        string
}
