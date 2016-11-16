package amazon

import "net/http"

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

func (req *CartClearRequest) buildQuery() map[string]interface{} {
	q := map[string]interface{}{}
	return q
}

func (req *CartClearRequest) httpMethod() string {
	return http.MethodGet
}

func (req *CartClearRequest) operation() string {
	return "CartClear"
}

// Do sends request for the API
func (req *CartClearRequest) Do() (*CartClearResponse, error) {
	_, err := req.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// CartClear returns new request for CartClear
func (client *Client) CartClear(responseGroups ...CartClearResponseGroup) *CartClearRequest {
	return &CartClearRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
