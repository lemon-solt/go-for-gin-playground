package testing_gin_ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestingGinGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "愛はあるんかスイーツ")
}
