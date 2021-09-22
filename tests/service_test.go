package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/golang/mock/gomock"

	"github.com/shopsmart/ssm2env/pkg/service"
	"github.com/shopsmart/ssm2env/tests/mocks"
)

var _ = Describe("Service", func() {

	var (
		inputPath = "/test/ssm2env"
		client    *mocks.MockSSMClient
		svc       service.Service
		err       error
	)

	BeforeEach(func() {
		client = mocks.NewMockSSMClient(mockController)
		svc, err = service.NewFromClient(client)
		Expect(err).Should(BeNil())
	})

	It("Should get the parameters from ssm", func() {
		in := ssm.GetParametersByPathInput{
			Path:           aws.String(inputPath),
			WithDecryption: aws.Bool(true),
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

		params, err := svc.GetParameters(inputPath)
		Expect(err).Should(BeNil())
		Expect(params).Should(Equal(paramsMap))
	})

})
