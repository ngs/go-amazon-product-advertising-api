package amazon

import "testing"

func TestCartRequestItemQuery(t *testing.T) {
	Test{
		map[string]string{
			"ASIN":     "test1",
			"Quantity": "2",
		},
		CartRequestItem{
			ASIN:     "test1",
			Quantity: 2,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"OfferListingId": "test1",
			"Quantity":       "2",
		},
		CartRequestItem{
			OfferListingID: "test1",
			Quantity:       2,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"ASIN": "test1",
		},
		CartRequestItem{
			ASIN: "test1",
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"OfferListingId": "test1",
		},
		CartRequestItem{
			OfferListingID: "test1",
		}.Query(),
	}.DeepEqual(t)
}

func TestCartCreateParametersAddItem(t *testing.T) {
	p := &CartRequestItems{}
	Test{0, len(p.Items)}.Compare(t)
	p.AddASIN("test1", 2)
	Test{1, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].ASIN}.Compare(t)
	Test{2, p.Items[0].Quantity}.Compare(t)
	p.AddOfferListingID("test2", 4)
	Test{2, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].ASIN}.Compare(t)
	Test{"", p.Items[0].OfferListingID}.Compare(t)
	Test{2, p.Items[0].Quantity}.Compare(t)
	Test{"", p.Items[1].ASIN}.Compare(t)
	Test{"test2", p.Items[1].OfferListingID}.Compare(t)
	Test{4, p.Items[1].Quantity}.Compare(t)
}
