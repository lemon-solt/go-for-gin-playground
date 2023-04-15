package httptesting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("愛のhttpでスクライブ", func() {
	Context("愛のhttp: Context", func() {
		It("", func() {
			Expect(nil).To(BeNil())    // エラーは発生しない(nil)ことが期待されます
			Expect("愛").To(Equal("愛")) // nameはBobだと期待されます
		})
	})

})
