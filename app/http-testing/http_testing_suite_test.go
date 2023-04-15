package httptesting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHttpTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HttpTesting Suite")
}
