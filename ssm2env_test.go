package ssm2env_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	_ "embed"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"gopkg.in/alessio/shellescape.v1"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/internal/testutils"
)

const MultilineString = `This is a multiline string.
There are six lines to this string.

A blank line was placed above this one and at the end.
This line has special characters: $({&*"'\}).
`

var _ = Describe("Ssm2env", func() {

	var (
		err error
		cfg *ssm2env.Config
		svc *testutils.MockService

		searchPath = "/does/not/matter"
		parameters = map[string]string{
			"Foo":                 "Bar",
			"includes-dashes-187": "foo",
			"special-<£#1":        "yes",
			"multiline":           MultilineString,
		}
		// singleLine = strings.ReplaceAll(MultilineString, "\n", "\\n")
		envVars = map[string]string{
			"FOO":                 "Bar",
			"INCLUDES_DASHES_187": "foo",
			"SPECIAL____1":        "yes",
			"MULTILINE":           MultilineString,
		}
		viperMap = map[string]interface{}{
			"foo":                 "Bar",
			"includes_dashes_187": "foo",
			"special____1":        "yes",
			"multiline":           MultilineString,
		}
	)

	BeforeEach(func() {
		Expect(err).Should(BeNil())
		cfg = &ssm2env.Config{
			SearchPath:       searchPath,
			Recursive:        false,
			MultilineSupport: true,
		}
		svc = testutils.NewMockService(testutils.MockController)

		svc.
			EXPECT().
			GetParameters(searchPath, false).
			Return(parameters, nil).
			Times(1)
	})

	It("Should collect the parameters and write the env formatted bytes to the buffer", func() {
		var expected []string
		for key, value := range envVars {
			expected = append(expected, fmt.Sprintf("%s=%s\n", key, shellescape.Quote(value)))
		}

		buffer := bytes.NewBuffer([]byte{})

		err = ssm2env.WriteEnvWithService(svc, buffer, cfg)
		Expect(err).Should(BeNil())

		lines := strings.Split(buffer.String(), "\n")
		for _, line := range lines {
			Expect(lines).Should(ContainElement(line))
		}
	})

	It("Should load the SSM parameters into the environment", func() {
		err = ssm2env.LoadEnvWithService(svc, cfg)
		Expect(err).Should(BeNil())

		for key, expected := range envVars {
			actual := os.Getenv(key)
			Expect(actual).Should(Equal(expected))
		}
	})

	It("Should load the SSM parameters into a viper config", func() {
		v := viper.New()

		err = ssm2env.LoadViperWithService(svc, v, cfg)
		Expect(err).Should(BeNil())

		actual := v.AllSettings()

		Expect(actual).Should(Equal(viperMap))
	})
})
