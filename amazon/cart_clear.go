package amazon

// CartClearResponseGroup represents constants those are capable ResponseGroups parameter
type CartClearResponseGroup string

const (
	// CartClearResponseGroupCart is a constant for Cart response group
	CartClearResponseGroupCart CartClearResponseGroup = "Cart"
)

// CartClearRequest represents request for CartClear operation
type CartClearRequest struct {
	ResponseGroups []CartClearResponseGroup
	Client         *Client
}

// CartClearResponse represents response for CartClear operation
type CartClearResponse struct {
	Error error
}

func (req *CartClearRequest) buildQuery() map[string]string {
	q := map[string]string{}
	return q
}

// Do send request for the API
func (req *CartClearRequest) Do() (*CartClearResponse, error) {
	return nil, nil
}

// CartClear returns new request for CartClear
func (client *Client) CartClear(responseGroups ...CartClearResponseGroup) *CartClearRequest {
	return &CartClearRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
