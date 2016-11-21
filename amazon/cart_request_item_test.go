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

func TestCartRequestItemsAddItem(t *testing.T) {
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

func TestCartModifyRequestItemQuery(t *testing.T) {
	Test{
		map[string]string{
			"CartItemId": "test1",
			"Quantity":   "2",
		},
		CartModifyRequestItem{
			CartItemID: "test1",
			Quantity:   &[]int{2}[0],
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"CartItemId": "test1",
			"Action":     "MoveToCart",
		},
		CartModifyRequestItem{
			CartItemID: "test1",
			Action:     CartModifyRequestItemActionMoveToCart,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"CartItemId": "test1",
			"Action":     "SaveForLater",
		},
		CartModifyRequestItem{
			CartItemID: "test1",
			Action:     CartModifyRequestItemActionSaveForLater,
		}.Query(),
	}.DeepEqual(t)
	Test{
		map[string]string{
			"CartItemId": "test1",
		},
		CartModifyRequestItem{
			CartItemID: "test1",
		}.Query(),
	}.DeepEqual(t)
}

func TestCartModifyRequestItemsAddItem(t *testing.T) {
	p := &CartModifyRequestItems{}
	Test{0, len(p.Items)}.Compare(t)
	p.ModifyQuantity("test1", 1)
	Test{1, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].CartItemID}.Compare(t)
	Test{1, *p.Items[0].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[0].Action}.Compare(t)
	p.SaveForLater("test2")
	Test{2, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].CartItemID}.Compare(t)
	Test{1, *p.Items[0].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[0].Action}.Compare(t)
	Test{"test2", p.Items[1].CartItemID}.Compare(t)
	if p.Items[1].Quantity != nil {
		t.Errorf("Expected nil but got %v", p.Items[1].Quantity)
	}
	Test{CartModifyRequestItemActionSaveForLater, p.Items[1].Action}.Compare(t)
	p.ModifyQuantity("test3", 0)
	Test{3, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].CartItemID}.Compare(t)
	Test{1, *p.Items[0].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[0].Action}.Compare(t)
	Test{"test2", p.Items[1].CartItemID}.Compare(t)
	if p.Items[1].Quantity != nil {
		t.Errorf("Expected nil but got %v", p.Items[1].Quantity)
	}
	Test{CartModifyRequestItemActionSaveForLater, p.Items[1].Action}.Compare(t)
	Test{"test3", p.Items[2].CartItemID}.Compare(t)
	Test{0, *p.Items[2].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[2].Action}.Compare(t)
	p.MoveToCart("test4")
	Test{4, len(p.Items)}.Compare(t)
	Test{"test1", p.Items[0].CartItemID}.Compare(t)
	Test{1, *p.Items[0].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[0].Action}.Compare(t)
	Test{"test2", p.Items[1].CartItemID}.Compare(t)
	if p.Items[1].Quantity != nil {
		t.Errorf("Expected nil but got %v", p.Items[1].Quantity)
	}
	Test{CartModifyRequestItemActionSaveForLater, p.Items[1].Action}.Compare(t)
	Test{"test3", p.Items[2].CartItemID}.Compare(t)
	Test{0, *p.Items[2].Quantity}.Compare(t)
	Test{CartModifyRequestItemActionNone, p.Items[2].Action}.Compare(t)
	Test{"test4", p.Items[3].CartItemID}.Compare(t)
	if p.Items[3].Quantity != nil {
		t.Errorf("Expected nil but got %v", p.Items[3].Quantity)
	}
	Test{CartModifyRequestItemActionMoveToCart, p.Items[3].Action}.Compare(t)
}
