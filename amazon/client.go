package amazon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	buildQuery() map[string]interface{}
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
func New(accessKeyID string, secretAccessKey string, region Region) (*Client, error) {
	if accessKeyID == "" {
		return nil, errors.New("AccessKeyID is not specified")
	}
	if secretAccessKey == "" {
		return nil, errors.New("SecretAccessKey is not specified")
	}
	if region == "" {
		return nil, errors.New("Region is not specified")
	}
	if region.Endpoint() == "" {
		return nil, fmt.Errorf("Invalid Region %v", region)
	}
	return &Client{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		Secure:          true,
	}, nil
}

// NewFromEnvionment returns new client from environment variables
func NewFromEnvionment() (*Client, error) {
	return New(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region(os.Getenv("AWS_REGION")),
	)
}

// Endpoint returns API endpoint
func (client *Client) Endpoint() string {
	if client.Secure {
		return client.Region.HTTPSEndpoint()
	}
	return client.Region.HTTPEndpoint()
}

func (client *Client) fillQuery(op OperationRequest) url.Values {
	ep := client.Endpoint()
	q := url.Values{}
	u, _ := url.Parse(ep)
	qmap := op.buildQuery()
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
		if str, ok := v.(string); ok {
			q.Set(k, str)
		} else if strar, ok := v.([]string); ok {
			for i, vv := range strar {
				q.Set(k+"."+strconv.Itoa(i+1), vv)
			}
		} else if strar, ok := v.([]map[string]string); ok {
			for i, m := range strar {
				for kk, vv := range m {
					q.Set(k+"."+strconv.Itoa(i+1)+"."+kk, vv)
				}
			}
		}
	}
	msg := op.httpMethod() + "\n" + u.Host + "\n" + u.Path + "\n" + q.Encode()
	mac := hmac.New(sha256.New, []byte(client.SecretAccessKey))
	mac.Write([]byte(msg))
	signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))
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
func (client *Client) DoRequest(op OperationRequest) (*http.Response, error) {
	url := client.SignedURL(op)
	method := op.httpMethod()
	if method == http.MethodGet {
		return http.Get(url)
	}
	if method == http.MethodPost {
		return http.PostForm(client.Endpoint(), client.fillQuery(op))
	}
	return nil, fmt.Errorf("Unsupported HTTP method: %v", method)
}
