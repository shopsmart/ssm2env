package ssm2env_test

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	_ "embed"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/spf13/viper"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/internal/testutils"
)

const Prefix = "/aws/service/global-infrastructure/regions"

var (
	//go:embed tests/expected.env
	RegionsEnv string

	//go:embed tests/regions.json
	RegionsJSON []byte

	RegionsMap map[string]interface{}
)

func init() {
	RegionsMap = map[string]interface{}{}
	err := json.Unmarshal(RegionsJSON, &RegionsMap)
	if err != nil {
		panic(err)
	}
}

var _ = Describe("Ssm2env Integration", func() {

	var (
		buffer *bytes.Buffer
		err    error
		cfg    *ssm2env.Config
	)

	BeforeEach(func() {
		buffer = bytes.NewBuffer([]byte{})
		Expect(err).Should(BeNil())
		cfg = &ssm2env.Config{
			SearchPath: Prefix,
		}
	})

	It("Should collect the parameters and write the env formatted bytes to the buffer", func() {
		err = ssm2env.WriteEnv(buffer, cfg)
		Expect(err).Should(BeNil())

		sorted := testutils.SortMultilineString(buffer.String())

		Expect(sorted).Should(Equal(RegionsEnv))
	})

	It("Should load the SSM parameters into the environment", func() {
		err = ssm2env.LoadEnv(cfg)
		Expect(err).Should(BeNil())

		for key, expected := range RegionsMap {
			actual := os.Getenv(strings.ToUpper(key))
			Expect(actual).Should(Equal(expected))
		}
	})

	It("Should load the SSM parameters into a viper config", func() {
		v := viper.New()

		err = ssm2env.LoadViper(v, cfg)
		Expect(err).Should(BeNil())

		actual := v.AllSettings()

		Expect(actual).Should(Equal(RegionsMap))
	})

})
