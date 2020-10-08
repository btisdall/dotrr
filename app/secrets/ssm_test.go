package secrets

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockSSMClient struct {
	ssmiface.SSMAPI
	secrets map[string]*ssm.Parameter
}

func generateParameterOutput(name, value string) *ssm.Parameter {
	return &ssm.Parameter{
		Value: aws.String(value),
	}
}

func (m *mockSSMClient) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return &ssm.GetParameterOutput{
		Parameter: m.secrets[*input.Name],
	}, nil
}

func newMockSSMClient() *mockSSMClient {
	client := mockSSMClient{
		secrets: make(map[string]*ssm.Parameter),
	}
	data := map[string]string{
		"/bentis/test": "verySekret",
		"hello":        "goodbye",
		"/this/that":   "bongo",
	}
	for k, v := range data {
		client.secrets[k] = generateParameterOutput(k, v)
	}
	return &client
}

func TestGetSecret(t *testing.T) {
	p := SSMProvider{Client: newMockSSMClient()}
	secret, _ := p.GetSecret("/bentis/test")
	expected := "verySekret"
	if secret != expected {
		t.Errorf("Expected %v got %v", expected, secret)
	}
}

func TestNewSSMProvider(t *testing.T) {
	p := NewSSMProvider()
	_, ok := p.Client.(*ssm.SSM)
	if !ok {
		t.Errorf("Expected Client to be *ssm.SSM got %T", p)
	}
}
