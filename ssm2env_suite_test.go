package ssm2env_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSsm2env(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ssm2env Suite")
}
