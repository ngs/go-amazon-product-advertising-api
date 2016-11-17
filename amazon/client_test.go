package amazon

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

func setNow(t time.Time, err error) {
	if err != nil {
		panic(err)
	}
	timeNowFunc = func() time.Time { return t }
}

type Test struct {
	expected interface{}
	actual   interface{}
}

func (test Test) Compare(t *testing.T) {
	if test.expected != test.actual {
		t.Errorf(`Expected "%v" but got "%v"`, test.expected, test.actual)
	}
}

func (test Test) DeepEqual(t *testing.T) {
	if !reflect.DeepEqual(test.expected, test.actual) {
		t.Errorf(`Expected "%v" but got "%v"`, test.expected, test.actual)
	}
}

func TestNew(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	for _, test := range []Test{
		Test{"AK", client.AccessKeyID},
		Test{"SK", client.SecretAccessKey},
		Test{RegionJapan, client.Region},
	} {
		test.Compare(t)
	}
}

func TestNewInvalidRegion(t *testing.T) {
	client, err := New("AK", "SK", "JAPAN")
	Test{"Invalid Region JAPAN", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptyRegion(t *testing.T) {
	client, err := New("AK", "SK", "")
	Test{"Region is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptyAccessKeyID(t *testing.T) {
	client, err := New("", "SK", RegionJapan)
	Test{"AccessKeyID is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptySecretAccessKey(t *testing.T) {
	client, err := New("AK", "", RegionJapan)
	Test{"SecretAccessKey is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewFromEnvionment(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "AK")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "SK")
	os.Setenv("AWS_REGION", "JP")
	client, _ := NewFromEnvionment()
	for _, test := range []Test{
		Test{"AK", client.AccessKeyID},
		Test{"SK", client.SecretAccessKey},
		Test{RegionJapan, client.Region},
	} {
		test.Compare(t)
	}
}

func TestClientEndpoint(t *testing.T) {
	secureClient, _ := New("AK", "SK", RegionJapan)
	insecureClient, _ := New("AK", "SK", RegionJapan)
	insecureClient.Secure = false
	for _, test := range []Test{
		Test{"https://webservices.amazon.co.jp/onca/xml", secureClient.Endpoint()},
		Test{"http://webservices.amazon.co.jp/onca/xml", insecureClient.Endpoint()},
	} {
		test.Compare(t)
	}
}

type mockOperation struct {
	method string
}

func (mop *mockOperation) buildQuery() map[string]interface{} {
	return map[string]interface{}{
		"string": "bar",
		"array":  []string{"foo", "bar", "baz"},
		"uint":   uint(200),
		"int":    100,
		"map": []map[string]string{
			map[string]string{"foo1": "bar1", "baz1": "qux1"},
			map[string]string{"foo2": "bar2", "baz2": "qux2"},
		},
	}
}

func (mop *mockOperation) httpMethod() string {
	if mop.method == "" {
		return http.MethodGet
	}
	return mop.method
}

func (mop *mockOperation) operation() string {
	return "Mock"
}

func TestClientSignedURL(t *testing.T) {
	client, _ := New("AK", "SK", RegionJapan)
	client.AssociateTag = "ngsio-22"
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	mockOp := &mockOperation{}
	url := client.SignedURL(mockOp)
	Test{
		"https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Operation=Mock&Service=AWSECommerceService&Signature=a8mPtL8sbgteMh0rKRDg4Cxi1H0i3D7M69zr5iXZBbA%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01&array.1=foo&array.2=bar&array.3=baz&int=100&map.1.baz1=qux1&map.1.foo1=bar1&map.2.baz2=qux2&map.2.foo2=bar2&string=bar&uint=200",
		url,
	}.Compare(t)
}

func TestDoGetRequest(t *testing.T) {
	defer gock.Off()
	gock.DisableNetworking()
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&Operation=Mock&Service=AWSECommerceService&Signature=a8mPtL8sbgteMh0rKRDg4Cxi1H0i3D7M69zr5iXZBbA%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01&array.1=foo&array.2=bar&array.3=baz&int=100&map.1.baz1=qux1&map.1.foo1=bar1&map.2.baz2=qux2&map.2.foo2=bar2&string=bar&uint=200").
		Reply(200).
		BodyString("ok")
	client, _ := New("AK", "SK", RegionJapan)
	client.AssociateTag = "ngsio-22"
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	mockOp := &mockOperation{}
	res, err := client.DoRequest(mockOp)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	Test{200, res.StatusCode}.Compare(t)
}

func TestDoPostRequest(t *testing.T) {
	defer gock.Off()
	gock.DisableNetworking()
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp").
		Post("/onca/xml").
		BodyString("AWSAccessKeyId=AK&AssociateTag=ngsio-22&Operation=Mock&Service=AWSECommerceService&Signature=HPbku1B2gRuCGvfQYuNdZCRSDBplZ4t-IeXe8Fcxvc0%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01&array.1=foo&array.2=bar&array.3=baz&int=100&map.1.baz1=qux1&map.1.foo1=bar1&map.2.baz2=qux2&map.2.foo2=bar2&string=bar&uint=200").
		Reply(200).
		BodyString("ok")
	client, _ := New("AK", "SK", RegionJapan)
	client.AssociateTag = "ngsio-22"
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	mockOp := &mockOperation{
		method: http.MethodPost,
	}
	res, err := client.DoRequest(mockOp)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	Test{200, res.StatusCode}.Compare(t)
}

func TestDoInvalidMethodRequest(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", RegionJapan)
	client.AssociateTag = "ngsio-22"
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	mockOp := &mockOperation{
		method: http.MethodDelete,
	}
	res, err := client.DoRequest(mockOp)
	Test{"Unsupported HTTP method: DELETE", err.Error()}.Compare(t)
	if res != nil {
		t.Errorf("Expected nil but got %v", res)
	}
}
