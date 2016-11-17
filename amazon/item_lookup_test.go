package amazon

import "testing"

func TestItemLookup(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.ItemLookup(ItemLookupParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestItemLookupBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	q := client.ItemLookup(ItemLookupParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
