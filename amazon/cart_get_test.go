package amazon

import "testing"

func TestCartGet(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.CartGet(CartGetParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartGetBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	q := client.CartGet(CartGetParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
