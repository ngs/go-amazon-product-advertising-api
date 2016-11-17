package amazon

import (
	"encoding/xml"
	"errors"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

const expectedGetBody = "AWSAccessKeyId=AK&AssociateTag=ngsio-22&Operation=Mock&Service=AWSECommerceService&Signature=wHPsmXHNme%2B%2F1bb39wTxqB51YgB2xBRe2r5WOzfqViQ%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01&array.1=foo&array.2=bar&array.3=baz&falsy=False&int=100&map.1.baz1=qux1&map.1.foo1=bar1&map.2.baz2=qux2&map.2.foo2=bar2&string=bar&truthy=True&uint=200"
const expectedPostBody = "AWSAccessKeyId=AK&AssociateTag=ngsio-22&Operation=Mock&Service=AWSECommerceService&Signature=yO2WsclEMEc357Q%2BCMSFn%2FNRh6DWbaJZM2zySysY%2F0A%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01&array.1=foo&array.2=bar&array.3=baz&falsy=False&int=100&map.1.baz1=qux1&map.1.foo1=bar1&map.2.baz2=qux2&map.2.foo2=bar2&string=bar&truthy=True&uint=200"

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
	client, _ := New("AK", "SK", "ngsio-22", "JP")
	for _, test := range []Test{
		Test{"AK", client.AccessKeyID},
		Test{"SK", client.SecretAccessKey},
		Test{"ngsio-22", client.AssociateTag},
		Test{RegionJapan, client.Region},
	} {
		test.Compare(t)
	}
}

func TestNewInvalidRegion(t *testing.T) {
	client, err := New("AK", "SK", "ngsio-22", "JAPAN")
	Test{"Invalid Region JAPAN", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptyRegion(t *testing.T) {
	client, err := New("AK", "SK", "ngsio-22", "")
	Test{"Region is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptyAccessKeyID(t *testing.T) {
	client, err := New("", "SK", "ngsio-22", RegionJapan)
	Test{"AccessKeyID is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptyAssociateTag(t *testing.T) {
	client, err := New("AK", "SK", "", RegionJapan)
	Test{"AssociateTag is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewEmptySecretAccessKey(t *testing.T) {
	client, err := New("AK", "", "ngsio-22", RegionJapan)
	Test{"SecretAccessKey is not specified", err.Error()}.Compare(t)
	if client != nil {
		t.Errorf(`Expected nil but got "%v"`, client)
	}
}

func TestNewFromEnvionment(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "AK")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "SK")
	os.Setenv("AWS_PRODUCT_REGION", "JP")
	os.Setenv("AWS_ASSOCIATE_TAG", "ngsio-22")
	client, _ := NewFromEnvionment()
	for _, test := range []Test{
		Test{"AK", client.AccessKeyID},
		Test{"SK", client.SecretAccessKey},
		Test{"ngsio-22", client.AssociateTag},
		Test{RegionJapan, client.Region},
	} {
		test.Compare(t)
	}
}

func TestClientEndpoint(t *testing.T) {
	secureClient, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	insecureClient, _ := New("AK", "SK", "ngsio-22", RegionJapan)
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

type mockResponse struct {
	XMLName xml.Name `xml:"mock"`
	Result  string   `xml:"result"`
}

func (mop *mockOperation) buildQuery() map[string]interface{} {
	return map[string]interface{}{
		"string": "bar",
		"array":  []string{"foo", "bar", "baz"},
		"uint":   uint(200),
		"int":    100,
		"falsy":  false,
		"truthy": true,
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
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	mockOp := &mockOperation{}
	url := client.SignedURL(mockOp)
	Test{
		"https://webservices.amazon.co.jp/onca/xml?" + expectedGetBody,
		url,
	}.Compare(t)
}

func TestDoGetRequest(t *testing.T) {
	defer gock.Off()
	t.SkipNow()
	gock.DisableNetworking()
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp/onca/xml?" + expectedGetBody).
		Reply(200).
		BodyString("<mock><result>OK</result></mock>")
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	mockOp := &mockOperation{}
	mockResp := mockResponse{}
	res, err := client.DoRequest(mockOp, &mockResp)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	Test{200, res.StatusCode}.Compare(t)
	Test{"OK", mockResp.Result}.Compare(t)
}

func TestDoPostRequest(t *testing.T) {
	defer gock.Off()
	gock.DisableNetworking()
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp").
		Post("/onca/xml").
		BodyString(expectedPostBody).
		Reply(200).
		BodyString("<mock><result>OK</result></mock>")
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	mockOp := &mockOperation{
		method: http.MethodPost,
	}
	mockResp := mockResponse{}
	res, err := client.DoRequest(mockOp, &mockResp)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	Test{200, res.StatusCode}.Compare(t)
	Test{"OK", mockResp.Result}.Compare(t)
}

func TestDoInvalidMethodRequest(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	mockOp := &mockOperation{
		method: http.MethodDelete,
	}
	mockResp := mockResponse{}
	res, err := client.DoRequest(mockOp, &mockResp)
	Test{"Unsupported HTTP method: DELETE", err.Error()}.Compare(t)
	if res != nil {
		t.Errorf("Expected nil but got %v", res)
	}
}

func TestDoHTTPError(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp/onca/xml").
		BodyString(expectedGetBody).
		ReplyError(errors.New("oops"))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	mockOp := &mockOperation{}
	mockResp := mockResponse{}
	res, err := client.DoRequest(mockOp, &mockResp)
	Test{
		"Get " + "https://webservices.amazon.co.jp/onca/xml?" + expectedGetBody + ": oops",
		err.Error()}.Compare(t)
	if res != nil {
		t.Errorf("Expected nil but got %v", res)
	}
}

func TestDoInvalidXML(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	gock.New("https://webservices.amazon.co.jp/onca/xml").
		BodyString(expectedGetBody).
		Reply(200).
		BodyString("<invalidmock><result>OK</result></invalidmock>")
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	client.AssociateTag = "ngsio-22"
	mockOp := &mockOperation{}
	mockResp := mockResponse{}
	res, err := client.DoRequest(mockOp, &mockResp)
	Test{
		"expected element type <mock> but have <invalidmock>",
		err.Error()}.Compare(t)
	if res != nil {
		t.Errorf("Expected nil but got %v", res)
	}
}

func TestDoErrorResponse(t *testing.T) {
	setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
	for _, op := range []string{
		"ItemSearch",
		"BrowseNodeLookup",
		"ItemLookup",
		"SimilarityLookup",
		"CartAdd",
		"CartClear",
		"CartCreate",
		"CartGet",
		"CartModify",
	} {
		fixtureIO, _ := os.Open("_fixtures/" + op + "ErrorResponse.xml")
		gock.New("https://webservices.amazon.co.jp/onca/xml").
			BodyString(expectedGetBody).
			Reply(200).
			Body(fixtureIO)
		client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
		client.AssociateTag = "ngsio-22"
		setNow(time.Parse(time.RFC822, "16 Nov 16 21:34 JST"))
		mockOp := &mockOperation{}
		mockResp := mockResponse{}
		res, err := client.DoRequest(mockOp, &mockResp)
		Test{
			"Error RequestExpired: Request has expired. Timestamp date is 2016-11-16T12:34:00Z. (c2fd7101-14f1-4c46-954b-d2bf492dd2eb)",
			err.Error()}.Compare(t)
		if res != nil {
			t.Errorf("Expected nil but got %v", res)
		}
	}

}
