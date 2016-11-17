package amazon

import "testing"

func TestCartAdd(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.CartAdd(CartAddParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartAddBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	q := client.CartAdd(CartAddParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
