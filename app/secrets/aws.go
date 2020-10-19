package secrets

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
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
			panic(fmt.Sprintf("Couldn't create AWS session: %s", error))
		}
	})
	return awsSession
}
