package ssm2env_test

import (
	"bytes"
	"sort"
	"strings"

	_ "embed"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/pkg/service"
)

const Prefix = "/aws/service/global-infrastructure/regions"

//go:embed tests/expected.env
var RegionsEnv string

var _ = Describe("Ssm2env", func() {

	var (
		buffer *bytes.Buffer
		svc    service.Service
		err    error
	)

	BeforeEach(func() {
		buffer = bytes.NewBuffer([]byte{})
		svc, err = service.New()
		Expect(err).Should(BeNil())
	})

	It("Should collect the parameters and write the env formatted bytes to the buffer", func() {
		err = ssm2env.Collect(svc, buffer, Prefix)
		Expect(err).Should(BeNil())

		ss := strings.Split(buffer.String(), "\n")
		sort.Strings(ss[:len(ss)-1]) // The last line is blank
		sorted := strings.Join(ss, "\n")

		Expect(sorted).Should(Equal(RegionsEnv))
	})

})
