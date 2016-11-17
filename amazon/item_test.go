package amazon

import (
	"encoding/xml"
	"testing"
	"time"
)

type TestDate struct {
	Date *Date
}

func TestUnmarshalDate(t *testing.T) {
	obj := TestDate{}
	if err := xml.Unmarshal([]byte("<TestDate><Date>2016-11-18</Date></TestDate>"), &obj); err != nil {
		t.Errorf("Got error %v", err)
	}
	Test{
		time.Date(2016, 11, 18, 0, 0, 0, 0, time.UTC).UnixNano(),
		obj.Date.UnixNano(),
	}.Compare(t)
}

func TestUnmarshalInvalidDate(t *testing.T) {
	obj := TestDate{}
	err := xml.Unmarshal([]byte("<TestDate><Date>2016/11/18</Date></TestDate>"), &obj)
	if err == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{
		`parsing time "2016/11/18" as "2006-01-02": cannot parse "/11/18" as "-"`,
		err.Error(),
	}.Compare(t)
}
