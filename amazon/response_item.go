package amazon

import "encoding/xml"

type responseItems struct {
	XMLName xml.Name `xml:"Items"`
	Request requestNode
}

type requestNode struct {
	XMLName xml.Name `xml:"Request"`
	IsValid bool
	Errors  errorsNode
}

func (req requestNode) Error() error {
	if req.Errors.Error() != "" {
		return req.Errors
	}
	return nil
}
