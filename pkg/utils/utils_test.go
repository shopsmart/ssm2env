package utils_test

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
			paramsMap := map[string]string{
				"foo": "<value-to-be-quoted>",
				"bar": "contains a single quote y'all",
				"baz": "uses multiple* special &characters, y'all",
			}
			paramsEnvString := `FOO='<value-to-be-quoted>'
BAR='contains a single quote y'"'"'all'
BAZ='uses multiple* special &characters, y'"'"'all'
`

			err := utils.EnvFormat(buffer, paramsMap)
			Expect(err).Should(BeNil())
			Expect(buffer.String(), Equal(paramsEnvString))
		})

	})

})
