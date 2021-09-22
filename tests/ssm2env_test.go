package tests_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/tests/mocks"
)

var _ = Describe("Ssm2env", func() {

	var (
		buffer *bytes.Buffer
		svc    *mocks.MockService
	)

	BeforeEach(func() {
		buffer = bytes.NewBuffer([]byte{})
		svc = mocks.NewMockService(mockController)
	})

	It("Should collect the parameters and write the env formatted bytes to the buffer", func() {
		svc.
			EXPECT().
			GetParameters(paramsPrefix).
			Return(paramsMap, nil)

		err := ssm2env.Collect(svc, buffer, paramsPrefix)
		Expect(err).Should(BeNil())
		Expect(buffer.String()).Should(Equal(paramsEnvString))
	})

})
