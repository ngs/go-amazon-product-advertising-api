package amazon

import (
	"os"
	"testing"
)

type Test struct {
	expected interface{}
	actual   interface{}
}

func (test Test) Compare(t *testing.T) {
	if test.expected != test.actual {
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
