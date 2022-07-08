package service_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/golang/mock/gomock"

	"github.com/shopsmart/ssm2env/internal/testutils"
	"github.com/shopsmart/ssm2env/pkg/service"
)

var (
	paramsMap = map[string]string{
		"foo": "<value-to-be-quoted>",
		"bar": "contains a single quote y'all",
		"baz": "uses multiple* special &characters, y'all",
	}

	paramsSlice []*ssm.Parameter
)

const (
	paramsPrefix = "/test/ssm2env"
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

var _ = Describe("Service", func() {

	var (
		inputPath = "/test/ssm2env"
		client    *testutils.MockSSMClient
		svc       service.Service
		err       error
	)

	BeforeEach(func() {
		client = testutils.NewMockSSMClient(testutils.MockController)
		svc, err = service.NewFromClient(client)
		Expect(err).Should(BeNil())
	})

	It("Should get the parameters from ssm", func() {
		in := ssm.GetParametersByPathInput{
			Path:           aws.String(inputPath),
			WithDecryption: aws.Bool(true),
			Recursive:      aws.Bool(false),
		}

		var _ = in

		client.
			EXPECT().
			GetParametersByPathPages(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error {
				resp := ssm.GetParametersByPathOutput{
					NextToken:  nil,
					Parameters: paramsSlice,
				}
				_ = fn(&resp, true)
				return nil
			}).
			Times(1)

		params, err := svc.GetParameters(inputPath, false)
		Expect(err).Should(BeNil())
		Expect(params).Should(Equal(paramsMap))
	})

	It("Should get the parameters from ssm recursively", func() {
		in := ssm.GetParametersByPathInput{
			Path:           aws.String(inputPath),
			WithDecryption: aws.Bool(true),
			Recursive:      aws.Bool(true),
		}

		var _ = in

		client.
			EXPECT().
			GetParametersByPathPages(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error {
				resp := ssm.GetParametersByPathOutput{
					NextToken:  nil,
					Parameters: paramsSlice,
				}
				_ = fn(&resp, true)
				return nil
			}).
			Times(1)

		params, err := svc.GetParameters(inputPath, true)
		Expect(err).Should(BeNil())
		Expect(params).Should(Equal(paramsMap))
	})

})
