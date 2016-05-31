package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListSRPs", func() {
	Context("when given a symmetrix id", func() {
		It("returns a valid list of SRPs", func() {
			_, err := client.ListSRPs(symmetrixID)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
