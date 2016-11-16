package amazon

import (
	"errors"
	"fmt"
	"os"
)

// Client AWAS Client
type Client struct {
	AccessKeyID     string
	SecretAccessKey string
	AssociateTag    string
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
