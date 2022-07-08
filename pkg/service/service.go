package service

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMClient abstracts out the single function used within the ssmiface
type SSMClient interface {
	GetParametersByPathPages(getParametersByPathInput *ssm.GetParametersByPathInput, fn func(resp *ssm.GetParametersByPathOutput, lastPage bool) bool) error
}

// Service will abstract out calls to AWS
type Service interface {
	GetParameters(searchPath string, recursive bool) (map[string]string, error)
}

// Impl will be the implementation of the Service interface
type Impl struct {
	SSMClient SSMClient
}

// New creates a new service
func New() (Service, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return NewFromClient(ssm.New(sess))
}

// NewFromClient creates a new service from an implementation of the SSMClient
func NewFromClient(ssmClient SSMClient) (Service, error) {
	return &Impl{
		SSMClient: ssmClient,
	}, nil
}

// GetParameters fetches all of the parameters under a path into a map
func (svc *Impl) GetParameters(searchPath string, recursive bool) (map[string]string, error) {
	params := map[string]string{}
	getParametersByPathInput := &ssm.GetParametersByPathInput{
		Path:           aws.String(searchPath),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(recursive),
	}

	err := svc.SSMClient.GetParametersByPathPages(getParametersByPathInput, func(resp *ssm.GetParametersByPathOutput, lastPage bool) bool {
		for _, param := range resp.Parameters {
			name := strings.TrimPrefix(*param.Name, fmt.Sprintf("%s/", searchPath))
			params[name] = *param.Value
		}
		return true
	})

	return params, err
}
