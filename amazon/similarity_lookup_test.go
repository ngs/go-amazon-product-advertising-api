package amazon

import "testing"

func TestSimilarityLookup(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	req := client.SimilarityLookup(SimilarityLookupParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestSimilarityLookupBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	q := client.SimilarityLookup(SimilarityLookupParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}
