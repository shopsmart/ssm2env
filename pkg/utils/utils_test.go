package utils_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shopsmart/ssm2env/internal/testutils"
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
				"foo":        "<value-to-be-quoted>",
				"bar":        "contains a single quote y'all",
				"baz":        "uses multiple* special &characters, y'all",
				"nested/foo": "demonstrates recursive sanitization",
			}
			paramsEnvString := `BAR='contains a single quote y'"'"'all'
BAZ='uses multiple* special &characters, y'"'"'all'
FOO='<value-to-be-quoted>'
NESTED_FOO='demonstrates recursive sanitization'
`

			err := utils.WriteEnv(buffer, paramsMap, true, false)
			Expect(err).Should(BeNil())

			sorted := testutils.SortMultilineString(buffer.String())
			Expect(sorted).Should(Equal(paramsEnvString))
		})

		It("Should not escape new lines if multiline support is enabled", func() {
			paramsMap := map[string]string{
				"foo": `multiline value
it could be over multiple lines
containing many, many, many words
`,
			}
			paramsEnvString := `FOO='multiline value
it could be over multiple lines
containing many, many, many words
'
`

			err := utils.WriteEnv(buffer, paramsMap, true, false)
			Expect(err).Should(BeNil())
			Expect(buffer.String()).Should(Equal(paramsEnvString))
		})

		It("Should escape new lines if multiline support is disabled", func() {
			paramsMap := map[string]string{
				"foo": `multiline value
it could be over multiple lines
containing many, many, many words
`,
			}
			paramsEnvString := `FOO='multiline value\nit could be over multiple lines\ncontaining many, many, many words\n'
`

			err := utils.WriteEnv(buffer, paramsMap, false, false)
			Expect(err).Should(BeNil())
			Expect(buffer.String()).Should(Equal(paramsEnvString))
		})

		It("Should prefix with export if enabled", func() {
			paramsMap := map[string]string{
				"foo": `multiline value
it could be over multiple lines
containing many, many, many words
`,
			}
			paramsEnvString := `export FOO='multiline value
it could be over multiple lines
containing many, many, many words
'
`

			err := utils.WriteEnv(buffer, paramsMap, true, true)
			Expect(err).Should(BeNil())
			Expect(buffer.String()).Should(Equal(paramsEnvString))
		})

	})

})
