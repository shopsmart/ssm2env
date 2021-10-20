package tests_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source=../pkg/service/service.go -destination=mocks/gen_service_mocks.go -package mocks SSMClient,Service

var (
	mockController *gomock.Controller

	paramsMap = map[string]string{
		"foo": "<value-to-be-quoted>",
		"bar": "contains a single quote y'all",
		"baz": "uses multiple* special &characters, y'all",
	}

	paramsSlice []*ssm.Parameter
)

const (
	paramsPrefix    = "/test/ssm2env"
	paramsEnvString = `FOO='<value-to-be-quoted>'
BAR='contains a single quote y'"'"'all'
BAZ='uses multiple* special &characters, y'"'"'all'
`
)

func init() {
	paramsSlice = []*ssm.Parameter{}
	for key, value := range paramsMap {
		paramsSlice = append(paramsSlice, &ssm.Parameter{
			Name:  aws.String(fmt.Sprintf("%s/%s", paramsPrefix, key)),
			Value: aws.String(value),
		})
	}
}

func TestTests(t *testing.T) {
	mockController = gomock.NewController(t)
	defer mockController.Finish()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}
