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

func TestBrowseNodeLookupBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.BrowseNodeLookup().buildQuery()
	Test{0, len(q)}.Compare(t)
}
