package amazon

import "testing"

func TestBrowseNodeLookup(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.BrowseNodeLookup(BrowseNodeLookupParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestBrowseNodeLookupBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.BrowseNodeLookup(BrowseNodeLookupParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
