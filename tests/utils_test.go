package tests_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shopsmart/ssm2env/pkg/utils"
)

var _ = Describe("Utils", func() {

	Describe("EnvFormat", func() {

		var (
			buffer *bytes.Buffer
		)

		BeforeEach(func() {
			buffer = bytes.NewBuffer([]byte{})
		})

		It("Should properly format a map of key values into env format", func() {
			err := utils.EnvFormat(buffer, paramsMap)
			Expect(err).Should(BeNil())
			Expect(buffer.String(), Equal(paramsEnvString))
		})

	})

})
