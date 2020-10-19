package secrets

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)


const (
	secretsManagerPrefix string = "aws-secretsmanager-secret:"
)

// SecretsManagerProvider implements an SSM secrets provider
type SecretsManagerProvider struct {
	Client secretsmanageriface.SecretsManagerAPI
}

// GetSecret returns the resolved value of an SSM key
func (s *SecretsManagerProvider) GetSecret(name string) (string, error) {
	key := strings.Replace(name, secretsManagerPrefix, "", 1)
	res, err := s.Client.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId:           aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("Could not retrieve a value from Secrets Manager for secret %s", err)
	}
	return aws.StringValue(res.SecretString), nil
}

// NewSecretsManagerProvider returns an initialised SSM provider
func NewSecretsManagerProvider() *SecretsManagerProvider {
	return &SecretsManagerProvider{
		Client: secretsmanager.New(NewAwsSession()),
	}
}
