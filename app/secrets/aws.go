package secrets

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/btisdall/dotrr/v2/app/util"
)

var (
	awsSession *session.Session
	once       sync.Once
)

// NewAwsSession returns a session singleton
func NewAwsSession() *session.Session {
	once.Do(func() {
		var error error
		awsSession, error = session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})
		if error != nil {
			util.Er("Couldn't create AWS session", error)
		}
	})
	return awsSession
}
