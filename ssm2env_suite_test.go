package ssm2env_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shopsmart/ssm2env/internal/testutils"
)

func TestSsm2env(t *testing.T) {
	testutils.Setup(t)
	defer testutils.Teardown()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Ssm2env Suite")
}
