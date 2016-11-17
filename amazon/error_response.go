package amazon

import (
	"encoding/xml"
	"fmt"
)

type errorNode struct {
	XMLName xml.Name `xml:"Error"`
	Code    string
	Message string
}

type errorsNode struct {
	Errors []errorNode `xml:"Error"`
}

// ErrorResponse represents error response from the API
type ErrorResponse interface {
	error
	Code() string
	Message() string
}

func (e errorResponseNode) code() string {
	return e.ErrorNode.Code
}

func (e errorResponseNode) message() string {
	return e.ErrorNode.Message
}

func (e errorResponseNode) Error() string {
	if e.code() != "" {
		return fmt.Sprintf("Error %v: %v (%v)", e.code(), e.message(), e.RequestID)
	}
	return ""
}

func (e errorsNode) code() string {
	if len(e.Errors) > 0 {
		return e.Errors[0].Code
	}
	return ""
}

func (e errorsNode) message() string {
	if len(e.Errors) > 0 {
		return e.Errors[0].Message
	}
	return ""
}

func (e errorsNode) Error() string {
	if e.code() != "" {
		return fmt.Sprintf("Error %v: %v", e.code(), e.message())
	}
	return ""
}

type errorResponseNode struct {
	ErrorNode errorNode `xml:"Error"`
	RequestID string    `xml:"RequestId"`
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
