package amazon

import "testing"

func TestCartCreate(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.CartCreate(CartCreateParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartCreateBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	q := client.CartCreate(CartCreateParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
