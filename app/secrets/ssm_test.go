package secrets

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockSSMClient struct {
	ssmiface.SSMAPI
	value string
}

func (m *mockSSMClient) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return &ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Value: aws.String(m.value),
		},
	}, nil
}

func newMockSSMClient(value string) *mockSSMClient {
	return &mockSSMClient{
		value: value,
	}
}

func TestGetSecret(t *testing.T) {
	p := SSMProvider{Client: newMockSSMClient("verySekret")}
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
