package secrets

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockSSMClient struct {
	ssmiface.SSMAPI
	secrets map[string]*ssm.Parameter
}

func generateParameterOutput(name, value string) *ssm.Parameter {
	return &ssm.Parameter{
		ARN:              aws.String("arn:aws:ssm:eu-west-1:111111111111:parameter/" + value),
		DataType:         aws.String("text"),
		LastModifiedDate: aws.Time(time.Now()),
		Name:             aws.String(name),
		Type:             aws.String("SecureString"),
		Value:            aws.String(value),
		Version:          aws.Int64(1),
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
