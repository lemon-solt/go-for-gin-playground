package testing_gin_ginkgo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hoge", func() {
	Context("愛はあるんかテスト", func() {
		It("愛あるITテスト", func() {
			Expect(nil).To(BeNil())    // エラーは発生しない(nil)ことが期待されます
			Expect("愛").To(Equal("愛")) // nameはBobだと期待されます

		})
	})

})
