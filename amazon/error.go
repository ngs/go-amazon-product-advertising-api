package amazon

import (
	"encoding/xml"
	"fmt"
)

// ErrorCode error code http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ErrorMessages.html
type ErrorCode string

const (
	// ExactParameterRequirement AWS.ExactParameterRequirement
	ExactParameterRequirement ErrorCode = "AWS.ExactParameterRequirement"
	// ExceededMaximumParameterValues AWS.ExceededMaximumParameterValues
	ExceededMaximumParameterValues ErrorCode = "AWS.ExceededMaximumParameterValues"
	// InsufficientParameterValues AWS.InsufficientParameterValues
	InsufficientParameterValues ErrorCode = "AWS.InsufficientParameterValues"
	// InternalError AWS.InternalError For SOAP, this will be presented as a SOAP fault rather than an error.
	InternalError ErrorCode = "AWS.InternalError"
	// InvalidAccount AWS.InvalidAccount
	InvalidAccount ErrorCode = "AWS.InvalidAccount"
	// InvalidAssociate AWS.InvalidAssociate You registered as an Amazon Associate in the requested locale. For more information, see Becoming an Associate.
	InvalidAssociate ErrorCode = "AWS.InvalidAssociate"
	// InvalidEnumeratedParameter AWS.InvalidEnumeratedParameter For example, SearchIndex has an explicit list of valid values.
	InvalidEnumeratedParameter ErrorCode = "AWS.InvalidEnumeratedParameter"
	// InvalidISO8601Time AWS.InvalidISO8601Time For example, this error is returned if your request has an invalid value for the Version parameter.
	InvalidISO8601Time ErrorCode = "AWS.InvalidISO8601Time"
	// InvalidOperationForMarketplace AWS.InvalidOperationForMarketplace
	InvalidOperationForMarketplace ErrorCode = "AWS.InvalidOperationForMarketplace"
	// InvalidOperationParameter AWS.InvalidOperationParameter For example, if your request has the AsinSearch operation, you will receive an error because AsinSearch is no longer supported.
	InvalidOperationParameter ErrorCode = "AWS.InvalidOperationParameter"
	// InvalidParameterCombination AWS.InvalidParameterCombination For example, if the CartAdd operation includes an ASIN and OfferListingId, you will receive an error.
	InvalidParameterCombination ErrorCode = "AWS.InvalidParameterCombination"
	// InvalidParameterValue AWS.InvalidParameterValue
	InvalidParameterValue ErrorCode = "AWS.InvalidParameterValue"
	// InvalidResponseGroup AWS.InvalidResponseGroup
	InvalidResponseGroup ErrorCode = "AWS.InvalidResponseGroup"
	// InvalidServiceParameter AWS.InvalidServiceParameter
	InvalidServiceParameter ErrorCode = "AWS.InvalidServiceParameter"
	// InvalidSubscriptionID AWS.InvalidSubscriptionId
	InvalidSubscriptionID ErrorCode = "AWS.InvalidSubscriptionId"
	// MaximumParameterRequirement AWS.MaximumParameterRequirement
	MaximumParameterRequirement ErrorCode = "AWS.MaximumParameterRequirement"
	// MinimumParameterRequirement AWS.MinimumParameterRequirement
	MinimumParameterRequirement ErrorCode = "AWS.MinimumParameterRequirement"
	// MissingOperationParameter AWS.MissingOperationParameter
	MissingOperationParameter ErrorCode = "AWS.MissingOperationParameter"
	// MissingParameterCombination AWS.MissingParameterCombination
	MissingParameterCombination ErrorCode = "AWS.MissingParameterCombination"
	// MissingParameters AWS.MissingParameters
	MissingParameters ErrorCode = "AWS.MissingParameters"
	// MissingParameterValueCombination AWS.MissingParameterValueCombination For example, an ItemLookup request for a Universal Product Code (UPC) must include the IdType and ItemId parameters. The value of IdType must be UPC.
	MissingParameterValueCombination ErrorCode = "AWS.MissingParameterValueCombination"
	// MissingServiceParameter AWS.MissingServiceParameter
	MissingServiceParameter ErrorCode = "AWS.MissingServiceParameter"
	// ParameterOutOfRange AWS.ParameterOutOfRange For example, ItemSearch allows you to fetch search results per page with the ItemPage parameter. The range of values for ItemPage is 1 to 10. If the value you chose is less than 1 or greater than 10, an error is returned.
	ParameterOutOfRange ErrorCode = "AWS.ParameterOutOfRange"
	// ParameterRepeatedInRequest AWS.ParameterRepeatedInRequest
	ParameterRepeatedInRequest ErrorCode = "AWS.ParameterRepeatedInRequest"
	// RestrictedParameterValueCombination AWS.RestrictedParameterValueCombination
	RestrictedParameterValueCombination ErrorCode = "AWS.RestrictedParameterValueCombination"
	// ExceededMaximumCartItems AWS.ECommerceService.ExceededMaximumCartItems
	ExceededMaximumCartItems ErrorCode = "AWS.ECommerceService.ExceededMaximumCartItems"
	// InvalidCartID AWS.ECommerceService.InvalidCartId
	InvalidCartID ErrorCode = "AWS.ECommerceService.InvalidCartId"
	// InvalidHMAC AWS.ECommerceService.InvalidHMAC The HMAC value is a unique token that associates a shopping cart with an Amazon customer, and a specific session on the Amazon marketplace.
	InvalidHMAC ErrorCode = "AWS.ECommerceService.InvalidHMAC"
	// InvalidQuantity AWS.ECommerceService.InvalidQuantity
	InvalidQuantity ErrorCode = "AWS.ECommerceService.InvalidQuantity"
	// ItemAlreadyInCart AWS.ECommerceService.ItemAlreadyInCart
	ItemAlreadyInCart ErrorCode = "AWS.ECommerceService.ItemAlreadyInCart"
	// ItemNotAccessible AWS.ECommerceService.ItemNotAccessible
	ItemNotAccessible ErrorCode = "AWS.ECommerceService.ItemNotAccessible"
	// ItemNotEligibleForCart AWS.ECommerceService.ItemNotEligibleForCart
	ItemNotEligibleForCart ErrorCode = "AWS.ECommerceService.ItemNotEligibleForCart"
	// NoExactMatches AWS.ECommerceService.NoExactMatches
	NoExactMatches ErrorCode = "AWS.ECommerceService.NoExactMatches"
	// NoSimilarities AWS.ECommerceService.NoSimilarities
	NoSimilarities ErrorCode = "AWS.ECommerceService.NoSimilarities"
	// RequestThrottled RequestThrottled For more information about rates, see Efficiency Guidelines.
	RequestThrottled ErrorCode = "RequestThrottled"
)

