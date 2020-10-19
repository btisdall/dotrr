package secrets

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type mockSecretsManagerClient struct {
	secretsmanageriface.SecretsManagerAPI
	secretString string
}

func (m *mockSecretsManagerClient) GetSecretValue(_ *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	return &secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(m.secretString),
	}, nil
}

func newMockSecretsManagerClient(s string) *mockSecretsManagerClient {
	client := mockSecretsManagerClient{
		secretString: s,
	}
	return &client
}

func TestGetSecretsManagerSecret(t *testing.T) {
	expected := "verySekret"
	p := SecretsManagerProvider{Client: newMockSecretsManagerClient(expected)}
	secret, _ := p.GetSecret("/bentis/test")
	if secret != expected {
		t.Errorf("Expected %v got %v", expected, secret)
	}
}

func TestNewSecretsManagerProvider(t *testing.T) {
	p := NewSecretsManagerProvider()
	_, ok := p.Client.(*secretsmanager.SecretsManager)
	if !ok {
		t.Errorf("Expected Client to be *secretsmanager.SecretsManager got %T", p)
	}
}
