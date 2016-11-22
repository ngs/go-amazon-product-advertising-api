package amazon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var timeNowFunc = time.Now

const (
	// Version API Version
	Version string = "2013-08-01"
	// Service is service name
	Service string = "AWSECommerceService"
	// timestampFormat
	timestampFormat string = "2006-01-02T15:04:05Z"
)

// OperationRequest interface
type OperationRequest interface {
	httpMethod() string
	operation() string
	Query() map[string]interface{}
}

// Client AWAS Client
type Client struct {
	AccessKeyID     string
	SecretAccessKey string
	AssociateTag    string
	Secure          bool
	Region
}

// New returns new client
func New(accessKeyID string, secretAccessKey string, associateTag string, region Region) (*Client, error) {
	if accessKeyID == "" {
		return nil, errors.New("AccessKeyID is not specified")
	}
	if secretAccessKey == "" {
		return nil, errors.New("SecretAccessKey is not specified")
	}
	if associateTag == "" {
		return nil, errors.New("AssociateTag is not specified")
	}
	if region == "" {
		return nil, errors.New("Region is not specified")
	}
	if !region.IsValid() {
		return nil, fmt.Errorf("Invalid Region %v", region)
	}
	return &Client{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		AssociateTag:    associateTag,
		Secure:          true,
	}, nil
}

// NewFromEnvionment returns new client from environment variables
func NewFromEnvionment() (*Client, error) {
	return New(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_ASSOCIATE_TAG"),
		Region(os.Getenv("AWS_PRODUCT_REGION")),
	)
}

// Endpoint returns API endpoint
func (client *Client) Endpoint() string {
	if client.Secure {
		return client.Region.HTTPSEndpoint()
	}
	return client.Region.HTTPEndpoint()
}

func setQueryValue(q url.Values, key string, value interface{}) url.Values {
	refv := reflect.ValueOf(value)
	refKind := refv.Kind()
	if str, ok := value.(string); ok {
		q.Set(key, str)
		return q
	}
	if b, ok := value.(bool); ok {
		if b {
			q.Set(key, "True")
		} else {
			q.Set(key, "False")
		}
		return q
	}
	if refKind == reflect.String {
		q.Set(key, refv.String())
		return q
	}
	if num, ok := value.(int); ok {
		q.Set(key, strconv.Itoa(num))
		return q
	}
	if num, ok := value.(uint); ok {
		q.Set(key, strconv.Itoa(int(num)))
		return q
	}
	if refKind == reflect.Slice {
		len := refv.Len()
		if key == "ResponseGroup" {
			rgs := make([]string, len)
			for i := 0; i < len; i++ {
				rgs[i] = refv.Index(i).String()
			}
			q.Set(key, strings.Join(rgs, ","))
			return q
		}
		for i := 0; i < len; i++ {
			q = setQueryValue(q, key+"."+strconv.Itoa(i+1), refv.Index(i).Interface())
		}
	}
	if m, ok := value.(map[string]string); ok {
		for k, v := range m {
			q = setQueryValue(q, key+"."+k, v)
		}
		return q
	}
	return q
}

func (client *Client) fillQuery(op OperationRequest) url.Values {
	ep := client.Endpoint()
	q := url.Values{}
	u, _ := url.Parse(ep)
	qmap := op.Query()
	q.Set("Service", Service)
	q.Set("AWSAccessKeyId", client.AccessKeyID)
	q.Set("Version", Version)
	q.Set("Operation", op.operation())
	if client.AssociateTag != "" {
		q.Set("AssociateTag", client.AssociateTag)
	}
	ts := timeNowFunc().UTC().Format(timestampFormat)
	q.Set("Timestamp", ts)
	for k, v := range qmap {
		q = setQueryValue(q, k, v)
	}
	queryKeys := make([]string, 0, len(q))
	for key := range q {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)
	queryKeysAndValues := make([]string, len(queryKeys))
	for i, key := range queryKeys {
		k := strings.Replace(url.QueryEscape(key), "+", "%20", -1)
		v := strings.Replace(url.QueryEscape(q.Get(key)), "+", "%20", -1)
		queryKeysAndValues[i] = k + "=" + v
	}
	query := strings.Join(queryKeysAndValues, "&")
	msg := op.httpMethod() + "\n" + u.Host + "\n" + u.Path + "\n" + query
	mac := hmac.New(sha256.New, []byte(client.SecretAccessKey))
	mac.Write([]byte(msg))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	q.Set("Signature", signature)
	return q
}

// SignedURL returns signed URL with specified query
func (client *Client) SignedURL(op OperationRequest) string {
	ep := client.Endpoint()
	url, _ := url.Parse(ep)
	q := client.fillQuery(op)
	url.RawQuery = q.Encode()
	return url.String()
}

// DoRequest sends HTTP request
func (client *Client) DoRequest(op OperationRequest, responseObject interface{}) (*http.Response, error) {
	url := client.SignedURL(op)
	method := op.httpMethod()
	var res *http.Response
	var err error
	if method == http.MethodGet {
		res, err = http.Get(url)
	} else if method == http.MethodPost {
		res, err = http.PostForm(client.Endpoint(), client.fillQuery(op))
	} else {
		return nil, fmt.Errorf("Unsupported HTTP method: %v", method)
	}
	if err != nil {
		return nil, err
	}
	if data, _ := ioutil.ReadAll(res.Body); data != nil {
		// fmt.Println(string(data))
		if err = xml.Unmarshal(data, responseObject); err != nil {
			if e := newItemSearchErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newBrowseNodeLookupErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newItemLookupErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newSimilarityLookupErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newCartAddErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newCartClearErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newCartCreateErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newCartGetErrorResponse(data); e != nil {
				return nil, e
			}
			if e := newCartModifyErrorResponse(data); e != nil {
				return nil, e
			}
			return nil, err
		}
	}
	return res, err
}
