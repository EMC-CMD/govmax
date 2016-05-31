package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListStorageGroup", func() {
	Context("when given a symmetrix id", func() {
		It("returns a listStoragegroupResponse", func() {
			_, err := client.ListStorageGroups(symmetrixID)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
