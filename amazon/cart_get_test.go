package amazon

import "testing"

func TestCartGet(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.CartGet()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartGetBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.CartGet().buildQuery()
	Test{0, len(q)}.Compare(t)
}
