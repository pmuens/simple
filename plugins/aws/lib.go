package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/pmuens/simple/util"
)

func CreateSession(region string) *session.Session {
	session, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		util.LogPanic(err)
	}

	return session
}
