package cmd

import (
	"reflect"
	"testing"

	"github.com/btisdall/dotrr/v2/app/config"
	"github.com/btisdall/dotrr/v2/app/secrets"
)

type TestProvider struct{}

func (d *TestProvider) GetSecret(s string) (string, error) {
	return "dummy:" + s, nil
}

func getTestProvider(_ string) secrets.Provider {
	return &TestProvider{}
}

func TestResolve(t *testing.T) {
	c := &config.Map{
		"This":  "That",
		"Apple": "Orange",
		"hello": "goodbye",
	}
	expected := config.Map{
		"This":  "dummy:That",
		"Apple": "dummy:Orange",
		"hello": "dummy:goodbye",
	}

	resolved := Resolve(c, getTestProvider)

	if !reflect.DeepEqual(resolved, expected) {
		t.Errorf("Map was not correctly resolved. Got: %v, expected: %v", resolved, expected)
	}
}
