package httptesting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Loved", func() {
	Context("loved test", func() {
		It("loved it test", func() {

			Expect("愛").To(Equal("愛"))
		})
	})
})
