package secrets

import (
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

var (
	ssmProvider *SSMProvider
	once        sync.Once
)

const (
	ssmPrefix string = "aws-ssm-parameter:"
)

// SSMProvider implements an SSM secrets provider
type SSMProvider struct {
	Client ssmiface.SSMAPI
}

// GetSecret returns the resolved value of an SSM key
func (s *SSMProvider) GetSecret(name string) (string, error) {
	ssmKey := strings.Replace(name, ssmPrefix, "", 1)
	res, err := s.Client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(ssmKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("Could not retrieve a value from SSM for secret %s", err)
	}
	return aws.StringValue(res.Parameter.Value), nil
}

// NewSSMProvider returns an initialised SSM provider singleton
func NewSSMProvider() *SSMProvider {
	once.Do(func() {
		sess, _ := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})

		ssmProvider = &SSMProvider{
			Client: ssm.New(sess),
		}
	})
	return ssmProvider
}
