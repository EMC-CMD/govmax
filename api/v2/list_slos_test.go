package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListSLOs", func() {
	Context("when given a symmetrix id", func() {
		It("returns a valid list of SLOs", func() {
			_, err := client.ListSLOs(symmetrixID)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