// Errors represents Errors
type Errors struct {
	XMLName   xml.Name `xml:"Errors"`
	ErrorNode []Error  `xml:"Error"`
}

// Error represents Error
type Error struct {
	Code    ErrorCode
	Message string
}

func (e Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("Error %v: %v", e.Code, e.Message)
	}
	return ""
}

// Error returns error string
func (e Errors) Error() string {
	if len(e.ErrorNode) > 0 {
		return e.ErrorNode[0].Error()
	}
	return ""
}

// ErrorResponse represents error response from the API
type ErrorResponse interface {
	error
	Code() string
	Message() string
}

func (e errorResponseNode) Code() ErrorCode {
	return e.ErrorNode.Code
}

func (e errorResponseNode) Message() string {
	return e.ErrorNode.Message
}

func (e errorResponseNode) Error() string {
	if e.Code() != "" {
		return fmt.Sprintf("Error %v: %v (%v)", e.Code(), e.Message(), e.RequestID)
	}
	return ""
}

type errorResponseNode struct {
	ErrorNode Error  `xml:"Error"`
	RequestID string `xml:"RequestId"`
}

type itemSearchErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"ItemSearchErrorResponse"`
}

type browseNodeLookupErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"BrowseNodeLookupErrorResponse"`
}

type itemLookupErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"ItemLookupErrorResponse"`
}

type similarityLookupErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"SimilarityLookupErrorResponse"`
}

type cartAddErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"CartAddErrorResponse"`
}

type cartClearErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"CartClearErrorResponse"`
}

type cartCreateErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"CartCreateErrorResponse"`
}

type cartGetErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"CartGetErrorResponse"`
}

type cartModifyErrorResponse struct {
	errorResponseNode
	XMLName xml.Name `xml:"CartModifyErrorResponse"`
}

func newItemSearchErrorResponse(data []byte) error {
	e := itemSearchErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newBrowseNodeLookupErrorResponse(data []byte) error {
	e := browseNodeLookupErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newItemLookupErrorResponse(data []byte) error {
	e := itemLookupErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newSimilarityLookupErrorResponse(data []byte) error {
	e := similarityLookupErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newCartAddErrorResponse(data []byte) error {
	e := cartAddErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newCartClearErrorResponse(data []byte) error {
	e := cartClearErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newCartCreateErrorResponse(data []byte) error {
	e := cartCreateErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newCartGetErrorResponse(data []byte) error {
	e := cartGetErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}

func newCartModifyErrorResponse(data []byte) error {
	e := cartModifyErrorResponse{}
	xml.Unmarshal(data, &e)
	if e.Error() != "" {
		return e
	}
	return nil
}
