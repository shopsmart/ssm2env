package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	log "github.com/sirupsen/logrus"

	"github.com/shopsmart/ssm2env/cmd"
	"github.com/shopsmart/ssm2env/pkg/service"
)

const (
	// EnvironmentVariable is the environment variable to search for a mock
	EnvironmentVariable = "TEST_SSM2ENV_PARAMETERS"
)

var version = "development"

func main() {
	log.Warn("Using ssm mock for params")

	svc, err := mockService()
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd.Execute(version, svc)
}

func mockService() (service.Service, error) {
	params := os.Getenv(EnvironmentVariable)
	if params == "" {
		return nil, fmt.Errorf("Missing required environment variable: %s", EnvironmentVariable)
	}

	client := mockSSMClient{
		Parameters: map[string]string{},
	}

	if err := json.Unmarshal([]byte(params), &client.Parameters); err != nil {
		return nil, err
	}

	return service.NewFromClient(&client)
}

type mockSSMClient struct {
	Parameters map[string]string
}

func (m *mockSSMClient) GetParametersByPathPages(getParametersByPathInput *ssm.GetParametersByPathInput, fn func(resp *ssm.GetParametersByPathOutput, lastPage bool) bool) error {
	p := []*ssm.Parameter{}

	prefix := *getParametersByPathInput.Path

	for key, value := range m.Parameters {
		p = append(p, &ssm.Parameter{
			Name:  aws.String(fmt.Sprintf("%s/%s", prefix, key)),
			Type:  aws.String("SecureString"),
			Value: aws.String(value),
		})
	}

	r := ssm.GetParametersByPathOutput{
		NextToken:  nil,
		Parameters: p,
	}

	_ = fn(&r, true)

	return nil
}
