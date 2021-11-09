package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Interface for session
type ISession interface {
	client.ConfigProvider
}

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

// NewSession Creates a new Session
func NewSession(c *Config) *session.Session {
	s, _ := session.NewSession(&aws.Config{
		Region:      &c.Region,
		Credentials: credentials.NewStaticCredentials(c.AccessKeyID, c.SecretAccessKey, ""),
	})

	return s
}
