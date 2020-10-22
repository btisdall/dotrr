package secrets

import "strings"

const (
	escapePrefix string = "\\"
)

// Provider defines an interface for retrieving a secret
type Provider interface {
	GetSecret(string) (string, error)
}

// PassThruProvider is a provider that simply returns the value
type PassThruProvider struct{}

// GetSecret returns the resolved value
func (s *PassThruProvider) GetSecret(name string) (string, error) {
	return name, nil
}

// EscapeProvider is a provider that returns the value with the prefix "\" removed
type EscapeProvider struct{}

// GetSecret returns the resolved value
func (s *EscapeProvider) GetSecret(name string) (string, error) {
	return strings.Replace(name, escapePrefix, "", 1), nil
}

// GetProviderFunction is a function that returns a provider
type GetProviderFunction func(string) Provider

// GetProvider returns a secret provider based on the prefix of the value
func GetProvider(key string) Provider {
	switch {
	case strings.HasPrefix(key, ssmPrefix):
		return NewSSMProvider()
	case strings.HasPrefix(key, secretsManagerPrefix):
		return NewSecretsManagerProvider()
	case strings.HasPrefix(key, escapePrefix):
		return &EscapeProvider{}
	default:
		return &PassThruProvider{}
	}
}
