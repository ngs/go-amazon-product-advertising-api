package amazon

import "testing"

func TestError(t *testing.T) {
	Test{"", Error{}.Error()}.Compare(t)
	Test{"Error foo: bar", Error{Code: "foo", Message: "bar"}.Error()}.Compare(t)
}

func TestErrors(t *testing.T) {
	Test{"", Errors{ErrorNode: []Error{}}.Error()}.Compare(t)
	Test{"Error foo: bar", Errors{ErrorNode: []Error{Error{Code: "foo", Message: "bar"}}}.Error()}.Compare(t)
}
