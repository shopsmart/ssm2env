package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"

	"github.com/shopsmart/ssm2env/pkg/utils"
)

// Service will abstract out calls to AWS
type Service struct {
	SSMClient ssmiface.SSMAPI
}

// New creates a new service
func New() (*Service, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := &Service{
		SSMClient: ssm.New(sess),
	}

	return svc, nil
}

// GetParameters fetches all of the parameters under a path into a map
func (svc *Service) GetParameters(searchPath string) (map[string]string, error) {
	params := map[string]string{}
	getParametersByPathInput := &ssm.GetParametersByPathInput{
		Path:           aws.String(searchPath),
		WithDecryption: aws.Bool(true),
	}

	err := svc.SSMClient.GetParametersByPathPages(getParametersByPathInput, func(resp *ssm.GetParametersByPathOutput, lastPage bool) bool {
		for _, param := range resp.Parameters {
			name := utils.SSMParameterToEnvVar(*param.Name, searchPath)
			params[name] = *param.Value
		}
		return true
	})

	return params, err
}
