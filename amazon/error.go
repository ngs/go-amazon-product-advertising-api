package amazon

import (
	"encoding/xml"
	"fmt"
)

// Errors represents Errors
type Errors struct {
	XMLName   xml.Name `xml:"Errors"`
	ErrorNode []Error  `xml:"Error"`
}

// Error represents Error
type Error struct {
	Code    string
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

func (e errorResponseNode) Code() string {
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
