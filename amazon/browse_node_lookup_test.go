package amazon

import "testing"

func TestBrowseNodeLookup(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.BrowseNodeLookup()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}
