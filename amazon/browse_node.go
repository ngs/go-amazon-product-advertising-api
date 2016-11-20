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
