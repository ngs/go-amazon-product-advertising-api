package amazon

import "testing"

func TestCartAdd(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.CartAdd()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}
