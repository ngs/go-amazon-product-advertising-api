package amazon

import "testing"

func TestCartCreate(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.CartCreate()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartCreateBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.CartCreate().buildQuery()
	Test{0, len(q)}.Compare(t)
}
