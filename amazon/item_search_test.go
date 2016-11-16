package amazon

import "testing"

func TestItemSearch(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.ItemSearch()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}
